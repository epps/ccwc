package count

import (
	"bufio"
	"os"
	"strings"
)

func Count(filename string, linesOption, wordsOption, bytesOption, charsOption bool) (lines, words, bytes, chars int, err error) {
	file, err := os.Open(filename)
	if err != nil {
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
