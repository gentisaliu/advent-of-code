package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	day1         = Day1{}
	testDataDay1 = []string{
		"1721",
		"979",
		"366",
		"299",
		"675",
		"1456",
	}
)

func TestDay1AnswerPart1(t *testing.T) {
	expected, _ := day1.AnswerPartOne(&testDataDay1)
	assert.Equal(t, 514579, expected)
}

func TestDayAnswerPart2(t *testing.T) {
	expected, _ := day1.AnswerPartTwo(&testDataDay1)
	assert.Equal(t, 241861950, expected)
}
