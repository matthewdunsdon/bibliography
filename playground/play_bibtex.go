package main

import (
	"bytes"
	"log"

	"github.com/matthewdunsdon/bibliography"
	"github.com/matthewdunsdon/bibliography/bibtex"
)

func main() {
	myBook := bibtex.Entry{
		CitationKey: "devOpsTroubleshooting",
		Entry: bibliography.Entry{
			Author: []string{"Kyle Rankin"},
			Title:  "DevOps Troubleshooting: Linux Server Best Practices",
			Year:   "2012",
		},
		EntryType: "book",
	}

	var b bytes.Buffer
	encoder := bibtex.NewEncoder(&b)
	err := encoder.Encode(&myBook)

	if err != nil {
		log.Fatal(err)
		return
	}

	println(b.String())
}
