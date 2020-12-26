package day7

import (
	"regexp"
	"strconv"
)

const myBag = "shiny gold"

var ruleRegex = regexp.MustCompile(`([0-9] |^)[a-z]+ [a-z]+`)

// Solution implements the day 7 solution
type Solution struct{}

// PartOne answers part 1 of the day 7 puzzle
func (d *Solution) PartOne(input *[]string) (int, error) {
	bagsGraph := createParentChildBagMatrix(input)
	return bagsGraph.countParentsWithChildBag(myBag), nil
}

func createParentChildBagMatrix(input *[]string) *weightedDirectedGraph {
	bags := &weightedDirectedGraph{
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

// PartTwo answers part 2 of the day 7 puzzle
func (d *Solution) PartTwo(input *[]string) (int, error) {
	bagsGraph := createParentChildBagMatrix(input)
	return bagsGraph.countIndividualChildren(myBag), nil
}

type weightedDirectedGraph struct {
	matrix map[string]map[string]int
}

func (w *weightedDirectedGraph) addEdge(parent string, child string, capacity int) {
	if _, parentFound := w.matrix[parent]; !parentFound {
		w.matrix[parent] = make(map[string]int)
	}
	w.matrix[parent][child] = capacity
}

func (w *weightedDirectedGraph) countParentsWithChildBag(childBag string) int {
	var count int
	for parentBag := range w.matrix {
		switch {
		case w.isDirectChild(parentBag, childBag):
			count++
		case w.isIndirectChild(parentBag, childBag):
			count++
		}
	}
	return count
}

func (w *weightedDirectedGraph) isDirectChild(parentBag, childBag string) bool {
	_, childFound := w.matrix[parentBag][childBag]
	return childFound
}

func (w *weightedDirectedGraph) isIndirectChild(parentBag, childBag string) bool {
	// BFS
	mapVisited := map[string]bool{parentBag: true}
	queue := []string{parentBag}
	for len(queue) > 0 {
		b := queue[0]
		queue = queue[1:] // dequeue
		for currChild := range w.matrix[b] {
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

func (w *weightedDirectedGraph) countIndividualChildren(parentBag string) int {
	var count int
	for childBag, capacity := range w.matrix[parentBag] {
		count += capacity * (1 + w.countIndividualChildren(childBag))
	}
	return count
}
