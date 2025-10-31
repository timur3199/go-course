package uniq

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Options struct {
	Count      bool
	Repeated   bool
	Unique     bool
	IgnoreCase bool
	NumFields  int
	NumChars   int
}

type Line struct {
	Text    string
	Count   int
	Original string
}

func processLine(line string, options Options) string {
	result := line
	
	// Process -f: skip fields
	if options.NumFields > 0 {
		fields := strings.Fields(result)
		if len(fields) > options.NumFields {
			result = strings.Join(fields[options.NumFields:], " ")
		} else {
			result = ""
		}
	}
	
	// Process -s: skip chars
	if options.NumChars > 0 && len(result) > options.NumChars {
		result = result[options.NumChars:]
	}
	
	// Process -i: ignore case
	if options.IgnoreCase {
		result = strings.ToLower(result)
	}
	
	return result
}

func Run(input io.Reader, output io.Writer, options Options) error {
	scanner := bufio.NewScanner(input)
	lines := make([]Line, 0)
	processedLines := make(map[string]int)

	// Read and process all lines
	for scanner.Scan() {
		originalLine := scanner.Text()
		processedLine := processLine(originalLine, options)
		
		if idx, exists := processedLines[processedLine]; exists {
			lines[idx].Count++
		} else {
			lines = append(lines, Line{
				Text:     processedLine,
				Count:    1,
				Original: originalLine,
			})
			processedLines[processedLine] = len(lines) - 1
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Write output based on options
	for _, line := range lines {
		switch {
		case options.Count:
			fmt.Fprintf(output, "%d %s\n", line.Count, line.Original)
		case options.Repeated && line.Count > 1:
			fmt.Fprintln(output, line.Original)
		case options.Unique && line.Count == 1:
			fmt.Fprintln(output, line.Original)
		case !options.Count && !options.Repeated && !options.Unique:
			fmt.Fprintln(output, line.Original)
		}
	}

	return nil
}