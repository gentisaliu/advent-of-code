package day9

import (
	"errors"
	"fmt"

	"github.com/gentisaliu/advent-of-code/2020/internal/list"
)

var (
	preambleLength = 25
)

// Solution implements the day 9 solution
type Solution struct{}

// PartOne answers part 1 of the day 9 puzzle
func (d *Solution) PartOne(input *[]string) (int, error) {
	numbers := list.ConvertStringSliceToNumberSlice(input)
	encodedData := &xmasEncodedData{numbers}
	return encodedData.findFirstCorruptedNumber()
}

// PartTwo answers part 2 of the day 9 puzzle
func (d *Solution) PartTwo(input *[]string) (int, error) {
	numbers := list.ConvertStringSliceToNumberSlice(input)
	encodedData := &xmasEncodedData{numbers}
	return encodedData.findEncryptionWeakness()
}

func getMinMax(numbers *[]int) (int, int) {
	min, max := (*numbers)[0], (*numbers)[0]
	for i := range *numbers {
		number := (*numbers)[i]
		if number > max {
			max = number
		} else if number < min {
			min = number
		}
	}
	return min, max
}

type xmasEncodedData struct {
	numbers *[]int
}

func (d *xmasEncodedData) findFirstCorruptedNumber() (int, error) {
	// initialize preamble
	preamble := preamble{}
	preamble.initialize(d.numbers)
	numberCount := len(*d.numbers)
	for i := preamble.endIndex; i < numberCount; i++ {
		numberToCheck := (*d.numbers)[i]
		numberCheckOk := false
		for preambleNumber := range preamble.numbersMap {
			if preamble.contains(numberToCheck - preambleNumber) {
				numberCheckOk = true
				break
			}
		}
		if numberCheckOk {
			preamble.update(d.numbers)
		} else {
			return numberToCheck, nil
		}
	}
	return 0, errors.New("no currupted entry found")
}

func (d *xmasEncodedData) findEncryptionWeakness() (int, error) {
	corruptedNumber, err := d.findFirstCorruptedNumber()
	if err != nil {
		return 0, err
	}
	corruptedNumberIndex := d.getElementIndex(corruptedNumber)
	sliceSet, err := d.findContiguousSetWithSum(corruptedNumber, 0, corruptedNumberIndex)
	if err != nil {
		return 0, err
	}
	min, max := getMinMax(&sliceSet)
	return min + max, nil
}

func (d *xmasEncodedData) getElementIndex(el int) int {
	for i := range *d.numbers {
		if (*d.numbers)[i] == el {
			return i
		}
	}
	return -1
}

func (d *xmasEncodedData) findContiguousSetWithSum(targetSum, startIndex, endIndex int) ([]int, error) {
	for i := range (*d.numbers)[startIndex:endIndex] {
		sum := 0
		for j, number := range (*d.numbers)[i:] {
			sum += number
			if sum == targetSum {
				return (*d.numbers)[i : i+j+1], nil
			}
		}
	}
	return nil, fmt.Errorf(
		"no contiguous set with sum %d found between index %d and %d", targetSum, startIndex, endIndex)
}

type preamble struct {
	// keys are the numbers, values are their number of occurrence
	numbersMap map[int]int
	startIndex int
	endIndex   int
}

func (p *preamble) initialize(numbers *[]int) {
	p.numbersMap = make(map[int]int)
	p.startIndex = 0
	p.endIndex = preambleLength

	for _, number := range (*numbers)[p.startIndex:p.endIndex] {
		p.add(number)
	}
}

func (p *preamble) update(numbers *[]int) {
	p.remove((*numbers)[p.startIndex])
	p.add((*numbers)[p.endIndex])
	p.startIndex++
	p.endIndex++
}

func (p *preamble) add(number int) {
	if !p.contains(number) {
		p.numbersMap[number] = 0
	}
	p.numbersMap[number]++
}

func (p *preamble) contains(number int) bool {
	_, found := p.numbersMap[number]
	return found
}

func (p *preamble) remove(number int) {
	if count, found := p.numbersMap[number]; found {
		if count > 1 {
			p.numbersMap[number]--
		} else {
			delete(p.numbersMap, number)
		}
	}
}
