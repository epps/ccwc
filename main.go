package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	wc "github.com/epps/ccwc/count"
)

var bytesOption bool
var linesOption bool
var wordsOption bool

func init() {
	const (
		bytes      = "c"
		bytesUsage = "The number of bytes in each input file is written to the standard output."
		lines      = "l"
		linesUsage = "The number of lines in each input file is written to the standard output."
		words      = "w"
		wordsUsage = "The number of words in each input file is written to the standard output."
	)

	// Set the log flags to 0 to avoid the timestamp.
	log.SetFlags(0)

	flag.BoolVar(&bytesOption, bytes, false, bytesUsage)
	flag.BoolVar(&linesOption, lines, false, linesUsage)
	flag.BoolVar(&wordsOption, words, false, wordsUsage)
}

func main() {
	flag.Parse()

	files := flag.Args()

	for _, f := range files {
		file, err := os.Open(f)
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Fatalf("failed to close file due to error: %v", err)
			}
		}(file)
		if err != nil {
			log.Fatalf("failed to open file due to error: %v", err)
		}
		output := ""
		if linesOption {
			lines, err := wc.CountLines(file)
			if err != nil {
				log.Fatalf("failed to count lines in file %s due to error: %v", f, err)
			}
			output = fmt.Sprintf("%s%d\t", output, lines)
		}
		if wordsOption {
			words, err := wc.CountWords(file)
			if err != nil {
				log.Fatalf("failed to count words in file %s due to error: %v", f, err)
			}
			output = fmt.Sprintf("%s%d\t", output, words)
		}
		if bytesOption {
			size, err := wc.CountBytes(file)
			if err != nil {
				log.Fatalf("failed to count bytes in file %s due to error: %v", f, err)
			}
			output = fmt.Sprintf("%s%d\t", output, size)
		}
		log.Printf("%s%s", output, f)
	}
}
