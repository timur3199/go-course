package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"uniq/internal/uniq"
)

func main() {
	var options uniq.Options

	flag.BoolVar(&options.Count, "c", false, "count the number of occurrences")
	flag.BoolVar(&options.Repeated, "d", false, "only output duplicate lines")
	flag.BoolVar(&options.Unique, "u", false, "only output unique lines")
	flag.BoolVar(&options.IgnoreCase, "i", false, "ignore case")
	flag.IntVar(&options.NumFields, "f", 0, "ignore the first num fields")
	flag.IntVar(&options.NumChars, "s", 0, "ignore the first num chars")

	flag.Parse()

	// Validate mutually exclusive flags
	if (options.Count && options.Repeated) || 
	   (options.Count && options.Unique) || 
	   (options.Repeated && options.Unique) {
		fmt.Fprintln(os.Stderr, "uniq: options -c, -d, -u are mutually exclusive")
		os.Exit(1)
	}

	// Get input and output files
	var input io.Reader = os.Stdin
	var output io.Writer = os.Stdout

	args := flag.Args()
	
	if len(args) > 0 {
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "uniq: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	}

	if len(args) > 1 {
		file, err := os.Create(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "uniq: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	}

	// Run uniq
	if err := uniq.Run(input, output, options); err != nil {
		fmt.Fprintf(os.Stderr, "uniq: %v\n", err)
		os.Exit(1)
	}
}