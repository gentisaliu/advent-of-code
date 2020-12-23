package puzzle

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	errInputIncorrect = `input '%v' has incorrect format. Valid format: '[number]-[number] [letter]: [password]', 
	where [minRepeat], [maxNumber] are numbers, [letter] is a single character, [password] is a single word containing 
	latin characters, digits or underscores`
	errInputMissingPassword = `input %v is missing the password to apply the policy to`
)

// Day2 implements the day 1 puzzle
type Day2 struct{}

// AnswerPartOne answers part 1 of the day 2 puzzle
func (d *Day2) AnswerPartOne(input *[]string) (int, error) {
	return countValidPasswords(input, func(parsed inputParsed) passwordPolicy {
		return &sledRentalPasswordPolicy{
			letter:   parsed.letter,
			minOccur: parsed.number1,
			maxOccur: parsed.number2,
		}
	})
}

func countValidPasswords(input *[]string, createPolicy func(parsed inputParsed) passwordPolicy) (int, error) {
	validPasswords := 0
	for _, v := range *input {
		inputParsed, err := parseInput(v)
		if err != nil {
			return validPasswords, err
		}
		policy := createPolicy(inputParsed)
		if policy.isPasswordValid(inputParsed.password) {
			validPasswords++
		}
	}
	return validPasswords, nil
}

// AnswerPartTwo answers part 2 of the day 2 puzzle
func (d *Day2) AnswerPartTwo(input *[]string) (int, error) {
	return countValidPasswords(input, func(parsed inputParsed) passwordPolicy {
		return &tobogganCorporatePasswordPolicy{
			letter:    parsed.letter,
			position1: parsed.number1 - 1,
			position2: parsed.number2 - 1,
		}
	})
}

func parseInput(inputTxt string) (inputParsed, error) {
	var input inputParsed
	var password string
	var err error

	re := regexp.MustCompile(`(\d+)-(\d+)\s+(\w):\s+(\w+)`)
	submatch := re.FindStringSubmatch(inputTxt)
	err = validateInput(inputTxt, submatch)
	if err != nil {
		return input, err
	}

	// parse input
	letter := rune(submatch[3][0])
	number1, _ := strconv.Atoi(submatch[1])
	number2, _ := strconv.Atoi(submatch[2])
	password = submatch[4]
	input = inputParsed{
		letter:   letter,
		number1:  number1,
		number2:  number2,
		password: password,
	}
	return input, err
}

func validateInput(input string, submatch []string) error {
	var err error
	switch {
	case len(submatch) != 5:
		err = fmt.Errorf(errInputIncorrect, input)
	case len(submatch[1]) == 0:
		err = fmt.Errorf(errInputIncorrect, input)
	case len(submatch[2]) == 0:
		err = fmt.Errorf(errInputIncorrect, input)
	case len(submatch[3]) != 1:
		err = fmt.Errorf(errInputIncorrect, input)
	case len(submatch[4]) == 0:
		err = fmt.Errorf(errInputIncorrect, input)
	}
	return err
}

type inputParsed struct {
	number1  int
	number2  int
	letter   rune
	password string
}

type passwordPolicy interface {
	isPasswordValid(password string) bool
}

type sledRentalPasswordPolicy struct {
	letter   rune
	minOccur int
	maxOccur int
}

type tobogganCorporatePasswordPolicy struct {
	letter    rune
	position1 int
	position2 int
}

func (p *sledRentalPasswordPolicy) isPasswordValid(password string) bool {
	occurrences := strings.Count(password, string(p.letter))
	return p.minOccur <= occurrences && occurrences <= p.maxOccur
}

func (p *tobogganCorporatePasswordPolicy) isPasswordValid(password string) bool {
	letterPos1 := password[p.position1]
	letterPos2 := password[p.position2]

	if letterPos1 == byte(p.letter) && letterPos2 != byte(p.letter) {
		return true
	} else if letterPos1 != byte(p.letter) && letterPos2 == byte(p.letter) {
		return true
	}
	return false
}
