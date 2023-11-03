package count

import (
	"bufio"
	b "bytes"
	"log"
	"os"
	"strings"
)

func CountFromFile(filename string, linesOption, wordsOption, bytesOption, charsOption bool) (lines, words, bytes, chars int, err error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file %s due to error: %v", filename, err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(file)
	if bytesOption {
		fStat, err := file.Stat()
		if err != nil {
			return lines, words, bytes, chars, err
		}
		bytes = int(fStat.Size())
	}
	if linesOption || wordsOption {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lines += 1
			line := scanner.Text()
			wordsInLine := strings.Fields(line)
			words += len(wordsInLine)
		}
	}
	if charsOption {
		file.Seek(0, 0)
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			chars += 1
		}
	}
	return
}

func CountFromBytes(input []byte, linesOption, wordsOption, bytesOption, charsOption bool) (lines, words, bytes, chars int, err error) {
	if bytesOption {
		bytes = len(input)
	}
	r := b.NewReader(input)
	if linesOption || wordsOption {
		scanner := bufio.NewScanner(r)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lines += 1
			line := scanner.Text()
			wordsInLine := strings.Fields(line)
			words += len(wordsInLine)
		}
	}
	if charsOption {
		r.Seek(0, 0)
		scanner := bufio.NewScanner(r)
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			chars += 1
		}
	}
	return
}
