package day10

import (
	"sort"

	"github.com/gentisaliu/advent-of-code/2020/internal/util"
)

// Solution implements the day 10 solution
type Solution struct{}

// PartOne answers part 1 of the day 10 puzzle
func (d *Solution) PartOne(input *[]string) (int, error) {
	jolts := util.ConvertStringsToNumbers(input)
	chainedAdapters := chainAdapters(*jolts)
	diff1, diff3 := countJoltDifferences(chainedAdapters)
	return diff1 * diff3, nil
}

func chainAdapters(jolts []int) []int {
	chargeOutletJolt := []int{0}
	joltsSorted := append(chargeOutletJolt, jolts...)
	sort.Ints(joltsSorted)

	maxJolt := joltsSorted[len(joltsSorted)-1]
	deviceAdapterJolt := maxJolt + 3
	return append(joltsSorted, deviceAdapterJolt)
}

func countJoltDifferences(jolts []int) (int, int) {
	joltsCount := len(jolts)
	differences := map[int]int{1: 0, 2: 0, 3: 0}
	for i := 1; i < joltsCount; i++ {
		differences[jolts[i]-jolts[i-1]]++
	}
	return differences[1], differences[3]
}

// PartTwo answers part 2 of the day 10 puzzle
func (d *Solution) PartTwo(input *[]string) (int, error) {
	jolts := util.ConvertStringsToNumbers(input)
	chainedAdapters := chainAdapters(*jolts)
	return countAdapterArrangements(&chainedAdapters), nil
}

func countAdapterArrangements(adapters *[]int) int {
	// dynamic programming approach with O(n^2) time complexity
	// tried iterative depth-first-search approach first,
	// which provided correct answers, but would not terminate
	// for the given puzzle input
	deviceAdapterJolt := (*adapters)[len(*adapters)-1]
	arrangementsByJolt := make(map[int]int)
	arrangementsByJolt[deviceAdapterJolt] = 1
	for i := len(*adapters) - 2; i >= 0; i-- {
		jolt1 := (*adapters)[i]
		for _, jolt2 := range *adapters {
			diff := jolt2 - jolt1
			if diff > 0 && diff <= 3 {
				arrangementsByJolt[jolt1] += arrangementsByJolt[jolt2]
			}
		}
	}
	chargingOutletJolt := (*adapters)[0]
	return arrangementsByJolt[chargingOutletJolt]
}
