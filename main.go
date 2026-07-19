package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
			logger := zerolog.New(cmd.ErrWriter).With().Timestamp().Str("command", cmd.Name).Logger()

			commitMessage := cmd.String("commit")
			majorPattern := cmd.String("major")
			minorPattern := cmd.String("minor")
			encodedMajorPattern := base64.StdEncoding.EncodeToString([]byte(majorPattern))
			encodedMinorPattern := base64.StdEncoding.EncodeToString([]byte(minorPattern))
			logger.Info().
				Str("commit_message", commitMessage).
				Str("major_pattern", encodedMajorPattern).
				Str("minor_pattern", encodedMinorPattern).
				Msg("processing increment request")

			incrementLevel := "patch"

			majorRegex, err := regexp.Compile(majorPattern)
			if err != nil {
				logger.Error().Err(err).Str("pattern", encodedMajorPattern).Msg("failed to compile major pattern")
				return fmt.Errorf("invalid major pattern: %w", err)
			}
			minorRegex, err := regexp.Compile(minorPattern)
			if err != nil {
				logger.Error().Err(err).Str("pattern", encodedMinorPattern).Msg("failed to compile minor pattern")
				return fmt.Errorf("invalid minor pattern: %w", err)
			}

			if majorRegex.MatchString(commitMessage) {
				incrementLevel = "major"
			} else if minorRegex.MatchString(commitMessage) {
				incrementLevel = "minor"
			}
			logger.Info().Str("increment_level", incrementLevel).Msg("calculated increment level")

			_, err = io.WriteString(cmd.Writer, incrementLevel)
			if err != nil {
				logger.Error().Err(err).Msg("failed to write increment level")
			}
			return err
		},
	}
}

func main() {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Str("app", "commit-increment").Logger()
	log.Info().Msg("initializing command")

	rootCmd := newRootCmd()
	rootCmd.Writer = os.Stdout
	rootCmd.ErrWriter = os.Stderr
	log.Info().Int("arg_count", len(os.Args)).Msg("running command")

	if err := rootCmd.Run(context.Background(), os.Args); err != nil {
		log.Error().Err(err).Msg("command execution failed")
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}

	log.Info().Msg("command completed successfully")
}
