package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/urfave/cli/v3"
)

func newRootCmd() *cli.Command {
	return &cli.Command{
		Name:  "increment",
		Usage: "Determine version increment level based on commit message",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "commit", Aliases: []string{"c"}, Usage: "Commit message", Required: true},
			&cli.StringFlag{Name: "major", Aliases: []string{"m"}, Usage: "Major pattern", Required: true},
			&cli.StringFlag{Name: "minor", Aliases: []string{"n"}, Usage: "Minor pattern", Required: true},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			commitMessage := cmd.String("commit")
			majorPattern := cmd.String("major")
			minorPattern := cmd.String("minor")

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

			_, err = io.WriteString(cmd.Writer, incrementLevel)
			return err
		},
	}
}

func main() {
	rootCmd := newRootCmd()
	rootCmd.Writer = os.Stdout
	rootCmd.ErrWriter = os.Stderr

	if err := rootCmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}
}
