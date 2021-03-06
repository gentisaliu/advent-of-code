package day4

import (
	"regexp"
	"strings"
)

var (
	pairDelimiter     = " "
	keyValueDelimiter = ":"
	byrRegex          = regexp.MustCompile(`^(19[2-9][0-9]|200[1-2])$`)
	iyrRegex          = regexp.MustCompile(`^(201[0-9]|2020)$`)
	eyrRegex          = regexp.MustCompile(`^(202[0-9]|2030)$`)
	hgtRegex          = regexp.MustCompile(`^(1[5-8][0-9]cm|19[0-3]cm|59in|6[0-9]in|7[0-6]in)$`)
	hclRegex          = regexp.MustCompile(`^#[0-9A-Fa-f]{6}$`)
	eclRegex          = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	pidRegex          = regexp.MustCompile(`^[0-9]{9}$`)
)

// Solution implements the day 4 puzzle
type Solution struct{}

// PartOne answers part 1 of the day 4 puzzle
func (d *Solution) PartOne(input *[]string) (int, error) {
	passports := parsePassports(input)
	passportCount := countPassportsWithNoMissingValues(&passports)
	return passportCount, nil
}

func parsePassports(input *[]string) []passport {
	var passports []passport
	var currentPassport passport
	for _, line := range *input {
		if strings.TrimSpace(line) == "" {
			passports = append(passports, currentPassport)
			currentPassport = passport{}
			continue
		}
		entries := strings.Split(line, pairDelimiter)
		for _, kvp := range entries {
			keyValue := strings.Split(kvp, keyValueDelimiter)
			key := keyValue[0]
			value := keyValue[1]
			switch key {
			case "byr":
				currentPassport.byr = value
			case "iyr":
				currentPassport.iyr = value
			case "eyr":
				currentPassport.eyr = value
			case "hgt":
				currentPassport.hgt = value
			case "hcl":
				currentPassport.hcl = value
			case "ecl":
				currentPassport.ecl = value
			case "pid":
				currentPassport.pid = value
			case "cid":
				currentPassport.cid = value
			}
		}
	}
	return passports
}

func countPassportsWithNoMissingValues(passports *[]passport) int {
	count := 0
	for _, passport := range *passports {
		if passport.requiredFieldsProvided() {
			count++
		}
	}
	return count
}

// PartTwo answers part 2 of the day 4 puzzle
func (d *Solution) PartTwo(input *[]string) (int, error) {
	passports := parsePassports(input)
	passportCount := countPassportsWithNoMissingAndValidValues(&passports)
	return passportCount, nil
}

func countPassportsWithNoMissingAndValidValues(passports *[]passport) int {
	count := 0
	for _, passport := range *passports {
		if passport.requiredFieldsProvided() && passport.fieldsHaveValidValues() {
			count++
		}
	}
	return count
}

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p *passport) requiredFieldsProvided() bool {
	return p.byr != "" && p.iyr != "" && p.eyr != "" &&
		p.hgt != "" && p.hcl != "" && p.ecl != "" && p.pid != ""
}

func (p *passport) fieldsHaveValidValues() bool {
	return byrRegex.MatchString(p.byr) &&
		iyrRegex.MatchString(p.iyr) &&
		eyrRegex.MatchString(p.eyr) &&
		hgtRegex.MatchString(p.hgt) &&
		hclRegex.MatchString(p.hcl) &&
		eclRegex.MatchString(p.ecl) &&
		pidRegex.MatchString(p.pid)
}
