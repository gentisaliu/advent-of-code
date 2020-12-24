package puzzle

import (
	"regexp"
	"strconv"
)

const myBag = "shiny gold"

var ruleRegex = regexp.MustCompile(`([0-9] |^)[a-z]+ [a-z]+`)

// Day7 implements the day 7 puzzle
type Day7 struct{}

// AnswerPartOne answers part 1 of the day 7 puzzle
func (d *Day7) AnswerPartOne(input *[]string) (int, error) {
	bagsMatrix := createParentChildBagMatrix(input)
	return bagsMatrix.countParentsWithChildBag(myBag), nil
}

func createParentChildBagMatrix(input *[]string) *parentChildBagMatrix {
	bags := &parentChildBagMatrix{
		matrix: make(map[string]map[string]int),
	}
	for _, line := range *input {
		matches := ruleRegex.FindAllString(line, -1)
		parentBag := matches[0]
		for _, childBagMatch := range matches[1:] {
			capacity, _ := strconv.Atoi(childBagMatch[0:1])
			childBag := childBagMatch[2:]
			bags.addEdge(parentBag, childBag, capacity)
		}
	}
	return bags
}

// AnswerPartTwo answers part 2 of the day 7 puzzle
func (d *Day7) AnswerPartTwo(input *[]string) (int, error) {
	bagsMatrix := createParentChildBagMatrix(input)
	return bagsMatrix.countIndividualChildren(myBag), nil
}

type parentChildBagMatrix struct {
	matrix map[string]map[string]int
}

func (p *parentChildBagMatrix) addEdge(parent string, child string, capacity int) {
	if _, parentFound := p.matrix[parent]; !parentFound {
		p.matrix[parent] = make(map[string]int)
	}
	p.matrix[parent][child] = capacity
}

func (p *parentChildBagMatrix) countParentsWithChildBag(childBag string) int {
	var count int
	for parentBag := range p.matrix {
		if _, childFound := p.matrix[parentBag][childBag]; childFound {
			count++
			continue
		}
		if p.isIndirectChildOfParent(parentBag, childBag) {
			count++
		}
	}
	return count
}

func (p *parentChildBagMatrix) isIndirectChildOfParent(parentBag, childBag string) bool {
	// BFS
	mapVisited := map[string]bool{parentBag: true}
	queue := []string{parentBag}
	for len(queue) > 0 {
		b := queue[0]
		queue = queue[1:] // dequeue
		for currChild := range p.matrix[b] {
			if _, visited := mapVisited[currChild]; !visited {
				mapVisited[currChild] = true
				if currChild == childBag {
					return true
				}
				queue = append(queue, currChild) // enqueue
			}
		}
	}
	return false
}

func (p *parentChildBagMatrix) countIndividualChildren(parentBag string) int {
	var count int
	for childBag, capacity := range p.matrix[parentBag] {
		count += capacity * (1 + p.countIndividualChildren(childBag))
	}
	return count
}
