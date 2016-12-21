package bibtex

import (
	"bytes"
	"errors"
	"sort"
	"testing"

	"github.com/matthewdunsdon/bibliography"
)

func sortedFieldClone(fields []field) (sortedClone []field) {
	sortedClone = make([]field, len(fields))
	copy(sortedClone, fields)
	sort.Sort(sortFieldByName{sortedClone})
	return
}

func TestGetEntityFields(t *testing.T) {
	testCases := []struct {
		testName         string
		entry            bibliography.Entry
		additionalFields map[string]string
		want             []field
	}{
		{testName: "empty"},
		{
			testName: "basic/book-DevOps",
			entry: bibliography.Entry{
				Author: []string{"Kyle Rankin"},
				Title:  "DevOps Troubleshooting: Linux Server Best Practices",
				Year:   "2012",
			},
			want: []field{
				field{name: "author", value: "Kyle Rankin"},
				field{name: "title", value: "DevOps Troubleshooting: Linux Server Best Practices"},
				field{name: "year", value: "2012"},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			got := getEntityFields(tc.entry, tc.additionalFields)
			if len(got) != len(tc.want) {
				t.Errorf("Expected the following fields to be returned %q, got %q", tc.want, got)
				return
			}

			for i, v := range sortedFieldClone(got) {
				if v != tc.want[i] {
					t.Errorf("Expected the following fields to be returned %q, got %q", tc.want, got)
					return
				}
			}
		})
	}
}

func TestWriteEntry(t *testing.T) {
	testCases := []struct {
		testName    string
		entryType   string
		citationKey string
		fields      []field
		want        string
	}{
		{
			testName:    "basic/book-DevOps",
			entryType:   "book",
			citationKey: "devOpsTroubleshooting",
			fields: []field{
				field{name: "author", value: "Kyle Rankin"},
				field{name: "title", value: "DevOps Troubleshooting: Linux Server Best Practices"},
				field{name: "year", value: "2012"},
			},
			want: "@book{devOpsTroubleshooting,\n  author = {Kyle Rankin},\n  title = {DevOps Troubleshooting: Linux Server Best Practices},\n  year = {2012}\n}\n",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			var writer bytes.Buffer
			writeEntry(&writer, tc.entryType, tc.citationKey, tc.fields)
			got := writer.String()
			if got != tc.want {
				t.Errorf("Expected the following content to be written %q, got %q", tc.want, got)
			}
		})
	}
}

func IgnoreTestEncode(t *testing.T) {
	testCases := []struct {
		testName string
		entry    Entry
		want     string
	}{
		{
			testName: "basic/book",
			entry: Entry{EntryType: "book", CitationKey: "rails",
				Entry: bibliography.Entry{
					Address:   "Raleigh, North Carolina",
					Author:    []string{"Ruby, Sam", "Thomas, Dave", "Hansson Heinemeier, David"},
					BookTitle: "Agile Web Development with Rails",
					Edition:   "third",
					Publisher: "The Pragmatic Bookshelf",
					Series:    "The Facets of Ruby",
					Title:     "Agile Web Development with Rails",
					Year:      "2009",
				},
				AdditionalFields: map[string]string{
					"keywords": "ruby, rails",
				},
			},
			want: `@book{rails,
  address = {Raleigh, North Carolina},
  author = {Ruby, Sam and Thomas, Dave and Hansson Heinemeier, David},
  booktitle = {Agile Web Development with Rails},
  edition = {third},
  keywords = {ruby, rails},
  publisher = {The Pragmatic Bookshelf},
  series = {The Facets of Ruby},
  title = {Agile Web Development with Rails},
  year = {2009}
}
`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			var writer bytes.Buffer
			encoder := &Encoder{writer: &writer}
			encoder.Encode(tc.entry)
			got := writer.String()
			if got != tc.want {
				t.Errorf("Expected to write proper test = %q, got %q", tc.want, got)
			}
		})
	}
}

// TestEncodeWithErrFieldSet tests that the err field is handled correctly
// As this is testing a property of the function, this could be converted to being a a property based test
// See https://golang.org/pkg/testing/quick/
func TestEncodeWithErrFieldSet(t *testing.T) {
	var writer bytes.Buffer
	originalErr := errors.New("Previous failure")
	encoder := &Encoder{writer: &writer, err: originalErr}
	err := encoder.Encode(Entry{})
	if err != originalErr {
		t.Errorf("Expected original error to be returned %q, got %q", originalErr, err)
	}
	if encoder.err != originalErr {
		t.Errorf("Expected err field not to be modified %q, got %q", originalErr, encoder.err)
	}
	if writer.Len() != 0 {
		t.Errorf("Expected writer field not to be written to, got %q", writer.String())
	}
}
