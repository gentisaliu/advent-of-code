package puzzle

import "strings"

// Day6 implements the day 6 puzzle
type Day6 struct{}

// AnswerPartOne answers part 1 of the day 6 puzzle
func (d *Day6) AnswerPartOne(input *[]string) (int, error) {
	yesAnswers := parseYesAnswers(input)
	return countYesAnswers(&yesAnswers), nil
}

func parseYesAnswers(input *[]string) []yesAnswersInGroup {
	var yesAnswersAllGroups []yesAnswersInGroup
	yesAnswersCurrGroup := newYesAnswersInGroup()
	for _, line := range *input {
		if strings.TrimSpace(line) == "" {
			yesAnswersAllGroups = append(yesAnswersAllGroups, yesAnswersCurrGroup)
			yesAnswersCurrGroup = newYesAnswersInGroup()
			continue
		}
		yesAnswersCurrGroup.peopleCount++
		for _, questionID := range line {
			yesAnswersCurrGroup.countByQuestion[questionID]++
		}
	}
	return yesAnswersAllGroups
}

func newYesAnswersInGroup() yesAnswersInGroup {
	return yesAnswersInGroup{
		countByQuestion: make(map[rune]int),
	}
}

func countYesAnswers(yesAnswersAllGroups *[]yesAnswersInGroup) int {
	var yesCount int
	for _, groupAnswers := range *yesAnswersAllGroups {
		yesCount += len(groupAnswers.countByQuestion)
	}
	return yesCount
}

// AnswerPartTwo answers part 2 of the day 6 puzzle
func (d *Day6) AnswerPartTwo(input *[]string) (int, error) {
	yesAnswers := parseYesAnswers(input)
	return countQuestionsAnsweredYesByAllInGroups(&yesAnswers), nil
}

func countQuestionsAnsweredYesByAllInGroups(allGroupsAnswers *[]yesAnswersInGroup) int {
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

type yesAnswersInGroup struct {
	countByQuestion map[rune]int
	peopleCount     int
}
