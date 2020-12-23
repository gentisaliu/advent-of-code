package puzzle

import (
	"testing"
)

var (
	day1         = Day1{}
	testDataDay1 = []string{
		"1721",
		"979",
		"366",
		"299",
		"675",
		"1456",
	}
)

func TestDay1AnswerPart1(t *testing.T) {
	want := 514579
	got, _ := day1.AnswerPartOne(&testDataDay1)

	if got != want {
		t.Errorf("Got: %d. Want: %d", got, want)
	}
}

func TestDayAnswerPart2(t *testing.T) {
	want := 241861950
	got, _ := day1.AnswerPartTwo(&testDataDay1)

	if got != want {
		t.Errorf("Got: %d. Want: %d", got, want)
	}
}
