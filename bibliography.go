// Package bibliography defines interfaces shared by other packages that convert
// bibliography to and from textual representations
package bibliography

// An Entry contains the common BibTeX field types.
//
// See https://en.wikipedia.org/wiki/BibTeX#Field_types for a list of the common field types used in BibTeX.
type Entry struct {

	// Address is the partial or full address of the publisher.
	// Often, this refers just to the city the publisher is located in.
	Address string

	// Author is the list of author(s) associated with the cited material.
	Author []string

	// BookTitle is the title of the book and is used when only part of of the book is being cited.
	BookTitle string

	// Chapter is the chapter number from which the citation is being made
	Chapter string

	// Edition is the edition of a book.
	// It is good practice to express this in long form such as "First" or "Second"
	Edition string

	// Editor is the list of editor(s) associated with the cited material
	Editor []string

	// HowPublished is used to describe the publishing method of the cited material when it was not a standard method.
	HowPublished string

	// Institution is used to name of the institution who were involved in the cited material.
	// This typically is used when the institution is not the actual publisher.
	Institution string

	// Journal is the name for the journal or magazine that the citation was taken from.
	Journal string

	// Key is used specifying or overriding the alphabetical order of entries
	Key string

	// Month is the month of publication.
	// For unpublished material, the month of creation can be used instead.
	Month string

	// Note is any addition information that is needed to describe the cited material.
	Note string

	// Number is the "(issue) number" of a journal, magazine, or tech-report.
	Number string

	// Organization is used to identify the conference sponsor accociated with the cited material.
	Organization string

	// Pages is used to describe the pages from which the citation has been made.
	// Typically, the page numbers will be separated either by commas or by double-hyphens.
	Pages string

	// Publisher is the name of the publisher.
	Publisher string

	// School is used when the cited material is a thesis to indicate the the school where the thesis was written.
	School string

	// Series is used when the cited material was publushed as part of series of books.
	Series string

	// Title is the title of the cited material.
	Title string

	// Type is used to provide additional categorisation to the type of publication.
	Type string

	// Volume is the description of the volume that the cited material, typically a journal or multi-volume book, came from.
	Volume string

	// Year is the month of publication.
	// For unpublished material, the year of creation can be used instead.
	Year string
}
