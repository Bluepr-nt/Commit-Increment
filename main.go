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
		Run: func(cmd *cobra.Command, args []string) {
			commitMessage, _ := cmd.Flags().GetString("commit")
			majorPattern, _ := cmd.Flags().GetString("major")
			minorPattern, _ := cmd.Flags().GetString("minor")

			incrementLevel := "patch"

			majorRegex, _ := regexp.Compile(majorPattern)
			minorRegex, _ := regexp.Compile(minorPattern)

			if majorRegex.MatchString(commitMessage) {
				incrementLevel = "major"
			} else if minorRegex.MatchString(commitMessage) {
				incrementLevel = "minor"
			}

			cmd.Print(incrementLevel)
		},
	}

	rootCmd.Flags().StringP("commit", "c", "", "Commit message")
	rootCmd.Flags().StringP("major", "m", "", "Major pattern")
	rootCmd.Flags().StringP("minor", "n", "", "Minor pattern")

	if err := rootCmd.MarkFlagRequired("commit"); err != nil {
		rootCmd.Println("Error marking 'commit' flag as required:", err)
		return nil, err
	}

	if err := rootCmd.MarkFlagRequired("major"); err != nil {
		rootCmd.Println("Error marking 'major' flag as required:", err)
		return nil, err
	}

	if err := rootCmd.MarkFlagRequired("minor"); err != nil {
		rootCmd.Println("Error marking 'minor' flag as required:", err)
		return nil, err
	}
	return rootCmd, nil
}

func main() {
	rootCmd, err := newRootCmd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		rootCmd.Println(err)
		os.Exit(1)
	}
}
