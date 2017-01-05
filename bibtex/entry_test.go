package bibtex

import (
	"bytes"
	"testing"

	"github.com/matthewdunsdon/bibliography"
)

func TestEncodeBibTeXWithBook(t *testing.T) {
	entry := Entry{EntryType: "book", CitationKey: "devOpsTroubleshooting",
		Entry: bibliography.Entry{
			Author: []string{"Kyle Rankin"},
			Title:  "DevOps Troubleshooting: Linux Server Best Practices",
			Year:   "2012",
		},
	}
	want := "@book{devOpsTroubleshooting,\n  author = {Kyle Rankin},\n  title = {DevOps Troubleshooting: Linux Server Best Practices},\n  year = {2012}\n}\n"

	var writer bytes.Buffer
	err := entry.EncodeBibTeX(&writer)
	got := writer.String()
	if err != nil {
		t.Errorf("Expected fields to be written to the writer without an error, got %v", err)
	}
	if got != want {
		t.Errorf("Expected the following text to be written to the writer %q, got %q", want, got)
	}
}

// TestEncodeBibTeXFromEntryAllFields tests that all fields are written correctly
// To keep this testing working as more fields are added, then either the test entry and want need to be generated
// or this could be converted to being a a property based test
// See https://golang.org/pkg/testing/quick/
func TestEncodeBibTeXFromEntryAllFields(t *testing.T) {
	entry := Entry{EntryType: "book", CitationKey: "devOpsTroubleshooting",
		Entry: bibliography.Entry{
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
		},
	}
	var writer bytes.Buffer

	err := entry.EncodeBibTeX(&writer)

	want := `@book{devOpsTroubleshooting,
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
  year = {2016}
}
`
	got := writer.String()
	if err != nil {
		t.Errorf("Expected fields to be written to the writer without an error, got %v", err)
	}
	if got != want {
		t.Errorf("Expected the following text to be written to the writer %q, got %q", want, got)
	}
}

func TestWriteFieldsWithoutEntryType(t *testing.T) {
	entry := Entry{EntryType: "", CitationKey: "devOpsTroubleshooting",
		Entry: bibliography.Entry{},
	}
	var writer bytes.Buffer

	err := entry.EncodeBibTeX(&writer)

	if err.Error() != entityTypeRequiredMessage {
		t.Errorf("Expected original error to be returned %q, got %q", entityTypeRequiredMessage, err)
	}
	if writer.Len() != 0 {
		t.Errorf("Expected writer field not to be written to, got %q", writer.String())
	}
}

func TestWriteFieldsWithoutCitationKey(t *testing.T) {
	entry := Entry{EntryType: "book", CitationKey: "",
		Entry: bibliography.Entry{},
	}
	var writer bytes.Buffer

	err := entry.EncodeBibTeX(&writer)

	if err.Error() != citationKeyRequiredMessage {
		t.Errorf("Expected original error to be returned %q, got %q", citationKeyRequiredMessage, err)
	}
	if writer.Len() != 0 {
		t.Errorf("Expected writer field not to be written to, got %q", writer.String())
	}
}
