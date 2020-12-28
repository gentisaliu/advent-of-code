package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	solution = Solution{}
	testData = []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}
	testDataOccupancyCount1 = []string{
		".......#.",
		"...#.....",
		".#.......",
		".........",
		"..#L....#",
		"....#....",
		".........",
		"#........",
		"...#.....",
	}
	testDataOccupancyCount2 = []string{
		".............",
		".L.L.#.#.#.#.",
		".............",
	}
	testDataOccupancyCount3 = []string{
		".##.##.",
		"#.#.#.#",
		"##...##",
		"...L...",
		"##...##",
		"#.#.#.#",
		".##.##.",
	}
)

func TestPart1(t *testing.T) {
	actual, _ := solution.PartOne(&testData)
	assert.Equal(t, 37, actual)
}

func TestPart2(t *testing.T) {
	actual, _ := solution.PartTwo(&testData)
	assert.Equal(t, 26, actual)
}

func TestCountVisibleOccupancy1(t *testing.T) {
	seats := seatLayout{&testDataOccupancyCount1}
	actual := seats.countVisibleOccupancy(3, 4)
	assert.Equal(t, 8, actual)
}

func TestCountVisibleOccupancy2(t *testing.T) {
	seats := seatLayout{&testDataOccupancyCount2}
	actual := seats.countVisibleOccupancy(1, 1)
	assert.Equal(t, 0, actual)
}

func TestCountVisibleOccupancy3(t *testing.T) {
	seats := seatLayout{&testDataOccupancyCount3}
	actual := seats.countVisibleOccupancy(3, 3)
	assert.Equal(t, 0, actual)
}
