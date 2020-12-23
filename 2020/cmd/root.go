package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "advent-of-code-2020",
	Short: "Provides answers to Advent of Code 2020 puzzles",
}

func init() {
	rootCmd.AddCommand(AnswerCmd())
}

// Execute executes the root command and handles any errors
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
