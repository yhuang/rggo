package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	options := make(map[string]bool)

	// Defining a boolean flag -b to count bytes instead of words
	bytesPtr := flag.Bool("b", false, "Count bytes")

	// Defining a boolean flag -l to count lines instead of words
	linesPtr := flag.Bool("l", false, "Count lines")

	// Parsing the flags provided by the user
	flag.Parse()

	options["bytes"] = *bytesPtr
	options["lines"] = *linesPtr

	// Calling the count function to count the number of words (or lines)
	// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin, options))
}

func count(r io.Reader, options map[string]bool) int {
	// A scanner is used to read text from the a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// If the count bytes flag is set, we want to count bytes
	if options["bytes"] {
		scanner.Split(bufio.ScanBytes)
		// If the count lines flag is set, we want to count lines
	} else if options["lines"] {
		scanner.Split(bufio.ScanLines)
		// If no flag is set, we count words
	} else {
		scanner.Split(bufio.ScanWords)
	}

	// Defining a counter
	wc := 0

	// For every byte, line, or word scanned, add 1 to the counter
	for scanner.Scan() {
		wc++
	}

	// Return the total
	return wc
}
