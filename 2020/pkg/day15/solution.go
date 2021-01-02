package day15

import (
	"strconv"
	"strings"
)

// Solution implements day 15 solution
type Solution struct{}

// PartOne answers part 1 of the day 15 puzzle
func (s *Solution) PartOne(input *[]string) (int, error) {
	numbers := parseInput(input)
	return playMemoryGame(&numbers, 2020), nil
}

func parseInput(input *[]string) (numbers []int) {
	numberTokens := strings.Split((*input)[0], ",")
	for _, numberToken := range numberTokens {
		number, _ := strconv.Atoi(numberToken)
		numbers = append(numbers, number)
	}
	return numbers
}

func playMemoryGame(startingNumbers *[]int, turns int) (result int) {
	m := newMemory()
	for i, number := range *startingNumbers {
		m.addNumber(number, i+1)
	}
	for turn := len(*startingNumbers) + 1; turn <= turns; turn++ {
		m.addNumber(m.age(), turn)
		result = m.lastNumber
	}
	return result
}

type memory struct {
	lastNumber    int
	numbersToTurn map[int][]int
}

func newMemory() memory {
	m := memory{}
	m.numbersToTurn = make(map[int][]int)
	return m
}

func (m *memory) age() int {
	if turns, found := m.numbersToTurn[m.lastNumber]; !found {
		return 0
	} else if turns[0] == -1 || turns[1] == -1 {
		return 0
	} else {
		return turns[1] - turns[0]
	}
}

func (m *memory) addNumber(number, turn int) {
	_, found := m.numbersToTurn[number]
	if !found {
		m.numbersToTurn[number] = make([]int, 2)
		m.numbersToTurn[number][0] = -1
		m.numbersToTurn[number][1] = -1
	}
	if m.numbersToTurn[number][0] == -1 {
		m.numbersToTurn[number][0] = turn
	} else if m.numbersToTurn[number][1] == -1 {
		m.numbersToTurn[number][1] = turn
	} else {
		m.numbersToTurn[number][0] = m.numbersToTurn[number][1]
		m.numbersToTurn[number][1] = turn
	}
	m.lastNumber = number
}

// PartTwo answers part 2 of the day 15 puzzle
func (s *Solution) PartTwo(input *[]string) (int, error) {
	numbers := parseInput(input)
	return playMemoryGame(&numbers, 30000000), nil
}
