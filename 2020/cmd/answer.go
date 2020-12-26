package cmd

import (
	"fmt"

	"github.com/gentisaliu/advent-of-code/2020/internal/fs"
	"github.com/gentisaliu/advent-of-code/2020/pkg/puzzle"
	"github.com/spf13/cobra"
)

var puzzles = map[int]puzzle.Puzzle{
	1: &puzzle.Day1{},
	2: &puzzle.Day2{},
	3: &puzzle.Day3{},
	4: &puzzle.Day4{},
	5: &puzzle.Day5{},
	6: &puzzle.Day6{},
	7: &puzzle.Day7{},
	8: &puzzle.Day8{},
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
	fmt.Printf("Day %v puzzle, input file %v\n", day, inputFilePath)

	lines, err := fs.ReadAllLines(inputFilePath)
	if err != nil {
		return fmt.Errorf("error reading input file: %v", err)
	}

	puzzle, dayFound := puzzles[day]
	if !dayFound {
		return fmt.Errorf("no puzzle for day %v found", day)
	}

	answer1, err1 := puzzle.AnswerPartOne(&lines)
	printAnswer("Answer Part 1: ", answer1, err1)

	answer2, err2 := puzzle.AnswerPartTwo(&lines)
	printAnswer("Answer Part 2: ", answer2, err2)

	return nil
}

func printAnswer(label string, answer int, err error) {
	fmt.Print(label)
	if err != nil {
		fmt.Printf("could not calculate answer: %v\n", err)
	} else {
		fmt.Printf("%d\n", answer)
	}
}
