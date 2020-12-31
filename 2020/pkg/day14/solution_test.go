package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	solution      = Solution{}
	testData1 = []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}
	testData2 = []string {
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}
)

func TestPart1(t *testing.T) {
	actual, _ := solution.PartOne(&testData1)
	assert.Equal(t, 165, actual)
}

func TestPart2(t *testing.T) {
	actual, _ := solution.PartTwo(&testData2)
	assert.Equal(t, 208, actual)
}