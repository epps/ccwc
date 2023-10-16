package main

import (
	"bytes"
	"log"
	"os"
	"testing"
)

type testCase struct {
	name           string
	args           []string
	expectedOutput string
}

func TestMain(t *testing.T) {

	testCases := []testCase{
		{
			name:           "Bytes Option (-c)",
			args:           []string{"cmd", "-c", "test.txt"},
			expectedOutput: "342384 test.txt\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Args = tc.args

			var buf bytes.Buffer
			log.SetOutput(&buf)
			main()
			log.SetOutput(os.Stderr)

			if buf.String() != tc.expectedOutput {
				t.Errorf("Expected '%s', got %s", tc.expectedOutput, buf.String())
			}
		})
	}
}
