package puzzle

// Puzzle defines the contract of a puzzle
type Puzzle interface {
	AnswerPartOne(input *[]string) int
	AnswerPartTwo(input *[]string) int
}
