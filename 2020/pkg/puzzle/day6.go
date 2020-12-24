package puzzle

import "strings"

// Day6 implements the day 6 puzzle
type Day6 struct{}

// AnswerPartOne answers part 1 of the day 6 puzzle
func (d *Day6) AnswerPartOne(input *[]string) (int, error) {
	groupsAnswers := parseQuestionsAnsweredYesByGroups(input)
	return countQuestionsAnsweredYes(&groupsAnswers), nil
}

func parseQuestionsAnsweredYesByGroups(input *[]string) []groupYesAnswers {
	var yesAnswersGroups []groupYesAnswers
	yesAnswersCurrGroup := newGroupYesAnswers()
	for _, line := range *input {
		if strings.TrimSpace(line) == "" {
			yesAnswersGroups = append(yesAnswersGroups, yesAnswersCurrGroup)
			yesAnswersCurrGroup = newGroupYesAnswers()
			continue
		}
		yesAnswersCurrGroup.peopleCount++
		for _, questionID := range line {
			yesAnswersCurrGroup.countByQuestion[questionID]++
		}
	}
	return yesAnswersGroups
}

func newGroupYesAnswers() groupYesAnswers {
	return groupYesAnswers{
		countByQuestion: make(map[rune]int),
	}
}

func countQuestionsAnsweredYes(groupsAnswers *[]groupYesAnswers) int {
	var yesCount int
	for _, groupAnswers := range *groupsAnswers {
		yesCount += len(groupAnswers.countByQuestion)
	}
	return yesCount
}

// AnswerPartTwo answers part 2 of the day 6 puzzle
func (d *Day6) AnswerPartTwo(input *[]string) (int, error) {
	groupsAnswers := parseQuestionsAnsweredYesByGroups(input)
	return countQuestionsAnsweredYesByAllInGroups(&groupsAnswers), nil
}

func countQuestionsAnsweredYesByAllInGroups(allGroupsAnswers *[]groupYesAnswers) int {
	var answeredYesByAllCount int
	for _, groupAnswers := range *allGroupsAnswers {
		for _, a := range groupAnswers.countByQuestion {
			if a == groupAnswers.peopleCount {
				answeredYesByAllCount++
			}
		}
	}
	return answeredYesByAllCount
}

type groupYesAnswers struct {
	countByQuestion map[rune]int
	peopleCount     int
}
