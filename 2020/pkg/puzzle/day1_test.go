package puzzle

import (
	"testing"
)

var (
	testdata = []string{
		"1721",
		"979",
		"366",
		"299",
		"675",
		"1456",
	}
	day1 = Day1{}
)

func TestAnswerPartOne(t *testing.T) {
	want := 514579
	got, _ := day1.AnswerPartOne(&testdata)

	if got != want {
		t.Errorf("Got: %d. Want: %d", got, want)
	}
}

func TestAnswertPartTwo(t *testing.T) {
	want := 241861950
	got, _ := day1.AnswerPartTwo(&testdata)

	if got != want {
		t.Errorf("Got: %d. Want: %d", got, want)
	}
}
