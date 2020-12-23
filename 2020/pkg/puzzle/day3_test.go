package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	day3         = Day3{}
	testDataDay3 = []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
)

func TestDay3AnswerPart1(t *testing.T) {
	actual, _ := day3.AnswerPartOne(&testDataDay3)
	assert.Equal(t, 7, actual)
}

func TestDay3AnswerPart2(t *testing.T) {
	actual, _ := day3.AnswerPartTwo(&testDataDay3)
	assert.Equal(t, 336, actual)
}
