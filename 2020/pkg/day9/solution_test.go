package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	solution = Solution{}
	testData = []string{
		"35",
		"20",
		"15",
		"25",
		"47",
		"40",
		"62",
		"55",
		"65",
		"95",
		"102",
		"117",
		"150",
		"182",
		"127",
		"219",
		"299",
		"277",
		"309",
		"576",
	}
)

func TestPart1(t *testing.T) {
	preambleLength = 5
	actual, _ := solution.PartOne(&testData)
	assert.Equal(t, 127, actual)
}

func TestPart2(t *testing.T) {
	preambleLength = 5
	actual, _ := solution.PartTwo(&testData)
	assert.Equal(t, 62, actual)
}
