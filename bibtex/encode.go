package bibtex

import (
	"fmt"
	"io"

	"github.com/matthewdunsdon/bibliography"
)

// An Encoder writes BiBTeX to an output stream
type Encoder struct {
	writer io.Writer
	err    error
}

type EncoderOptions struct {
	sort bool
}

func getEntityFields(entry bibliography.Entry, additionalFields map[string]string) (fields []field) {
	fields = make([]field, 0, len(additionalFields))
	for key, value := range additionalFields {
		fields = append(fields, field{name: key, value: value})
	}
	fields = appendFieldsFromEntry(fields, entry)
	return
}

func writeEntry(writer io.Writer, entryType string, citationKey string, entryFields fields) (err error) {
	fmt.Fprintf(writer, "@%s{%s", entryType, citationKey)
	for _, field := range entryFields {
		fmt.Fprintf(writer, ",\n  %s = {%s}", field.name, field.value)
	}
	fmt.Fprint(writer, "\n}\n")
	return
}

// Encode writes the BiBTeX encoding of entry to the stream,
func (enc *Encoder) Encode(entry Entry) (err error) {
	err = enc.err
	if err != nil {
		return
	}

	entryFields := getEntityFields(entry.Entry, entry.AdditionalFields)
	orderEntityByName(entryFields)

	enc.err = writeEntry(enc.writer, entry.EntryType, entry.CitationKey, entryFields)
	err = enc.err
	return
}
