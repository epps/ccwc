package count

import (
	"bufio"
	"os"
)

func CountBytes(file *os.File) (int, error) {
	fStat, err := file.Stat()
	if err != nil {
		return 0, nil
	}
	return int(fStat.Size()), nil
}

func CountLines(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)
	lineCount := 0
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lineCount++
	}
	return lineCount, nil
}

func CountWords(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)
	wordCount := 0
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordCount++
	}
	return wordCount, nil
}
