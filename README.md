# matthewdunsdon/bibliography
[![GoDoc](https://godoc.org/github.com/matthewdunsdon/bibliography?status.svg)](https://godoc.org/github.com/matthewdunsdon/bibliography)
[![Build Status](https://travis-ci.org/matthewdunsdon/bibliography.svg?branch=master)](https://travis-ci.org/matthewdunsdon/bibliography)
[![Coverage Status](https://coveralls.io/repos/github/matthewdunsdon/bibliography/badge.svg?branch=master)](https://coveralls.io/github/matthewdunsdon/bibliography?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/matthewdunsdon/bibliography)](https://goreportcard.com/report/github.com/matthewdunsdon/bibliography)

Package `matthewdunsdon/bibliography` implements bibliography encoders and
decoders for common textual formats.

The main features are:

* It defines a `bibliography.Entry` type with the common [BibTeX field types](https://en.wikipedia.org/wiki/BibTeX#Field_types)
* It provides basic BibTeX encoding for `bibliography.Entry` and any supplied additional fields

---

* [Install](#install)
* [License](./LICENSE)

---

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get github.com/matthewdunsdon/bibliography
```

## License

MIT licensed. See the LICENSE file for details.
