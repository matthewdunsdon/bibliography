// Package bibtex implements encoding of BibTeX.
package bibtex

import (
	"io"

	"github.com/matthewdunsdon/bibliography"
)

// Entry represents the data needed to describe an entry in a BibTeX file
type Entry struct {
	CitationKey      string
	EntryType        string
	Entry            bibliography.Entry
	AdditionalFields map[string]string
}

// NewEncoder returns a new encoder that writes to writer.
func NewEncoder(writer io.Writer) (encoder *Encoder) {
	encoder = &Encoder{writer: writer}
	return
}
