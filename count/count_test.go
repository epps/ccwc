package count

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

type testCase struct {
	name        string
	linesOption bool
	wordsOption bool
	bytesOption bool
	charsOption bool
	wcArgs      []string
}

const testFilepath = "../test.txt"

func TestCount(t *testing.T) {
	if _, err := os.Stat(testFilepath); err != nil {
		t.Fatalf("failed to get information on test file due to error: %v\nensure you have 'test.txt' at the root of the project", err)
	}
	testCases := []testCase{
		{
			name:        "Lines Option (-l)",
			linesOption: true,
			wcArgs:      []string{"wc", "-l", testFilepath},
		},
		{
			name:        "Words Option (-w)",
			wordsOption: true,
			wcArgs:      []string{"wc", "-w", testFilepath},
		},
		{
			name:        "Bytes Option (-c)",
			bytesOption: true,
			wcArgs:      []string{"wc", "-c", testFilepath},
		},
		{
			name:        "Characters Option (-m)",
			charsOption: true,
			wcArgs:      []string{"wc", "-m", testFilepath},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lines, words, bytes, chars, err := Count(testFilepath, tc.linesOption, tc.wordsOption, tc.bytesOption, tc.charsOption)
			if err != nil {
				t.Fatalf("failed to get actual value due to error: %v", err)
			}
			expected := getExpectedCount(t, tc.wcArgs)
			if tc.linesOption {
				if lines != expected[0] {
					t.Fatalf("expected %d but received %d", expected, lines)
				}
			}
			if tc.wordsOption {
				if words != expected[0] {
					t.Fatalf("expected %d but received %d", expected, words)
				}
			}
			if tc.bytesOption && !tc.charsOption {
				if bytes != expected[0] {
					t.Fatalf("expected %d but received %d", expected, bytes)
				}
			}
			if tc.charsOption {
				if chars != expected[0] {
					t.Fatalf("expected %d but received %d", expected, chars)
				}
			}
		})
	}
}

func getExpectedCount(t *testing.T, args []string) []int {
	expectedOutputBytes, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		t.Fatalf("failed to get expected output due to error: %v", err)
	}
	expectedOutput := strings.Trim(string(expectedOutputBytes), " ")
	parts := strings.Split(expectedOutput, " ")
	if len(parts) == 0 {
		t.Fatalf("expected output is empty")
	}
	outputsWithoutFile := parts[:len(parts)-1]
	expected := make([]int, len(outputsWithoutFile))
	for idx, o := range outputsWithoutFile {
		i, err := strconv.ParseInt(o, 0, 64)
		if err != nil {
			t.Fatalf("failed to parse %s to int due to error: %v", parts[0], err)
		}
		expected[idx] = int(i)
	}
	return expected
}
