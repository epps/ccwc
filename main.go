package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var bytesOption bool

func init() {
	const (
		bytes      = "c"
		bytesUsage = "The number of bytes in each input file is written to the standard output."
	)

	// Set the log flags to 0 to avoid the timestamp.
	log.SetFlags(0)

	flag.BoolVar(&bytesOption, bytes, false, bytesUsage)
}

func main() {
	flag.Parse()

	files := flag.Args()

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if bytesOption {
			log.Printf("%v %v\n", len(data), file)
		}
	}
}
