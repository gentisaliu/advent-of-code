package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	solution  = Solution{}
	testData1 = []string{
		"16",
		"10",
		"15",
		"5",
		"1",
		"11",
		"7",
		"19",
		"6",
		"12",
		"4",
	}
	testData2 = []string{
		"28",
		"33",
		"18",
		"42",
		"31",
		"14",
		"46",
		"20",
		"48",
		"47",
		"24",
		"23",
		"49",
		"45",
		"19",
		"38",
		"39",
		"11",
		"1",
		"32",
		"25",
		"35",
		"8",
		"17",
		"7",
		"9",
		"4",
		"2",
		"34",
		"10",
		"3",
	}
)

func TestPart1Test1(t *testing.T) {
	actual, _ := solution.PartOne(&testData1)
	assert.Equal(t, 35, actual)
}

func TestPart1Test2(t *testing.T) {
	actual, _ := solution.PartOne(&testData2)
	assert.Equal(t, 220, actual)
}

func TestPart2Test1(t *testing.T) {
	actual, _ := solution.PartTwo(&testData1)
	assert.Equal(t, 8, actual)
}

func TestPart2Test2(t *testing.T) {
	actual, _ := solution.PartTwo(&testData2)
	assert.Equal(t, 19208, actual)
}
