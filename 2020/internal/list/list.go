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
