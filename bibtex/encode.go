package bibtex

import (
	"fmt"
	"io"
)

// An Encoder writes BiBTeX to an output stream
type Encoder struct {
	writer io.Writer
	err    error
}

// Encode writes the BiBTeX encoding of entry to the stream,
func (enc *Encoder) Encode(entry Entry) (err error) {
	if enc.err != nil {
		return enc.err
	}

	_, err = fmt.Fprintf(enc.writer, "@%s{%s", entry.EntryType, entry.CitationKey)
	if err == nil {
		err = writeFieldsFromEntry(enc.writer, entry.Entry, entry.AdditionalFields)
	}
	if err == nil {
		_, err = fmt.Fprint(enc.writer, "\n}\n")
	}
	return
}
