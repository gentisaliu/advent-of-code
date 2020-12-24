package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	day6         = Day6{}
	testDataDay6 = []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
		"",
	}
)

func TestDay6Part1(t *testing.T) {
	actual, _ := day6.AnswerPartOne(&testDataDay6)
	assert.Equal(t, 11, actual)
}

func TestDay6Part2(t *testing.T) {
	actual, _ := day6.AnswerPartTwo(&testDataDay6)
	assert.Equal(t, 6, actual)
}
