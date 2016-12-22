package bibtex

import (
	"fmt"
	"io"
	"strings"

	"github.com/matthewdunsdon/bibliography"
)

// writeFieldsFromEntry will write the contents of an entry and any additionalFields fields in a bibtex field format to a writer.
// To avoid the overhead of using reflection at run-time, each property is explicitly checked and if present is written to the stream.
// Further investigation into "go generate" is needed as it could move the overhead to compile-time and allow the code to be generated automatically.
func writeFieldsFromEntry(writer io.Writer, entry bibliography.Entry, additionalFields map[string]string) (err error) {
	if len(entry.Address) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  address = {%s}", entry.Address)
		if err != nil {
			return
		}
	}
	if len(entry.Author) > 0 {
		authorString := strings.Join(entry.Author, " and ")
		_, err = fmt.Fprintf(writer, ",\n  author = {%s}", authorString)
		if err != nil {
			return
		}
	}
	if len(entry.BookTitle) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  booktitle = {%s}", entry.BookTitle)
		if err != nil {
			return
		}
	}
	if len(entry.Chapter) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  chapter = {%s}", entry.Chapter)
		if err != nil {
			return
		}
	}
	if len(entry.Edition) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  edition = {%s}", entry.Edition)
		if err != nil {
			return
		}
	}
	if len(entry.Editor) > 0 {
		editorString := strings.Join(entry.Editor, " and ")
		_, err = fmt.Fprintf(writer, ",\n  editor = {%s}", editorString)
		if err != nil {
			return
		}
	}
	if len(entry.HowPublished) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  howpublished = {%s}", entry.HowPublished)
		if err != nil {
			return
		}
	}
	if len(entry.Institution) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  institution = {%s}", entry.Institution)
		if err != nil {
			return
		}
	}
	if len(entry.Journal) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  journal = {%s}", entry.Journal)
		if err != nil {
			return
		}
	}
	if len(entry.Key) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  key = {%s}", entry.Key)
		if err != nil {
			return
		}
	}
	if len(entry.Month) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  month = {%s}", entry.Month)
		if err != nil {
			return
		}
	}
	if len(entry.Note) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  note = {%s}", entry.Note)
		if err != nil {
			return
		}
	}
	if len(entry.Number) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  number = {%s}", entry.Number)
		if err != nil {
			return
		}
	}
	if len(entry.Organization) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  organization = {%s}", entry.Organization)
		if err != nil {
			return
		}
	}
	if len(entry.Pages) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  pages = {%s}", entry.Pages)
		if err != nil {
			return
		}
	}
	if len(entry.Publisher) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  publisher = {%s}", entry.Publisher)
		if err != nil {
			return
		}
	}
	if len(entry.School) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  school = {%s}", entry.School)
		if err != nil {
			return
		}
	}
	if len(entry.Series) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  series = {%s}", entry.Series)
		if err != nil {
			return
		}
	}
	if len(entry.Title) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  title = {%s}", entry.Title)
		if err != nil {
			return
		}
	}
	if len(entry.Type) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  type = {%s}", entry.Type)
		if err != nil {
			return
		}
	}
	if len(entry.Volume) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  volume = {%s}", entry.Volume)
		if err != nil {
			return
		}
	}
	if len(entry.Year) > 0 {
		_, err = fmt.Fprintf(writer, ",\n  year = {%s}", entry.Year)
		if err != nil {
			return
		}
	}
	for key, value := range additionalFields {
		if err == nil {
			_, err = fmt.Fprintf(writer, ",\n  %s = {%s}", key, value)
		}
	}
	return
}
