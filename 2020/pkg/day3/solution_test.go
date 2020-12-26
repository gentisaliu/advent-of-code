package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	solution = Solution{}
	testData = []string{
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

func TestPart1(t *testing.T) {
	actual, _ := solution.PartOne(&testData)
	assert.Equal(t, 7, actual)
}

func TestPart2(t *testing.T) {
	actual, _ := solution.PartTwo(&testData)
	assert.Equal(t, 336, actual)
}
