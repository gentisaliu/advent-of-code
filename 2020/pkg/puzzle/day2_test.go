package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	day2         = Day2{}
	testDataDay2 = []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
)

func TestDay2Part1(t *testing.T) {
	actual, _ := day2.AnswerPartOne(&testDataDay2)
	assert.Equal(t, 2, actual)
}

func TestDay2Part2(t *testing.T) {
	actual, _ := day2.AnswerPartTwo(&testDataDay2)
	assert.Equal(t, 1, actual)
}
