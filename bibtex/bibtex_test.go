package bibtex

import (
	"bytes"
	"io"
	"testing"
)

func TestNewEncoder(t *testing.T) {
	var buffer bytes.Buffer
	var suppliedWriter io.Writer = &buffer
	encoder := NewEncoder(suppliedWriter)

	if encoder.writer != suppliedWriter {
		t.Errorf("Expected writer field to be the supplied writer %v, got %v", suppliedWriter, encoder.writer)
	}
	if encoder.err != nil {
		t.Errorf("Expected err field to be nil, got %v", encoder.err)
	}
	if buffer.Len() != 0 {
		t.Errorf("Expected no data to be written when encoder created, got %q", buffer.String())
	}
}
