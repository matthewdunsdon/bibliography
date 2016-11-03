// Package bibliography defines interfaces shared by other packages that convert
// data to and from byte-level and textual representations
package bibliography

// An Entry represents the bubliography
type Entry struct {

	// Address specifies the partial (typically city) or full address of the publisher
	Address string

	// The name(s) of the author(s) (in the case of more than one author, separated by and)
	Author string

	// The title of the book, if only part of it is being cited
	BookTitle string

	// The chapter number
	Chapter string

	// The edition of a book, long form (such as "First" or "Second")
	Edition string

	// The name(s) of the editor(s)
	Editor string

	// How it was published, if the publishing method is nonstandard
	HowPublished string

	// The institution that was involved in the publishing, but not necessarily the publisher
	Institution string

	// The journal or magazine the work was published in
	Journal string

	// A hidden field used for specifying or overriding the alphabetical order of entries (when the "author" and "editor" fields are missing). Note that this is very different from the key (mentioned just after this list) that is used to cite or cross-reference the entry.
	Key string

	// The month of publication (or, if unpublished, the month of creation)
	Month string

	// Miscellaneous extra information
	Note string

	// The "(issue) number" of a journal, magazine, or tech-report, if applicable. (Most publications have a "volume", but no "number" field.)
	Number string

	// The conference sponsor
	Organization string

	// Page numbers, separated either by commas or double-hyphens.
	Pages string

	// The publisher's name
	Publisher string

	// The school where the thesis was written
	School string

	// The series of books the book was published in (e.g. "The Hardy Boys" or "Lecture Notes in Computer Science")
	Series string

	// The title of the work
	Title string

	// The field overriding the default type of publication (e.g. "Research Note" for techreport, "{PhD} dissertation" for phdthesis, "Section" for inbook/incollection)
	Type string

	// The volume of a journal or multi-volume book
	Volume string

	// The year of publication (or, if unpublished, the year of creation)
	Year string
}
