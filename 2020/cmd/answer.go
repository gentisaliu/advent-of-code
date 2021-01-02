package cmd

import (
	"fmt"
	"time"

	"github.com/gentisaliu/advent-of-code/2020/internal/util"
	"github.com/gentisaliu/advent-of-code/2020/pkg/day"
	"github.com/gentisaliu/advent-of-code/2020/pkg/day1"
	"github.com/gentisaliu/advent-of-code/2020/pkg/day10"
	"github.com/gentisaliu/advent-of-code/2020/pkg/day11"
	"github.com/gentisaliu/advent-of-code/2020/pkg/day12"
	"github.com/gentisaliu/advent-of-code/2020/pkg/day13"
	"github.com/gentisaliu/advent-of-code/2020/pkg/day14"
	"github.com/gentisaliu/advent-of-code/2020/pkg/day15"
	"github.com/gentisaliu/advent-of-code/2020/pkg/day2"
	"github.com/gentisaliu/advent-of-code/2020/pkg/day3"
	"github.com/gentisaliu/advent-of-code/2020/pkg/day4"
	"github.com/gentisaliu/advent-of-code/2020/pkg/day5"
	"github.com/gentisaliu/advent-of-code/2020/pkg/day6"
	"github.com/gentisaliu/advent-of-code/2020/pkg/day7"
	"github.com/gentisaliu/advent-of-code/2020/pkg/day8"
	"github.com/gentisaliu/advent-of-code/2020/pkg/day9"
	"github.com/spf13/cobra"
)

var solutions = map[int]day.Solution{
	1:  &day1.Solution{},
	2:  &day2.Solution{},
	3:  &day3.Solution{},
	4:  &day4.Solution{},
	5:  &day5.Solution{},
	6:  &day6.Solution{},
	7:  &day7.Solution{},
	8:  &day8.Solution{},
	9:  &day9.Solution{},
	10: &day10.Solution{},
	11: &day11.Solution{},
	12: &day12.Solution{},
	13: &day13.Solution{},
	14: &day14.Solution{},
	15: &day15.Solution{},
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

	input, err := util.ReadAllLines(inputFilePath)
	if err != nil {
		return fmt.Errorf("error reading input file: %v", err)
	}

	solution, dayFound := solutions[day]
	if !dayFound {
		return fmt.Errorf("no solution for day %v found", day)
	}

	err = partOne(&solution, &input)
	err = partTwo(&solution, &input)

	return nil
}

func partOne(solution *day.Solution, input *[]string) error {
	start := time.Now()
	answer, err := (*solution).PartOne(input)
	elapsed := time.Since(start)

	printAnswer(1, answer, err, elapsed)

	return err
}

func partTwo(solution *day.Solution, input *[]string) error {
	start := time.Now()
	answer, err := (*solution).PartTwo(input)
	elapsed := time.Since(start)

	printAnswer(2, answer, err, elapsed)

	return err
}

func printAnswer(part int, answer int, err error, duration time.Duration) {
	switch {
	case err != nil:
		fmt.Printf("error calculating answer: %v\n", err)
	default:
		fmt.Printf("Answer part %d: %d (duration: %v)\n", part, answer, duration)
	}
}
