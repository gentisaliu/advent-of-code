package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	solution = Solution{}
	testData = []string{
		"FBFBBFFRLR",
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	}
)

func TestPart1(t *testing.T) {
	actual, _ := solution.PartOne(&testData)
	assert.Equal(t, 820, actual)
}
