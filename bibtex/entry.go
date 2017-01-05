package bibtex

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/matthewdunsdon/bibliography"
)

const (
	entityTypeRequiredMessage  = "An entry type must be specified in order for entry to be encoded to BibTeX"
	citationKeyRequiredMessage = "A citation key must be specified in order for entry to be encoded to BibTeX"
)

// Entry describes the information associated with a bibliography in a BibTeX file
type Entry struct {
	bibliography.Entry
	CitationKey string
	EntryType   string
}

// EncodeBibTeX encodes an entry in a BibTeX format.
// To avoid the overhead of using reflection at run-time, each property is explicitly checked and if present is written to the stream.
func (e *Entry) EncodeBibTeX(writer io.Writer) (err error) {
	if len(e.EntryType) == 0 {
		return errors.New(entityTypeRequiredMessage)
	}
	if len(e.CitationKey) == 0 {
		return errors.New(citationKeyRequiredMessage)
	}

	_, err = fmt.Fprintf(writer, "@%s{%s", e.EntryType, e.CitationKey)
	if len(e.Address) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  address = {%s}", e.Address); err != nil {
			return
		}
	}
	if len(e.Author) > 0 {
		authorString := strings.Join(e.Author, " and ")
		if _, err = fmt.Fprintf(writer, ",\n  author = {%s}", authorString); err != nil {
			return
		}
	}
	if len(e.BookTitle) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  booktitle = {%s}", e.BookTitle); err != nil {
			return
		}
	}
	if len(e.Chapter) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  chapter = {%s}", e.Chapter); err != nil {
			return
		}
	}
	if len(e.Edition) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  edition = {%s}", e.Edition); err != nil {
			return
		}
	}
	if len(e.Editor) > 0 {
		editorString := strings.Join(e.Editor, " and ")
		if _, err = fmt.Fprintf(writer, ",\n  editor = {%s}", editorString); err != nil {
			return
		}
	}
	if len(e.HowPublished) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  howpublished = {%s}", e.HowPublished); err != nil {
			return
		}
	}
	if len(e.Institution) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  institution = {%s}", e.Institution); err != nil {
			return
		}
	}
	if len(e.Journal) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  journal = {%s}", e.Journal); err != nil {
			return
		}
	}
	if len(e.Key) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  key = {%s}", e.Key); err != nil {
			return
		}
	}
	if len(e.Month) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  month = {%s}", e.Month); err != nil {
			return
		}
	}
	if len(e.Note) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  note = {%s}", e.Note); err != nil {
			return
		}
	}
	if len(e.Number) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  number = {%s}", e.Number); err != nil {
			return
		}
	}
	if len(e.Organization) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  organization = {%s}", e.Organization); err != nil {
			return
		}
	}
	if len(e.Pages) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  pages = {%s}", e.Pages); err != nil {
			return
		}
	}
	if len(e.Publisher) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  publisher = {%s}", e.Publisher); err != nil {
			return
		}
	}
	if len(e.School) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  school = {%s}", e.School); err != nil {
			return
		}
	}
	if len(e.Series) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  series = {%s}", e.Series); err != nil {
			return
		}
	}
	if len(e.Title) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  title = {%s}", e.Title); err != nil {
			return
		}
	}
	if len(e.Type) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  type = {%s}", e.Type); err != nil {
			return
		}
	}
	if len(e.Volume) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  volume = {%s}", e.Volume); err != nil {
			return
		}
	}
	if len(e.Year) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  year = {%s}", e.Year); err != nil {
			return
		}
	}
	if err == nil {
		_, err = fmt.Fprint(writer, "\n}\n")
	}
	return
}
