package day5

import (
	"errors"
	"strings"

	"github.com/gentisaliu/advent-of-code/2020/internal/util"
)

// Solution implements the day 5 puzzle
type Solution struct{}

// PartOne answers part 1 of the day 5 puzzle
func (d *Solution) PartOne(input *[]string) (int, error) {
	seats := decodeBoardingPasses(input)
	return findMaxSeatID(&seats), nil
}

func findMaxSeatID(seats *[]seat) int {
	var maxSeatID int
	for _, seat := range *seats {
		if seat.getID() > maxSeatID {
			maxSeatID = seat.getID()
		}
	}
	return maxSeatID
}

func decodeBoardingPasses(input *[]string) []seat {
	var seats []seat
	for _, boardingPass := range *input {
		seat := decodeBoardingPass(boardingPass)
		seats = append(seats, seat)
	}
	return seats
}

func decodeBoardingPass(boardingPass string) seat {
	rowPartition := boardingPass[0:7]
	row := decodePartition(rowPartition, "F", "B")
	colPartition := boardingPass[7:]
	col := decodePartition(colPartition, "L", "R")

	return seat{row: row, column: col}
}

func decodePartition(partition string, lowerHalfLetter string, upperHalfLetter string) int {
	binary := convertLettersToBinary(partition, lowerHalfLetter, upperHalfLetter)
	return int(util.ConvertBinaryToInt(binary))
}

func convertLettersToBinary(partition string, lowerHalfLetter string, upperHalfLetter string) string {
	binary := strings.ReplaceAll(strings.ToUpper(partition), string(lowerHalfLetter), "0")
	return strings.ReplaceAll(binary, string(upperHalfLetter), "1")
}

// PartTwo answers part 2 of the day 5 puzzle
func (d *Solution) PartTwo(input *[]string) (int, error) {
	seats := decodeBoardingPasses(input)
	seatsIDs := getSeatIDs(&seats)
	missingSeatID, err := findMissingSeatBetweenSeatsInList(&seatsIDs)
	return missingSeatID, err
}

func getSeatIDs(seats *[]seat) map[int]bool {
	seatIDs := make(map[int]bool)
	for _, seat := range *seats {
		seatIDs[seat.getID()] = true
	}
	return seatIDs
}

func findMissingSeatBetweenSeatsInList(seatIDs *map[int]bool) (int, error) {
	// Do not look for seats at the very front (row 0) and back (row 127)
	minSeatID, maxSeatID := calculateSeatID(1, 0), calculateSeatID(126, 8)
	var seatID = minSeatID + 1
	for seatID < maxSeatID-1 {
		if !(*seatIDs)[seatID] && (*seatIDs)[seatID-1] && (*seatIDs)[seatID+1] {
			return seatID, nil
		}
		seatID++
	}
	return -1, errors.New("could not find missing seat between two seats in list")
}

func calculateSeatID(row, column int) int {
	return row*8 + column
}

type seat struct {
	row    int
	column int
}

func (s *seat) getID() int {
	return calculateSeatID(s.row, s.column)
}
