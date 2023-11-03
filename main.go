package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	wc "github.com/epps/ccwc/count"
)

var bytesOption bool
var linesOption bool
var wordsOption bool
var charsOption bool

func init() {
	const (
		bytes      = "c"
		bytesUsage = "The number of bytes in each input file is written to the standard output."
		lines      = "l"
		linesUsage = "The number of lines in each input file is written to the standard output."
		words      = "w"
		wordsUsage = "The number of words in each input file is written to the standard output."
		chars      = "m"
		charsUsage = "The number of characters in each input file is written to the standard output."
	)

	// Set the log flags to 0 to avoid the timestamp.
	log.SetFlags(0)

	flag.BoolVar(&bytesOption, bytes, false, bytesUsage)
	flag.BoolVar(&linesOption, lines, false, linesUsage)
	flag.BoolVar(&wordsOption, words, false, wordsUsage)
	flag.BoolVar(&charsOption, chars, false, charsUsage)
}

func main() {
	flag.Parse()

	files := flag.Args()

	// Supports the default actions (i.e. when no flags are passed, the
	// count is run as if the -c -l and -w options were selected).
	if !linesOption && !wordsOption && !bytesOption && !charsOption {
		linesOption, wordsOption, bytesOption = true, true, true
	}

	// Cancels out the bytes option in the even both bytes and character
	// options are selected
	if charsOption {
		bytesOption = false
	}

	if len(files) == 0 {
		input, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("failed to read from stdin due to error: %v", err)
		}

		lines, words, bytes, chars, err := wc.CountFromBytes(input, linesOption, wordsOption, bytesOption, charsOption)
		if err != nil {
			log.Fatalf(err.Error())
		}
		LogCounts(linesOption, wordsOption, bytesOption, charsOption, lines, words, bytes, chars, "")
	}

	for _, f := range files {
		lines, words, bytes, chars, err := wc.CountFromFile(f, linesOption, wordsOption, bytesOption, charsOption)
		if err != nil {
			log.Fatalf(err.Error())
		}
		LogCounts(linesOption, wordsOption, bytesOption, charsOption, lines, words, bytes, chars, f)
	}
}

func LogCounts(l, w, b, c bool, lines, words, bytes, chars int, f string) {
	output := ""
	if l {
		output = fmt.Sprintf("%s%d\t", output, lines)
	}
	if w {
		output = fmt.Sprintf("%s%d\t", output, words)
	}
	if b {
		output = fmt.Sprintf("%s%d\t", output, bytes)
	}
	if c {
		output = fmt.Sprintf("%s%d\t", output, chars)
	}
	log.Printf("%s%s", output, f)
}
