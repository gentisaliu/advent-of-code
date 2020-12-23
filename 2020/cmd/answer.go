package cmd

import (
	"fmt"

	"github.com/gentisaliu/advent-of-code/2020/internal/fs"
	"github.com/gentisaliu/advent-of-code/2020/pkg/puzzle"
	"github.com/spf13/cobra"
)

var puzzles = map[int]puzzle.Puzzle{
	1: puzzle.Day1{},
}

// AnswerCmd
func AnswerCmd() *cobra.Command {
	var day int
	var inputFilePath string

	cmd := &cobra.Command{
		Use:   "answer -day=DAY -input=INPUT",
		Short: "Calculate answer of the puzzle for the specified day and input",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := answer(day, inputFilePath)
			return err
		},
	}
	cmd.Flags().IntVarP(&day, "day", "d", 1, "day for which to get the puzzle answers")
	cmd.Flags().StringVarP(&inputFilePath, "input", "i", "", "input file for the puzzle")
	cmd.MarkFlagRequired("day")
	cmd.MarkFlagRequired("input")

	return cmd
}

func answer(day int, inputFilePath string) error {
	fmt.Printf("Day %v puzzle, input file %v", day, inputFilePath)
	fmt.Println("Computing answers for part 1 and 2.")

	lines, err := fs.ReadAllLines(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed reading lines from input file: %v", err)
	}

	puzzle, puzzleFound := puzzles[day]
	if !puzzleFound {
		return fmt.Errorf("no puzzle for day %v found", day)
	}
	ansP1, ansP2 := puzzle.AnswerPartOne(&lines), puzzle.AnswerPartTwo(&lines)

	fmt.Printf("Answers: Part 1: %v, Part 2: %v\n", ansP1, ansP2)

	return nil
}
