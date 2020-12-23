package puzzle

import (
	"errors"
	"strconv"
	"strings"
)

// Day5 implements the day 5 puzzle
type Day5 struct{}

// AnswerPartOne answers part 1 of the day 5 puzzle
func (d *Day5) AnswerPartOne(input *[]string) (int, error) {
	seats := decodeBoardingPasses(input)
	return getHighestSeatID(&seats), nil
}

func getHighestSeatID(seats *[]seat) int {
	var highestSeatID int
	for _, seat := range *seats {
		if seat.getID() > highestSeatID {
			highestSeatID = seat.getID()
		}
	}
	return highestSeatID
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
	binary := convertPartitionToBinary(partition, lowerHalfLetter, upperHalfLetter)
	return convertBinaryToInt(binary)
}

func convertPartitionToBinary(partition string, lowerHalfLetter string, upperHalfLetter string) string {
	binary := strings.ReplaceAll(strings.ToUpper(partition), string(lowerHalfLetter), "0")
	return strings.ReplaceAll(binary, string(upperHalfLetter), "1")
}

func convertBinaryToInt(binaryString string) int {
	n, _ := strconv.ParseInt(binaryString, 2, 8)
	return int(n)
}

// AnswerPartTwo answers part 2 of the day 5 puzzle
func (d *Day5) AnswerPartTwo(input *[]string) (int, error) {
	occupiedSeats := decodeBoardingPasses(input)
	occupiedSeatsIDs := getSeatIDs(&occupiedSeats)
	mySeatID, err := findEmptySeatSurroundedByOccupiedSeats(&occupiedSeatsIDs)
	return mySeatID, err
}

func getSeatIDs(seats *[]seat) map[int]bool {
	seatIDs := make(map[int]bool)
	for _, seat := range *seats {
		seatIDs[seat.getID()] = true
	}
	return seatIDs
}

func findEmptySeatSurroundedByOccupiedSeats(seatIDs *map[int]bool) (int, error) {
	// Do not look for seats at the very front (row 0) and back (row 127) rows
	minSeatID, maxSeatID := calculateSeatID(1, 0), calculateSeatID(126, 8)
	var seatID = minSeatID + 1
	for seatID < maxSeatID-1 {
		if !(*seatIDs)[seatID] && (*seatIDs)[seatID-1] && (*seatIDs)[seatID+1] {
			return seatID, nil
		}
		seatID++
	}
	return -1, errors.New("could not find empty seat surrounded by two occupied seats")
}

type seat struct {
	row    int
	column int
}

func (s *seat) getID() int {
	return calculateSeatID(s.row, s.column)
}

func calculateSeatID(row, column int) int {
	return row*8 + column
}
