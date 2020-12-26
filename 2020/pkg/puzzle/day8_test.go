package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	day8         = Day8{}
	testDataDay8 = []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}
)

func TestDay8Part1(t *testing.T) {
	actual, _ := day8.AnswerPartOne(&testDataDay8)
	assert.Equal(t, 5, actual)
}

func TestDay8Part2(t *testing.T) {
	actual, _ := day8.AnswerPartTwo(&testDataDay8)
	assert.Equal(t, 8, actual)
}
