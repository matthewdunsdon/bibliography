package main

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

var scratch [64]byte

type encOpts struct {
	// quoted causes primitive fields to be encoded inside JSON strings.
	quoted bool
	// escapeHTML causes '<', '>', and '&' to be escaped in JSON strings.
	escapeHTML bool
}

type encoderFunc func(b *bytes.Buffer, v reflect.Value, opts encOpts)

func boolEncoder(b *bytes.Buffer, v reflect.Value, opts encOpts) {
	if opts.quoted {
		b.WriteByte('"')
	}
	if v.Bool() {
		b.WriteString("true")
	} else {
		b.WriteString("false")
	}
	if opts.quoted {
		b.WriteByte('"')
	}
}

func intEncoder(b *bytes.Buffer, v reflect.Value, opts encOpts) {
	d := strconv.AppendInt(scratch[:0], v.Int(), 10)
	if opts.quoted {
		b.WriteByte('"')
	}
	b.Write(d)
	if opts.quoted {
		b.WriteByte('"')
	}
}

func uintEncoder(b *bytes.Buffer, v reflect.Value, opts encOpts) {
	d := strconv.AppendUint(scratch[:0], v.Uint(), 10)
	if opts.quoted {
		b.WriteByte('"')
	}
	b.Write(d)
	if opts.quoted {
		b.WriteByte('"')
	}
}

type floatEncoder int // number of bits

var (
	float32Encoder = (floatEncoder(32)).encode
	float64Encoder = (floatEncoder(64)).encode
)

func (bits floatEncoder) encode(b *bytes.Buffer, v reflect.Value, opts encOpts) {
	f := v.Float()
	if math.IsInf(f, 0) || math.IsNaN(f) {
		panic(&UnsupportedValueError{v, strconv.FormatFloat(f, 'g', -1, int(bits))})
	}

	// Convert as if by ES6 number to string conversion.
	// This matches most other JSON generators.
	// See golang.org/issue/6384 and golang.org/issue/14135.
	// Like fmt %g, but the exponent cutoffs are different
	// and exponents themselves are not padded to two digits.
	d := scratch[:0]
	abs := math.Abs(f)
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if abs != 0 {
		if bits == 64 && (abs < 1e-6 || abs >= 1e21) || bits == 32 && (float32(abs) < 1e-6 || float32(abs) >= 1e21) {
			fmt = 'e'
		}
	}
	d = strconv.AppendFloat(d, f, fmt, -1, int(bits))
	if fmt == 'e' {
		// clean up e-09 to e-9
		n := len(d)
		if n >= 4 && d[n-4] == 'e' && d[n-3] == '-' && d[n-2] == '0' {
			d[n-2] = d[n-1]
			d = d[:n-1]
		}
	}

	if opts.quoted {
		b.WriteByte('"')
	}
	b.Write(d)
	if opts.quoted {
		b.WriteByte('"')
	}
}

func stringEncoder(b *bytes.Buffer, v reflect.Value, opts encOpts) {
	if v.Type() == numberType {
		numStr := v.String()
		// In Go1.5 the empty string encodes to "0", while this is not a valid number literal
		// we keep compatibility so check validity after this.
		if numStr == "" {
			numStr = "0" // Number's zero-val
		}
		if !isValidNumber(numStr) {
			panic(fmt.Errorf("json: invalid number literal %q", numStr))
		}
		b.WriteString(numStr)
		return
	}
	if opts.quoted {
		sb, err := Marshal(v.String())
		if err != nil {
			panic(err)
		}
		e.string(string(sb), opts.escapeHTML)
	} else {
		e.string(v.String(), opts.escapeHTML)
	}
}

func interfaceEncoder(b *bytes.Buffer, v reflect.Value, opts encOpts) {
	if v.IsNil() {
		b.WriteString("null")
		return
	}
	e.reflectValue(v.Elem(), opts)
}

type structEncoder struct {
	fields    []field
	fieldEncs []encoderFunc
}

