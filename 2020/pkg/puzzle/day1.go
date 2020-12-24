package puzzle

import (
	"fmt"
	"strconv"
)

const targetSum = 2020

// Day1 implements the day 1 puzzle
type Day1 struct{}

// AnswerPartOne answers part 1 of the day 1 puzzle
func (d *Day1) AnswerPartOne(input *[]string) (int, error) {
	numbers := parseNumbersFromInput(input)
	nr1, nr2, err := findTwoNumbersWithSum(&numbers, targetSum)
	return nr1 * nr2, err
}

func parseNumbersFromInput(aSlice *[]string) []int {
	var numbers []int
	for _, v := range *aSlice {
		number, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
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

// AnswerPartTwo answers part 2 of the day 1 puzzle
func (d *Day1) AnswerPartTwo(input *[]string) (int, error) {
	numbers := parseNumbersFromInput(input)
	nr1, nr2, nr3, err := findThreeNumbersWithSum(&numbers, targetSum)
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
