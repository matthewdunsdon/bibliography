package bibtex

import (
	"io"

	"github.com/matthewdunsdon/bibliography"
)

/*
@STRING") {
		this.string();
} else if (d == "@PREAMBLE") {
		this.preamble();
} else if (d == "@COMMENT") {
		this.comment();
} else {
		this.entry(d);
*/
/*
const ( // iota is reset to 0
	c0 = iota // c0 == 0
	c1 = iota // c1 == 1
	c2 = iota // c2 == 2
)
*/

// Entry aa
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
