package puzzle

import "strings"

// Day6 implements the day 5 puzzle
type Day6 struct{}

// AnswerPartOne answers part 1 of the day 6 puzzle
func (d *Day6) AnswerPartOne(input *[]string) (int, error) {
	groupsAnswers := parseYesAnswersOfGroups(input)
	return countQuestionsAnsweredYes(&groupsAnswers), nil
}

func parseYesAnswersOfGroups(input *[]string) []groupAnswers {
	var answersGroups []groupAnswers
	answersCurrGroup := newGroupAnswers()
	for _, line := range *input {
		if strings.TrimSpace(line) == "" {
			answersGroups = append(answersGroups, answersCurrGroup)
			answersCurrGroup = newGroupAnswers()
			continue
		}
		answersCurrGroup.peopleCount++
		for _, questionID := range line {
			answersCurrGroup.yesAnswerCountByQuestion[questionID]++
		}
	}
	return answersGroups
}

func newGroupAnswers() groupAnswers {
	return groupAnswers{
		yesAnswerCountByQuestion: make(map[rune]int),
	}
}

func countQuestionsAnsweredYes(groupsAnswers *[]groupAnswers) int {
	var yesCount int
	for _, groupAnswers := range *groupsAnswers {
		yesCount += len(groupAnswers.yesAnswerCountByQuestion)
	}
	return yesCount
}

// AnswerPartTwo answers part 2 of the day 6 puzzle
func (d *Day6) AnswerPartTwo(input *[]string) (int, error) {
	groupsAnswers := parseYesAnswersOfGroups(input)
	return countQuestionsAnsweredYesByAllInGroups(&groupsAnswers), nil
}

func countQuestionsAnsweredYesByAllInGroups(allGroupsAnswers *[]groupAnswers) int {
	var answeredYesByAllCount int
	for _, groupAnswers := range *allGroupsAnswers {
		for _, a := range groupAnswers.yesAnswerCountByQuestion {
			if a == groupAnswers.peopleCount {
				answeredYesByAllCount++
			}
		}
	}
	return answeredYesByAllCount
}

type groupAnswers struct {
	yesAnswerCountByQuestion map[rune]int
	peopleCount              int
}
