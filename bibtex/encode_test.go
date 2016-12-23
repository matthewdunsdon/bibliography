package bibtex

import (
	"bytes"
	"errors"
	"testing"

	"github.com/matthewdunsdon/bibliography"
)

func TestEncode(t *testing.T) {
	entry := Entry{EntryType: "book", CitationKey: "devOpsTroubleshooting",
		Entry: bibliography.Entry{
			Author: []string{"Kyle Rankin"},
			Title:  "DevOps Troubleshooting: Linux Server Best Practices",
			Year:   "2012",
		},
		AdditionalFields: map[string]string{"isbn": "9780321832047"},
	}

	var buffer bytes.Buffer
	encoder := &Encoder{writer: &buffer}
	err := encoder.Encode(entry)

	want := "@book{devOpsTroubleshooting,\n  author = {Kyle Rankin},\n  title = {DevOps Troubleshooting: Linux Server Best Practices},\n  year = {2012},\n  isbn = {9780321832047}\n}\n"
	got := buffer.String()
	if err != nil {
		t.Errorf("Expected entry to be encoded without an error, got %v", err)
	}
	if got != want {
		t.Errorf("Expected the following text to be written to the writer %q, got %q", want, got)
	}
}

// TestEncodeWithErrFieldSet tests that the err field is handled correctly
// As this is testing a property of the function, this could be converted to being a a property based test
// See https://golang.org/pkg/testing/quick/
func TestEncodeWithErrFieldSet(t *testing.T) {
	var buffer bytes.Buffer
	originalErr := errors.New("Previous failure")

	encoder := &Encoder{writer: &buffer, err: originalErr}
	err := encoder.Encode(Entry{})

	if err != originalErr {
		t.Errorf("Expected original error to be returned %q, got %q", originalErr, err)
	}
	if encoder.err != originalErr {
		t.Errorf("Expected err field not to be modified %q, got %q", originalErr, encoder.err)
	}
	if buffer.Len() != 0 {
		t.Errorf("Expected writer field not to be written to, got %q", buffer.String())
	}
}
