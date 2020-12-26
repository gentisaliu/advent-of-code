package day8

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	opJmp                = 1
	opAcc                = 2
	opNop                = 3
	exitCodeOk           = 0
	exitCodeError        = 1
	exitCodeInfiniteLoop = 2
)

// Solution implements the day 8 solution
type Solution struct{}

// PartOne answers part 1 of the day 8 puzzle
func (d *Solution) PartOne(input *[]string) (int, error) {
	cpu := processor{}
	cpu.enableInfiniteLoopDetection()
	cpu.loadProgram(input)
	_, err := cpu.executeProgram()
	return cpu.accumulator, err
}

// PartTwo answers part 2 of the day 8 puzzle
func (d *Solution) PartTwo(input *[]string) (int, error) {
	cpu := processor{}
	cpu.enableInfiniteLoopDetection()
	cpu.loadProgram(input)
	exitCode, err := cpu.executeProgram()
	if exitCode == exitCodeInfiniteLoop && err == nil {
		for lineNr := range *input {
			program, programChanged := changeJmpNopInstructionInProgramLine(input, lineNr)
			if !programChanged {
				continue
			}
			cpu.loadProgram(program)
			exitCode, err = cpu.executeProgram()
			if exitCode == exitCodeOk {
				break
			}
		}
	}
	return cpu.accumulator, err
}

func changeJmpNopInstructionInProgramLine(program *[]string, lineNr int) (*[]string, bool) {
	instruction := (*program)[lineNr]
	instructionChanged := instruction

	if strings.HasPrefix(instruction, "jmp") {
		instructionChanged = strings.ReplaceAll(instruction, "jmp", "nop")
	} else if strings.HasPrefix(instruction, "nop") {
		instructionChanged = strings.ReplaceAll(instruction, "nop", "jmp")
	}

	if instructionChanged != instruction {
		newProgram := replaceSliceElement(program, lineNr, instructionChanged)
		return newProgram, true
	}
	return program, false
}

func replaceSliceElement(slice *[]string, index int, newElem string) *[]string {
	newSlice := make([]string, len(*slice))
	copy(newSlice, *slice)
	newSlice[index] = newElem
	return &newSlice
}

type opCode int

type instruction struct {
	instruction string
	opCode      opCode
	argument    int
	ordinal     int
}

type processor struct {
	instructionSequence  *[]string
	programCounter       int
	accumulator          int
	detectInfiniteLoops  bool
	executedInstructions map[int]bool
}

func (c *processor) loadProgram(instructionSequence *[]string) {
	c.instructionSequence = instructionSequence
	c.programCounter = 0
	c.accumulator = 0
	c.executedInstructions = make(map[int]bool)
}

func (c *processor) fetchInstruction() (string, int) {
	instruction := (*c.instructionSequence)[c.programCounter]
	ordinal := c.programCounter
	c.programCounter++
	return instruction, ordinal
}

func (c *processor) decodeInstruction(instructionTxt string, ordinal int) (instruction, error) {
	var err error
	decodedInstruction := instruction{}
	decodedInstruction.opCode, err = c.decodeOpCode(instructionTxt)
	decodedInstruction.argument, err = c.decodeArgument(instructionTxt)
	decodedInstruction.ordinal = ordinal
	if err != nil {
		err = fmt.Errorf("error decoding instruction '%v': %v", instructionTxt, err)
	}
	return decodedInstruction, err
}

func (c *processor) decodeOpCode(instructionTxt string) (opCode, error) {
	opcode := instructionTxt[0:3]
	switch opcode {
	case "acc":
		return opAcc, nil
	case "jmp":
		return opJmp, nil
	case "nop":
		return opNop, nil
	default:
		return 0, fmt.Errorf("operation %v is not supported (supported: acc, jmp, nop)", opcode)
	}
}

func (c *processor) decodeArgument(instructionTxt string) (int, error) {
	argument, _ := strconv.Atoi(instructionTxt[5:])
	sign := instructionTxt[4:5]
	switch sign {
	case "-":
		return -argument, nil
	case "+":
		return argument, nil
	default:
		return 0, fmt.Errorf("%v is not a valid sign character (valid: +, -)", sign)
	}
}

func (c *processor) executeInstruction(i *instruction) {
	switch i.opCode {
	case opAcc:
		c.accumulator += i.argument
	case opJmp:
		c.programCounter = i.ordinal + i.argument
	}
}

func (c *processor) performInstructionCycle() error {
	instructionTxt, ordinal := c.fetchInstruction()
	instruction, err := c.decodeInstruction(instructionTxt, ordinal)
	c.executeInstruction(&instruction)
	c.executedInstructions[ordinal] = true
	return err
}

func (c *processor) enableInfiniteLoopDetection() {
	c.detectInfiniteLoops = true
}

func (c *processor) executeProgram() (int, error) {
	for !c.endOfInstructionSequenceReached() {
		if c.detectInfiniteLoops && c.instructionPreviouslyExecuted() {
			return exitCodeInfiniteLoop, nil
		}
		err := c.performInstructionCycle()
		if err != nil {
			return exitCodeError, err
		}
	}
	return exitCodeOk, nil
}

func (c *processor) instructionPreviouslyExecuted() bool {
	_, found := c.executedInstructions[c.programCounter]
	return found
}

func (c *processor) endOfInstructionSequenceReached() bool {
	return len(*c.instructionSequence) <= c.programCounter
}
