package puzzle

// Puzzle defines the contract of a puzzle
type Puzzle interface {
	AnswerPartOne(input *[]string) (int, error)
	AnswerPartTwo(input *[]string) (int, error)
}
