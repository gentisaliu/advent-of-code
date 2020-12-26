package day2

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	entitiesRegEx = regexp.MustCompile(`^(\d+)-(\d+)\s+(\w):\s+(\w+)$`)
)

// Solution solves the day 2 puzzle
type Solution struct{}

// PartOne answers part 1 of the day 2 puzzle
func (d *Solution) PartOne(input *[]string) (int, error) {
	entities, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	passwords := parsePasswords(&entities, createSledRentalPasswordPolicy)
	return countValidPasswords(&passwords), nil
}

func parseInput(input *[]string) ([]inputEntities, error) {
	var entities []inputEntities
	for _, v := range *input {
		inputLineEntities := extractEntitiesFromLineOfInput(v)
		entities = append(entities, inputLineEntities)
	}
	return entities, nil
}

func extractEntitiesFromLineOfInput(inputLine string) inputEntities {
	var entities inputEntities
	var password string

	// extract entities
	submatch := entitiesRegEx.FindStringSubmatch(inputLine)
	letter := rune(submatch[3][0])
	number1, _ := strconv.Atoi(submatch[1])
	number2, _ := strconv.Atoi(submatch[2])
	password = submatch[4]
	entities = inputEntities{
		letter:   letter,
		number1:  number1,
		number2:  number2,
		password: password,
	}
	return entities
}

func parsePasswords(entities *[]inputEntities, policyFunc func(entities inputEntities) passwordPolicy) []password {
	var passwords []password
	for _, item := range *entities {
		passwords = append(passwords, password{
			passwordValue: item.password,
			policy:        policyFunc(item),
		})
	}
	return passwords
}

func createSledRentalPasswordPolicy(entities inputEntities) passwordPolicy {
	return &sledRentalPasswordPolicy{
		letter:   entities.letter,
		minOccur: entities.number1,
		maxOccur: entities.number2,
	}
}

func countValidPasswords(passwords *[]password) int {
	validPasswords := 0
	for _, password := range *passwords {
		if password.isValid() {
			validPasswords++
		}
	}
	return validPasswords
}

// PartTwo answers part 2 of the day 2 puzzle
func (d *Solution) PartTwo(input *[]string) (int, error) {
	entities, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	passwords := parsePasswords(&entities, createTobogganCorporatePasswordPolicy)
	return countValidPasswords(&passwords), nil
}

func createTobogganCorporatePasswordPolicy(entities inputEntities) passwordPolicy {
	return &tobogganCorporatePasswordPolicy{
		letter:        entities.letter,
		positionLeft:  entities.number1 - 1,
		positionRight: entities.number2 - 1,
	}
}

type inputEntities struct {
	number1  int
	number2  int
	letter   rune
	password string
}

type passwordPolicy interface {
	isValidPassword(password string) bool
}

type sledRentalPasswordPolicy struct {
	letter   rune
	minOccur int
	maxOccur int
}

type tobogganCorporatePasswordPolicy struct {
	letter        rune
	positionLeft  int
	positionRight int
}

type password struct {
	passwordValue string
	policy        passwordPolicy
}

func (p *password) isValid() bool {
	return p.policy.isValidPassword(p.passwordValue)
}

func (p *sledRentalPasswordPolicy) isValidPassword(password string) bool {
	occurrences := strings.Count(password, string(p.letter))
	return p.minOccur <= occurrences && occurrences <= p.maxOccur
}

func (p *tobogganCorporatePasswordPolicy) isValidPassword(password string) bool {
	letterLeft := password[p.positionLeft]
	letterRight := password[p.positionRight]

	occursLeft := letterLeft == byte(p.letter)
	occursRight := letterRight == byte(p.letter)

	return occursLeft != occursRight
}
