package bibtex

import (
	"sort"
	"strings"

	"github.com/matthewdunsdon/bibliography"
)

type field struct {
	name  string
	value string
}

type fields []field

func (f fields) Len() int      { return len(f) }
func (f fields) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

func appendFieldsFromEntry(entryFields []field, entry bibliography.Entry) []field {
	if entry.Address != "" {
		entryFields = append(entryFields, field{name: "address", value: entry.Address})
	}
	if len(entry.Author) > 0 {
		authorString := strings.Join(entry.Author, " and ")
		entryFields = append(entryFields, field{name: "author", value: authorString})
	}
	if entry.BookTitle != "" {
		entryFields = append(entryFields, field{name: "booktitle", value: entry.BookTitle})
	}
	if entry.Chapter != "" {
		entryFields = append(entryFields, field{name: "chapter", value: entry.Chapter})
	}
	if entry.Edition != "" {
		entryFields = append(entryFields, field{name: "edition", value: entry.Edition})
	}
	if len(entry.Editor) > 0 {
		editorString := strings.Join(entry.Editor, " and ")
		entryFields = append(entryFields, field{name: "address", value: editorString})
	}
	if entry.HowPublished != "" {
		entryFields = append(entryFields, field{name: "howpublished", value: entry.HowPublished})
	}
	if entry.Institution != "" {
		entryFields = append(entryFields, field{name: "institution", value: entry.Institution})
	}
	if entry.Journal != "" {
		entryFields = append(entryFields, field{name: "journal", value: entry.Journal})
	}
	if entry.Key != "" {
		entryFields = append(entryFields, field{name: "key", value: entry.Key})
	}
	if entry.Month != "" {
		entryFields = append(entryFields, field{name: "month", value: entry.Month})
	}
	if entry.Number != "" {
		entryFields = append(entryFields, field{name: "number", value: entry.Number})
	}
	if entry.Organization != "" {
		entryFields = append(entryFields, field{name: "organization", value: entry.Organization})
	}
	if entry.Pages != "" {
		entryFields = append(entryFields, field{name: "pages", value: entry.Pages})
	}
	if entry.Publisher != "" {
		entryFields = append(entryFields, field{name: "publisher", value: entry.Publisher})
	}
	if entry.School != "" {
		entryFields = append(entryFields, field{name: "school", value: entry.School})
	}
	if entry.Series != "" {
		entryFields = append(entryFields, field{name: "series", value: entry.Series})
	}
	if entry.Title != "" {
		entryFields = append(entryFields, field{name: "title", value: entry.Title})
	}
	if entry.Type != "" {
		entryFields = append(entryFields, field{name: "type", value: entry.Type})
	}
	if entry.Volume != "" {
		entryFields = append(entryFields, field{name: "volume", value: entry.Volume})
	}
	if entry.Year != "" {
		entryFields = append(entryFields, field{name: "year", value: entry.Year})
	}
	return entryFields
}

type sortFieldByName struct{ fields }

// Less reports whether the field with
// index i should sort before the element with index j.
func (s sortFieldByName) Less(i, j int) (isLess bool) {
	isLess = s.fields[i].name < s.fields[j].name
	return
}

func orderEntityByName(entryFields []field) {
	sort.Sort(sortFieldByName{entryFields})
}
