package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// flags
var (
	packageFlag = flag.String("package", "", "package containing type")
	typeFlag    = flag.String("type", "", "Type to generate an BibTeX serialization for")
	outputFlag  = flag.String("output", "", "output file name; default packagedir/<type>_string.go")
)

const (
	newLine     = "\n"
	useHelp     = "Run 'bibtexer -help' for more information.\n"
	helpMessage = `Generates Go code for optimized BibTeX serialization.
Usage: bibtexer [flags]`
)

func printHelp() {
	fmt.Fprintln(os.Stderr, helpMessage)
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
}

var extRe = regexp.MustCompile(`(.*)(\.go)$`)

func main() {
	log.SetPrefix("bibtexer: ")
	log.SetFlags(0)

	// Don't print full help unless -help was requested.
	// Just gently remind users that it's there.
	flag.Usage = func() { fmt.Fprint(os.Stderr, newLine, useHelp) }
	flag.CommandLine.Init(os.Args[0], flag.ContinueOnError) // hack
	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		// (err has already been printed)
		if err == flag.ErrHelp {
			printHelp()
		}
		os.Exit(2)
	}

	packageEmpty, typeEmpty, outputEmpty := len(*packageFlag) == 0, len(*typeFlag) == 0, len(*outputFlag) != 0
	if packageEmpty || typeEmpty {
		if packageEmpty {
			fmt.Fprintln(os.Stderr, "empty -package flag value; must be set")
		}
		if typeEmpty {
			fmt.Fprintln(os.Stderr, "empty -type flag value; must be set")
		}
		flag.Usage()
		os.Exit(2)
	}

	var outputPath string
	if outputEmpty {
		outputPath = *outputFlag
	} else {
		outputPath = fmt.Sprintf("%s_bibtexer.go", strings.ToLower(*typeFlag))
	}

	options := Options{
		TypeName:   *typeFlag,
		ImportPath: *packageFlag,
		Output:     outputPath,
	}

	if err := Run(&options); err != nil {
		log.Fatal(err)
	}
}
