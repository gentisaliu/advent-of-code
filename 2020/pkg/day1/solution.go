package day1

import (
	"fmt"

	"github.com/gentisaliu/advent-of-code/2020/internal/util"
)

const targetSum = 2020

// Solution solves the day 1 puzzle
type Solution struct{}

// PartOne answers part 1 of the day 1 puzzle
func (d *Solution) PartOne(input *[]string) (int, error) {
	numbers := util.ConvertStringsToNumbers(input)

	nr1, nr2, err := findTwoNumbersWithSum(numbers, targetSum)
	return nr1 * nr2, err
}

func findTwoNumbersWithSum(numbers *[]int, sum int) (int, int, error) {
	count := len(*numbers)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if (*numbers)[i]+(*numbers)[j] == sum {
				return (*numbers)[i], (*numbers)[j], nil
			}
		}
	}
	return 0, 0, fmt.Errorf("could not find two numbers with sum %d", sum)
}

// PartTwo answers part 2 of the day 1 puzzle
func (d *Solution) PartTwo(input *[]string) (int, error) {
	numbers := util.ConvertStringsToNumbers(input)
	nr1, nr2, nr3, err := findThreeNumbersWithSum(numbers, targetSum)
	return nr1 * nr2 * nr3, err
}

func findThreeNumbersWithSum(numbers *[]int, sum int) (int, int, int, error) {
	count := len(*numbers)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			for k := j + 1; k < count; k++ {
				if (*numbers)[i]+(*numbers)[j]+(*numbers)[k] == sum {
					return (*numbers)[i], (*numbers)[j], (*numbers)[k], nil
				}
			}
		}
	}
	return 0, 0, 0, fmt.Errorf("could not find three numbers with sum %d", sum)
}
