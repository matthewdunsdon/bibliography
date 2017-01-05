package bibtex

import (
	"fmt"
	"io"
)

// An Encoder writes BibTeX to an output stream.
type Encoder struct {
	writer io.Writer
	err    error
}

// NewEncoder returns a new encoder that writes to writer.
func NewEncoder(writer io.Writer) (encoder *Encoder) {
	encoder = &Encoder{writer: writer}
	return
}

// Encodable is the interface describing
type Encodable interface {
	EncodeBibTeX(io.Writer) error
}

// EncodeBibTeX writes the BibTeX encoding of entry to the stream.
func (e *Encoder) EncodeBibTeX(entry Encodable) error {
	if e.err != nil {
		return e.err
	}

	e.err = entry.EncodeBibTeX(e.writer)
	return e.err
}

// Encode writes the BibTeX encoding of entry to the stream.
func (e *Encoder) Encode(v interface{}) error {
	if e.err != nil {
		return e.err
	}

	if entry, ok := v.(Encodable); ok {
		return e.EncodeBibTeX(entry)
	}

	return fmt.Errorf("bibtex encoding not available for type %T", v)
}
