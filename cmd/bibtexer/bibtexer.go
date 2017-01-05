package main

import (
	"fmt"
	"go/build"
	"go/parser"
	"go/types"
	"html/template"
	"log"

	"golang.org/x/tools/go/loader"
)

// Options describes the information needed to generate BibTeX serialization
// code for a type belonging to a specific package
type Options struct {
	TypeName   string
	ImportPath string
	Output     string
}

var tmpl, _ = template.New("").Parse(`// Created by bibtexer; DO NOT EDIT
package {{.PackageName}}

import (
{{range .Deps}}	"{{.}}"
{{end}})

func ({{.RecieverName}} *{{.TypeName}}) encode(writer io.Writer) (err error) {
	if len(entry.Address) > 0 {
		if _, err = fmt.Fprintf(writer, ",\n  address = {%s}", entry.Address); err != nil {
			return
		}
	}

// {{.InterfaceName}} is an interface generated for {{.Type}}.
type {{.InterfaceName}} interface {
{{range .Interface}}	{{.}}
{{end}}}
`)

func Run(o *Options) (err error) {
	lconf := loader.Config{
		Build:       &build.Default,
		AllowErrors: true,
		// AllErrors makes the parser always return an AST instead of
		// bailing out after 10 errors and returning an empty ast.File.
		ParserMode:          parser.AllErrors,
		TypeChecker:         types.Config{Error: func(err error) {}},
		TypeCheckFuncBodies: func(p string) bool { return p == o.ImportPath },
	}

	lconf.ImportWithTests(o.ImportPath)

	lprog, err := lconf.Load()
	if err != nil {
		return
	}

	pkg := lprog.Package(o.ImportPath).Pkg
	obj := pkg.Scope().Lookup(o.TypeName)
	if obj == nil {
		return fmt.Errorf("Unable to find %q in package %q", o.TypeName, pkg.Path())
	}

	tnObj, ok := obj.(*types.TypeName)
	if !ok {
		return fmt.Errorf("Entity %q in package %q is not a named type, got %T", o.TypeName, pkg.Path(), obj)
	}

	tStruct, ok := tnObj.Type().Underlying().(*types.Struct)
	if !ok {
		return fmt.Errorf("Named type %q in package %q is not a defining a struct, got %T", o.TypeName, pkg.Path(), tnObj.Type().Underlying())
	}

	for i := 0; i < tStruct.NumFields(); i++ {
		f := tStruct.Field(i)
		t := tStruct.Tag(i)
		if f.Anonymous() {
			log.Println("- recursion needed:", f)
		} else {
			log.Println("- field:", f)
		}
		if len(t) != 0 {
			log.Println("- tag:", t)
		}
	}

	log.Println("hello world")
	log.Println("- output:", o.Output)

	return
}
