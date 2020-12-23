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

func TestDay2AnswerPart1(t *testing.T) {
	expected, _ := day2.AnswerPartOne(&testDataDay2)
	assert.Equal(t, 2, expected)
}

func TestDay2AnswerPart2(t *testing.T) {
	expected, _ := day2.AnswerPartTwo(&testDataDay2)
	assert.Equal(t, 1, expected)
}
