package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	solution = Solution{}
	testData = []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
)

func TestPart1(t *testing.T) {
	actual, _ := solution.PartOne(&testData)
	assert.Equal(t, 2, actual)
}

func TestPart2(t *testing.T) {
	actual, _ := solution.PartTwo(&testData)
	assert.Equal(t, 1, actual)
}
