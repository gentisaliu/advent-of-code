package puzzle

import "testing"

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
	want := 7
	got, _ := day3.AnswerPartOne(&testDataDay3)

	if got != want {
		t.Errorf("Got: %d. Want: %d", got, want)
	}
}

func TestDay3AnswerPart2(t *testing.T) {
	want := 336
	got, _ := day3.AnswerPartTwo(&testDataDay3)

	if got != want {
		t.Errorf("Got: %d. Want: %d", got, want)
	}
}
