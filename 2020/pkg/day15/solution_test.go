package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	solution  = Solution{}
	testData1 = []string{
		"0,3,6",
	}
	testData2 = []string{
		"1,3,2",
	}
	testData3 = []string{
		"2,1,3",
	}
	testData4 = []string{
		"1,2,3",
	}
	testData5 = []string{
		"2,3,1",
	}
	testData6 = []string{
		"3,2,1",
	}
	testData7 = []string{
		"3,1,2",
	}
)

func TestPart1Test1(t *testing.T) {
	actual, _ := solution.PartOne(&testData1)
	assert.Equal(t, 436, actual)
}

func TestPart1Test2(t *testing.T) {
	actual, _ := solution.PartOne(&testData2)
	assert.Equal(t, 1, actual)
}

func TestPart1Test3(t *testing.T) {
	actual, _ := solution.PartOne(&testData3)
	assert.Equal(t, 10, actual)
}

func TestPart1Test4(t *testing.T) {
	actual, _ := solution.PartOne(&testData4)
	assert.Equal(t, 27, actual)
}

func TestPart1Test5(t *testing.T) {
	actual, _ := solution.PartOne(&testData5)
	assert.Equal(t, 78, actual)
}

func TestPart1Test6(t *testing.T) {
	actual, _ := solution.PartOne(&testData6)
	assert.Equal(t, 438, actual)
}

func TestPart1Test7(t *testing.T) {
	actual, _ := solution.PartOne(&testData7)
	assert.Equal(t, 1836, actual)
}

func TestPart2Test1(t *testing.T) {
	actual, _ := solution.PartTwo(&testData1)
	assert.Equal(t, 175594, actual)
}

func TestPart2Test2(t *testing.T) {
	actual, _ := solution.PartTwo(&testData2)
	assert.Equal(t, 2578, actual)
}

func TestPart2Test3(t *testing.T) {
	actual, _ := solution.PartTwo(&testData3)
	assert.Equal(t, 3544142, actual)
}

func TestPart2Test4(t *testing.T) {
	actual, _ := solution.PartTwo(&testData4)
	assert.Equal(t, 261214, actual)
}

func TestPart2Test5(t *testing.T) {
	actual, _ := solution.PartTwo(&testData5)
	assert.Equal(t, 6895259, actual)
}

func TestPart2Test6(t *testing.T) {
	actual, _ := solution.PartTwo(&testData6)
	assert.Equal(t, 18, actual)
}

func TestPart2Test7(t *testing.T) {
	actual, _ := solution.PartTwo(&testData7)
	assert.Equal(t, 362, actual)
}