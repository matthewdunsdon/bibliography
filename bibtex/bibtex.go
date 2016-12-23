package bibtex

import (
	"io"

	"github.com/matthewdunsdon/bibliography"
)

// Entry represents the data associated with a bibtex entry
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
