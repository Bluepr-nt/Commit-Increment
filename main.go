package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

func newRootCmd() (*cobra.Command, error) {
	var rootCmd = &cobra.Command{
		Use:   "increment",
		Short: "Determine version increment level based on commit message",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SetOut(os.Stdout)
			cmd.SetErr(os.Stderr)
			commitMessage, _ := cmd.Flags().GetString("commit")
			majorPattern, _ := cmd.Flags().GetString("major")
			minorPattern, _ := cmd.Flags().GetString("minor")

			incrementLevel := "patch"

			majorRegex, err := regexp.Compile(majorPattern)
			if err != nil {
				return fmt.Errorf("invalid major pattern: %w", err)
			}
			minorRegex, err := regexp.Compile(minorPattern)
			if err != nil {
				return fmt.Errorf("invalid minor pattern: %w", err)
			}

			if majorRegex.MatchString(commitMessage) {
				incrementLevel = "major"
			} else if minorRegex.MatchString(commitMessage) {
				incrementLevel = "minor"
			}

			cmd.Print(incrementLevel)
			return nil
		},
	}

	rootCmd.Flags().StringP("commit", "c", "", "Commit message")
	rootCmd.Flags().StringP("major", "m", "", "Major pattern")
	rootCmd.Flags().StringP("minor", "n", "", "Minor pattern")

	if err := rootCmd.MarkFlagRequired("commit"); err != nil {
		rootCmd.PrintErrf("Error marking 'commit' flag as required: %v\n", err)
		return nil, err
	}

	if err := rootCmd.MarkFlagRequired("major"); err != nil {
		rootCmd.PrintErrf("Error marking 'major' flag as required: %v\n", err)
		return nil, err
	}

	if err := rootCmd.MarkFlagRequired("minor"); err != nil {
		rootCmd.PrintErrf("Error marking 'minor' flag as required: %v\n", err)
		return nil, err
	}
	return rootCmd, nil
}

func main() {
	rootCmd, err := newRootCmd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating command: %v\n", err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		rootCmd.PrintErrf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}
