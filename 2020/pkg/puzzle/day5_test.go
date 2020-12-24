package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	day5         = Day5{}
	testDataDay5 = []string{
		"FBFBBFFRLR",
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	}
)

func TestDay5Part1(t *testing.T) {
	actual, _ := day5.AnswerPartOne(&testDataDay5)
	assert.Equal(t, 820, actual)
}
