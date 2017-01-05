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
* It provides to tools to provide BibTeX encoding for custom interfaces using compile-time code generation

---

* [Install](#install)
* [Tools](#tools)
  * [cmd/bibtexer](#cmdbibtexer-)
* [License](./LICENSE)

---

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get github.com/matthewdunsdon/bibliography
```

## Tools

### cmd/bibtexer [![GoDoc](https://godoc.org/github.com/matthewdunsdon/bibliography/cmd/bibtexer?status.png)](https://godoc.org/github.com/matthewdunsdon/bibliography/cmd/bibtexer)

Generates an interface for a named type.

*Usage*

```bash
~ $ interfacer -help
```
```
Usage of interfacer:
  -all
        Include also unexported methods.
  -as string
        Generated interface name. (default "main.Interface")
  -for string
        Type to generate an interface for.
  -o string
        Output file. (default "-")
```

*Example*

```bash
~ $ interfacer -for \"os\".File -as mock.File
```
```go
// Created by interfacer; DO NOT EDIT

package mock

import (
        "os"
)

// File is an interface generated for "os".File.
type File interface {
        Chdir() error
        Chmod(os.FileMode) error
        Chown(int, int) error
        Close() error
        Fd() uintptr
        Name() string
        Read([]byte) (int, error)
        ReadAt([]byte, int64) (int, error)
        Readdir(int) ([]os.FileInfo, error)
        Readdirnames(int) ([]string, error)
        Seek(int64, int) (int64, error)
        Stat() (os.FileInfo, error)
        Sync() error
        Truncate(int64) error
        Write([]byte) (int, error)
        WriteAt([]byte, int64) (int, error)
        WriteString(string) (int, error)
}
```

## License

MIT licensed. See the LICENSE file for details.
