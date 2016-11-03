// Package encoding ...
package encoding

import (
	"io"

	"github.com/matthewdunsdon/bibliography"
)

// An Encoder writes BiBTeX to an output stream
type Encoder struct {
	writer io.Writer
	err    error
}

// NewEncoder returns a new encoder that writes to writer.
func NewEncoder(writer io.Writer) (encoder *Encoder) {
	encoder = &Encoder{writer: writer}
	return
}

// Encode writes the BiBTeX encoding of entry to the stream,
func (enc *Encoder) Encode(entry bibliography.Entry) (err error) {
	err = enc.err
	if err != nil {
		return
	}

	return
}
