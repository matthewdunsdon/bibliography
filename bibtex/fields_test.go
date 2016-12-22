package bibtex

import (
	"bytes"
	"testing"

	"github.com/matthewdunsdon/bibliography"
)

func TestWriteFieldsFromEntryWithBookAndIsbn(t *testing.T) {
	entry := bibliography.Entry{
		Author: []string{"Kyle Rankin"},
		Title:  "DevOps Troubleshooting: Linux Server Best Practices",
		Year:   "2012",
	}
	additionalFields := map[string]string{"isbn": "9780321832047"}
	want := ",\n  author = {Kyle Rankin},\n  title = {DevOps Troubleshooting: Linux Server Best Practices},\n  year = {2012},\n  isbn = {9780321832047}"

	var writer bytes.Buffer
	writeFieldsFromEntry(&writer, entry, additionalFields)
	got := writer.String()
	if got != want {
		t.Errorf("Expected to write proper test = %q, got %q", want, got)
	}
}

// TestWriteFieldsFromEntryAllFields tests that all fields are written correctly
// To keep this testing working as more fields are added, then either the test entry and want need to be generated
// or this could be converted to being a a property based test
// See https://golang.org/pkg/testing/quick/
func TestWriteFieldsFromEntryAllFields(t *testing.T) {
	entry := bibliography.Entry{
		Address:      "City, Country",
		Author:       []string{"Author, An", "Person, Another"},
		BookTitle:    "Title - of book",
		Chapter:      "2-4",
		Edition:      "first",
		Editor:       []string{"Editor, The"},
		HowPublished: "In test file",
		Institution:  "That place",
		Journal:      "A journal",
		Key:          "Misc Entry",
		Month:        "jan",
		Note:         "A note",
		Number:       "123.A",
		Organization: "Hello Org.",
		Pages:        "111-117, 131-160",
		Publisher:    "Pub",
		School:       "Excelsior",
		Series:       "Test series",
		Title:        "The Misc Entry",
		Type:         "Test",
		Volume:       "From the Volume: II",
		Year:         "2016",
	}
	var additionalFields map[string]string
	var writer bytes.Buffer

	writeFieldsFromEntry(&writer, entry, additionalFields)

	want := `,
  address = {City, Country},
  author = {Author, An and Person, Another},
  booktitle = {Title - of book},
  chapter = {2-4},
  edition = {first},
  editor = {Editor, The},
  howpublished = {In test file},
  institution = {That place},
  journal = {A journal},
  key = {Misc Entry},
  month = {jan},
  note = {A note},
  number = {123.A},
  organization = {Hello Org.},
  pages = {111-117, 131-160},
  publisher = {Pub},
  school = {Excelsior},
  series = {Test series},
  title = {The Misc Entry},
  type = {Test},
  volume = {From the Volume: II},
  year = {2016}`
	got := writer.String()
	if got != want {
		t.Errorf("Expected to write proper test = %q, got %q", want, got)
	}
}
