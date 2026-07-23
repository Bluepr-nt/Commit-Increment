package main

import (
	"bytes"
	"context"
	"io"
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
			name:           "Major increment with regex pattern",
			commitMessage:  "feat!: introduce a breaking change",
			majorPattern:   ".*!.*",
			minorPattern:   "^feat.*",
			expectedOutput: "major",
		},
		{
			name:           "Minor increment with regex pattern",
			commitMessage:  "feat: introduce a new feature",
			majorPattern:   ".*!.*",
			minorPattern:   "^feat.*",
			expectedOutput: "minor",
		},
		{
			name:           "Patch increment with regex pattern",
			commitMessage:  "fix: fix a bug",
			majorPattern:   ".*!.*",
			minorPattern:   "^feat.*",
			expectedOutput: "patch",
		},
		{
			name:           "Minor increment with multiline commit message",
			commitMessage:  "feat: introduce a new feature\n\nadditional context for the change",
			majorPattern:   ".*!:.*|BREAKING CHANGE:",
			minorPattern:   "^feat(\\(.+\\))?: .*$",
			expectedOutput: "minor",
		},
		{
			name:           "Major increment with breaking change in body",
			commitMessage:  "feat: introduce a new feature\n\nBREAKING CHANGE: incompatible API update",
			majorPattern:   "BREAKING CHANGE:",
			minorPattern:   "^feat(\\(.+\\))?: .*$",
			expectedOutput: "major",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rootCmd := newRootCmd()

			args := []string{
				"increment",
				"--commit", tc.commitMessage,
				"--major", tc.majorPattern,
				"--minor", tc.minorPattern,
			}

			buffer := new(bytes.Buffer)
			rootCmd.Writer = buffer
			rootCmd.ErrWriter = io.Discard

			if err := rootCmd.Run(context.Background(), args); err != nil {
				t.Fatal(err)
			}

			result := buffer.String()
			if result != tc.expectedOutput {
				t.Fatalf("expected \"%s\" but got \"%s\"", tc.expectedOutput, result)
			}
		})
	}
}

func TestInvalidRegexPatterns(t *testing.T) {
	testCases := []struct {
		name         string
		majorPattern string
		minorPattern string
	}{
		{
			name:         "Invalid major pattern",
			majorPattern: "[invalid",
			minorPattern: "valid",
		},
		{
			name:         "Invalid minor pattern",
			majorPattern: "valid",
			minorPattern: "[invalid",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rootCmd := newRootCmd()

			args := []string{
				"increment",
				"--commit", "some commit message",
				"--major", tc.majorPattern,
				"--minor", tc.minorPattern,
			}

			buffer := new(bytes.Buffer)
			rootCmd.Writer = buffer
			rootCmd.ErrWriter = io.Discard

			if err := rootCmd.Run(context.Background(), args); err == nil {
				t.Fatal("expected error for invalid regex pattern, but got none")
			}
		})
	}
}
