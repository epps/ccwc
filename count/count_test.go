package count

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

type testCase struct {
	name    string
	countFn func(file *os.File) (int, error)
	wcArgs  []string
}

const testFilepath = "../test.txt"

func TestCount(t *testing.T) {
	if _, err := os.Stat(testFilepath); err != nil {
		t.Fatalf("failed to get information on test file due to error: %v\nensure you have 'test.txt' at the root of the project", err)
	}
	file, err := os.Open(testFilepath)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			t.Fatalf("failed to close test file due to error: %v", err)
		}
	}(file)
	if err != nil {
		t.Fatalf("failed to open file due to error: %v", err)
	}

	testCases := []testCase{
		{
			name:    "Bytes Option (-c)",
			countFn: CountBytes,
			wcArgs:  []string{"wc", "-c", testFilepath},
		},
		{
			name:    "Lines Option (-l)",
			countFn: CountLines,
			wcArgs:  []string{"wc", "-l", testFilepath},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.countFn(file)
			if err != nil {
				t.Fatalf("failed to get actual value due to error: %v", err)
			}
			expected := getExpectedCount(t, tc.wcArgs)

			if actual != expected {
				t.Fatalf("expected %d but received %d", expected, actual)
			}
		})
	}
}

func getExpectedCount(t *testing.T, args []string) int {
	expectedOutputBytes, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		t.Fatalf("failed to get expected output due to error: %v", err)
	}
	expectedOutput := strings.Trim(string(expectedOutputBytes), " ")
	parts := strings.Split(expectedOutput, " ")
	if len(parts) == 0 {
		t.Fatalf("expected output is empty")
	}
	i, err := strconv.ParseInt(parts[0], 0, 64)
	if err != nil {
		t.Fatalf("failed to parse %s to int due to error: %v", parts[0], err)
	}
	return int(i)
}
