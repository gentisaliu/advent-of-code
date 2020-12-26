package day

// Solution defines the contract for the solution of a given day
type Solution interface {
	PartOne(input *[]string) (int, error)
	PartTwo(input *[]string) (int, error)
}
