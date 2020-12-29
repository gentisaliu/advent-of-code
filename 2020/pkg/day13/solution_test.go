package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	solution  = Solution{}
	testData1 = []string{
		"939",
		"7,13,x,x,59,x,31,19",
	}
	testData2 = []string{
		"",
		"17,x,13,19",
	}
	testData3 = []string{
		"",
		"67,7,59,61",
	}
	testData4 = []string{
		"",
		"67,x,7,59,61",
	}
	testData5 = []string{
		"",
		"67,7,x,59,61",
	}
	testData6 = []string{
		"",
		"1789,37,47,1889",
	}
)

func TestPart1(t *testing.T) {
	actual, _ := solution.PartOne(&testData1)
	assert.Equal(t, 295, actual)
}

func TestPart2Test1(t *testing.T) {
	actual, _ := solution.PartTwo(&testData1)
	assert.Equal(t, 1068781, actual)
}

func TestPart2Test2(t *testing.T) {
	actual, _ := solution.PartTwo(&testData2)
	assert.Equal(t, 3417, actual)
}

func TestPart2Test3(t *testing.T) {
	actual, _ := solution.PartTwo(&testData3)
	assert.Equal(t, 754018, actual)
}

func TestPart2Test4(t *testing.T) {
	actual, _ := solution.PartTwo(&testData4)
	assert.Equal(t, 779210, actual)
}

func TestPart2Test5(t *testing.T) {
	actual, _ := solution.PartTwo(&testData5)
	assert.Equal(t, 1261476, actual)
}

func TestPart2Test6(t *testing.T) {
	actual, _ := solution.PartTwo(&testData6)
	assert.Equal(t, 1202161486, actual)
}
