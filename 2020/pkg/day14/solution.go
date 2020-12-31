package day14

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gentisaliu/advent-of-code/2020/internal/util"
)

var (
	maskRegex   = regexp.MustCompile(`^mask = ([0|1|X]{36})$`)
	memoryRegex = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
)

const (
	bitmaskLength = 36
)

// Solution implements day 14 solution
type Solution struct{}

// PartOne answers part 1 of the day 14 puzzle
func (s *Solution) PartOne(input *[]string) (int, error) {
	memory := programMemory{}
	program := ferryDockingProgram{input, &memory}
	program.executeV1()
	return int(memory.sumMemoryValues()), nil
}

// PartTwo answers part 2 of the day 14 puzzle
func (s *Solution) PartTwo(input *[]string) (int, error) {
	memory := programMemory{}
	program := ferryDockingProgram{input, &memory}
	program.executeV2()
	return int(memory.sumMemoryValues()), nil
}

type programMemory struct {
	bitmask string
	content map[int]int
}

func (m *programMemory) sumMemoryValues() int {
	sum := 0
	for _, memoryValue := range m.content {
		sum += memoryValue
	}
	return sum
}

type ferryDockingProgram struct {
	instructions *[]string
	memory       *programMemory
}

func (p *ferryDockingProgram) executeV1() {
	p.execute(func(memAddr, memValue int) {
		p.memory.content[memAddr] = p.maskMemoryValue(memValue)
	})
}

func (p *ferryDockingProgram) execute(memoryInstructionFunc func(int, int)) {
	p.memory.content = make(map[int]int)

	for _, instruction := range *p.instructions {
		if maskRegex.MatchString(instruction) {
			submatch := maskRegex.FindStringSubmatch(instruction)
			p.memory.bitmask = submatch[1]
		} else if memoryRegex.MatchString(instruction) {
			memAddr, memValue := parseMemInstruction(instruction)
			memoryInstructionFunc(memAddr, memValue)
		}
	}
}

func parseMemInstruction(instruction string) (int, int) {
	submatch := memoryRegex.FindStringSubmatch(instruction)
	memAddrStr, valueStr := submatch[1], submatch[2]
	value, _ := strconv.Atoi(valueStr)
	memAddr, _ := strconv.Atoi(memAddrStr)
	return memAddr, value
}

func (p *ferryDockingProgram) executeV2() {
	p.execute(func(memAddr, memValue int) {
		memAddrBin := util.ConvertIntToBinary(memAddr, bitmaskLength)
		memAddrMasked := p.maskMemoryAddress(memAddrBin)
		memAddresses := p.replaceFloatingBits(memAddrMasked)
		for _, addr := range memAddresses {
			addrVal := util.ConvertBinaryToInt(addr)
			p.memory.content[int(addrVal)] = memValue
		}
	})
}

func (p *ferryDockingProgram) maskMemoryValue(memValue int) int {
	memValueBin := util.ConvertIntToBinary(memValue, bitmaskLength)
	memValueBinRune := []rune(memValueBin)
	for i, bitChar := range p.memory.bitmask {
		if bitChar != '0' && bitChar != '1' {
			continue
		}
		memValueBinRune[i] = bitChar
	}
	memBinStr := string(memValueBinRune)
	return int(util.ConvertBinaryToInt(memBinStr))
}

func (p *ferryDockingProgram) maskMemoryAddress(memAddrBin string) string {
	memAddrBinRune := []rune(memAddrBin)
	for i, bitChar := range p.memory.bitmask {
		if bitChar == 'X' || bitChar == '1' {
			memAddrBinRune[i] = bitChar
		}
	}
	return string(memAddrBinRune)
}

func (p *ferryDockingProgram) replaceFloatingBits(memAddrBin string) (addresses []string) {
	stack := []string{memAddrBin}
	for len(stack) > 0 {
		memAddrBinCurrent := stack[0]
		stack = stack[1:]
		floatBitIndex := strings.Index(memAddrBinCurrent, "X")
		if floatBitIndex == -1 {
			addresses = append(addresses, memAddrBinCurrent)
			continue
		}
		memAddrBinCurrent0 := util.SetStringCharAtIndex(memAddrBinCurrent, floatBitIndex, '0')
		memAddrBinCurrent1 := util.SetStringCharAtIndex(memAddrBinCurrent, floatBitIndex, '1')
		stack = append(stack, memAddrBinCurrent0)
		stack = append(stack, memAddrBinCurrent1)
	}
	return addresses
}
