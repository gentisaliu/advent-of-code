package list

import (
	"strconv"
)

// ConvertStringSliceToNumberSlice converts a slice of strings to a slice of ints
func ConvertStringSliceToNumberSlice(aSlice *[]string) *[]int {
	var numbers []int
	for _, v := range *aSlice {
		number, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}
	return &numbers
}

// GetMinMax finds the minimum and maximum in a list of integers
func GetMinMax(numbers *[]int) (int, int) {
	min, max := (*numbers)[0], (*numbers)[0]
	for _, number := range *numbers {
		if number > max {
			max = number
		} else if number < min {
			min = number
		}
	}
	return min, max
}

// GetIndexOfFirstNumberOccurrence returns the index of the first occurrence of the given integer in the list
func GetIndexOfFirstNumberOccurrence(numbers *[]int, el int) int {
	for i := range *numbers {
		if (*numbers)[i] == el {
			return i
		}
	}
	return -1
}
