package day13

import (
	"strconv"
	"strings"
)

// Solution implements the day 13 solution
type Solution struct{}

// PartOne answers part 1 of the day 13 puzzle
func (s *Solution) PartOne(input *[]string) (int, error) {
	earliestDepart, busIDs := parsePartOneInput(input)
	busID, waitTime := findEarliestBus(earliestDepart, busIDs)
	return busID * waitTime, nil
}

func parsePartOneInput(input *[]string) (int, []int) {
	earliestDepart, _ := strconv.Atoi((*input)[0])
	busIDs := parsePartOneBusIDs((*input)[1])
	return earliestDepart, busIDs
}

func parsePartOneBusIDs(busIDsRow string) []int {
	busIDs := make([]int, 0)
	busIDsTokens := strings.Split(busIDsRow, ",")
	for _, busIDRaw := range busIDsTokens {
		if busIDRaw == "x" {
			continue
		}
		busID, _ := strconv.Atoi(busIDRaw)
		busIDs = append(busIDs, busID)
	}
	return busIDs
}

func findEarliestBus(earliestDepart int, busIDs []int) (int, int) {
	minBusID := busIDs[0]
	minWaitTime := calcWaitTime(earliestDepart, minBusID)
	for _, busID := range busIDs {
		waitTime := calcWaitTime(earliestDepart, busID)
		if waitTime < minWaitTime {
			minWaitTime = waitTime
			minBusID = busID
		}
	}
	return minBusID, minWaitTime
}

func calcWaitTime(earliestDepart, busID int) int {
	return busID - (earliestDepart % busID)
}

// PartTwo answers part 2 of the day 13 puzzle
func (s *Solution) PartTwo(input *[]string) (int, error) {
	busIDs := parsePartTwoInput(input)
	timestamp := findConsecutiveBusDeparts(busIDs)
	return timestamp, nil
}

func parsePartTwoInput(input *[]string) []int {
	busIDsRow := (*input)[0]
	if len(*input) == 2 {
		busIDsRow = (*input)[1]
	}
	return parsePartTwoBusIDs(busIDsRow)
}

func parsePartTwoBusIDs(busIDsRow string) []int {
	busIDs := make([]int, 0)
	busIDsTokens := strings.Split(busIDsRow, ",")
	for _, busIDToken := range busIDsTokens {
		busID := 1
		if busIDToken != "x" {
			busID, _ = strconv.Atoi(busIDToken)
		}
		busIDs = append(busIDs, busID)
	}
	return busIDs
}

func findConsecutiveBusDeparts(busIDs []int) int {
	timestamp, step := 0, 1
	for i, busID := range busIDs {
		if busID == 1 {
			continue
		}
		for (timestamp+i)%busID > 0 {
			timestamp += step
		}
		step *= busID
	}
	return timestamp
}