func (se *structEncoder) encode(b *bytes.Buffer, v reflect.Value, opts encOpts) {
	b.WriteByte('{')
	first := true
	for i, f := range se.fields {
		fv := fieldByIndex(v, f.index)
		if !fv.IsValid() || f.omitEmpty && isEmptyValue(fv) {
			continue
		}
		if first {
			first = false
		} else {
			b.WriteByte(',')
		}
		e.string(f.name, opts.escapeHTML)
		b.WriteByte(':')
		opts.quoted = f.quoted
		se.fieldEncs[i](e, fv, opts)
	}
	b.WriteByte('}')
}

func newStructEncoder(t reflect.Type) encoderFunc {
	fields := cachedTypeFields(t)
	se := &structEncoder{
		fields:    fields,
		fieldEncs: make([]encoderFunc, len(fields)),
	}
	for i, f := range fields {
		se.fieldEncs[i] = newTypeEncoder(typeByIndex(t, f.index), true)
	}
	return se.encode
}

type mapEncoder struct {
	elemEnc encoderFunc
}

func (me *mapEncoder) encode(b *bytes.Buffer, v reflect.Value, opts encOpts) {
	if v.IsNil() {
		b.WriteString("null")
		return
	}
	b.WriteByte('{')

	// Extract and sort the keys.
	keys := v.MapKeys()
	sv := make([]reflectWithString, len(keys))
	for i, v := range keys {
		sv[i].v = v
		if err := sv[i].resolve(); err != nil {
			panic(&MarshalerError{v.Type(), err})
		}
	}
	// sort.Slice(sv, func(i, j int) bool { return sv[i].s < sv[j].s })

	for i, kv := range sv {
		if i > 0 {
			b.WriteByte(',')
		}
		e.string(kv.s, opts.escapeHTML)
		b.WriteByte(':')
		me.elemEnc(e, v.MapIndex(kv.v), opts)
	}
	b.WriteByte('}')
}

func newMapEncoder(t reflect.Type) encoderFunc {
	switch t.Key().Kind() {
	case reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
	default:
		if !t.Key().Implements(textMarshalerType) {
			return unsupportedTypeEncoder
		}
	}
	me := &mapEncoder{newTypeEncoder(t.Elem(), true)}
	return me.encode
}

// sliceEncoder just wraps an arrayEncoder, checking to make sure the value isn't nil.
type sliceEncoder struct {
	arrayEnc encoderFunc
}

func (se *sliceEncoder) encode(b *bytes.Buffer, v reflect.Value, opts encOpts) {
	if v.IsNil() {
		b.WriteString("null")
		return
	}
	se.arrayEnc(e, v, opts)
}

func newSliceEncoder(t reflect.Type) encoderFunc {
	// Byte slices get special treatment; arrays don't.
	if t.Elem().Kind() == reflect.Uint8 {
		p := reflect.PtrTo(t.Elem())
		if !p.Implements(marshalerType) && !p.Implements(textMarshalerType) {
			return encodeByteSlice
		}
	}
	enc := &sliceEncoder{newArrayEncoder(t)}
	return enc.encode
}

type arrayEncoder struct {
	elemEnc encoderFunc
}

func (ae *arrayEncoder) encode(b *bytes.Buffer, v reflect.Value, opts encOpts) {
	b.WriteByte('[')
	n := v.Len()
	for i := 0; i < n; i++ {
		if i > 0 {
			b.WriteByte(',')
		}
		ae.elemEnc(e, v.Index(i), opts)
	}
	b.WriteByte(']')
}

func newArrayEncoder(t reflect.Type) encoderFunc {
	enc := &arrayEncoder{newTypeEncoder(t.Elem(), true)}
	return enc.encode
}

type ptrEncoder struct {
	elemEnc encoderFunc
}

func (pe *ptrEncoder) encode(b *bytes.Buffer, v reflect.Value, opts encOpts) {
	if v.IsNil() {
		b.WriteString("null")
		return
	}
	pe.elemEnc(e, v.Elem(), opts)
}

func newPtrEncoder(t reflect.Type) encoderFunc {
	enc := &ptrEncoder{newTypeEncoder(t.Elem(), true)}
	return enc.encode
}

func unsupportedTypeEncoder(b *bytes.Buffer, v reflect.Value, _ encOpts) {
	panic(&UnsupportedTypeError{v.Type()})
}
