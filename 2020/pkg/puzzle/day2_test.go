package puzzle

import "testing"

var (
	day2         = Day2{}
	testDataDay2 = []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
)

func TestDay2AnswerPart1(t *testing.T) {
	want := 2
	got, _ := day2.AnswerPartOne(&testDataDay2)

	if got != want {
		t.Errorf("Got: %d. Want: %d", got, want)
	}
}

func TestDay2AnswerPart2(t *testing.T) {
	want := 1
	got, _ := day2.AnswerPartTwo(&testDataDay2)

	if got != want {
		t.Errorf("Got: %d. Want: %d", got, want)
	}
}
