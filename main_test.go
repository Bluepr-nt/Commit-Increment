package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	testCases := []struct {
		name           string
		commitMessage  string
		majorPattern   string
		minorPattern   string
		expectedOutput string
	}{
		{
			name:           "Test case for major increment",
			commitMessage:  "This commit introduces a breaking change",
			majorPattern:   "breaking change",
			minorPattern:   "new feature",
			expectedOutput: "major",
		},
		{
			name:           "Test case for minor increment",
			commitMessage:  "This commit introduces a new feature",
			majorPattern:   "breaking change",
			minorPattern:   "new feature",
			expectedOutput: "minor",
		},
		{
			name:           "Test case for patch increment",
			commitMessage:  "This commit fixes a bug",
			majorPattern:   "breaking change",
			minorPattern:   "new feature",
			expectedOutput: "patch",
		},
		{
			name:           "Test case for major increment",
			commitMessage:  "feat!: introduce a breaking change",
			majorPattern:   ".*!.*",
			minorPattern:   "^feat.*",
			expectedOutput: "major",
		},
		{
			name:           "Test case for minor increment",
			commitMessage:  "feat: introduce a new feature",
			majorPattern:   ".*!.*",
			minorPattern:   "^feat.*",
			expectedOutput: "minor",
		},
		{
			name:           "Test case for patch increment",
			commitMessage:  "fix: fix a bug",
			majorPattern:   ".*!.*",
			minorPattern:   "^feat.*",
			expectedOutput: "patch",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rootCmd, err := newRootCmd()
			if err != nil {
				t.Fatal(err)
			}

			rootCmd.SetArgs([]string{
				"--commit", tc.commitMessage,
				"--major", tc.majorPattern,
				"--minor", tc.minorPattern,
			})

			buffer := new(bytes.Buffer)
			rootCmd.SetOut(buffer)
			rootCmd.SetErr(buffer)

			if err := rootCmd.Execute(); err != nil {
				t.Fatal(err)
			}

			result := buffer.String()
			if result != tc.expectedOutput {
				t.Fatalf("expected \"%s\" but got \"%s\"", tc.expectedOutput, result)
			}
		})
	}
}
