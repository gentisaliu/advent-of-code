package puzzle

// Day3 implements the day 3 puzzle
type Day3 struct{}

// AnswerPartOne answers part 1 of the day 3 puzzle
func (d *Day3) AnswerPartOne(input *[]string) (int, error) {
	s := slope{right: 3, down: 1}
	w := world{*input}
	return countTreesToWorldBottom(&w, s), nil
}

func countTreesToWorldBottom(w *world, s slope) int {
	x, y := 0, 0
	worldH := w.getHeight()
	treeCount := 0
	for y < worldH {
		if w.isTreeLocatedAtPosition(x, y) {
			treeCount++
		}
		x += s.right
		y += s.down
	}
	return treeCount
}

// AnswerPartTwo answers part 2 of the day 3 puzzle
func (d *Day3) AnswerPartTwo(input *[]string) (int, error) {
	slopes := []slope{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}
	w := world{*input}
	treeCountProd := 1
	for _, s := range slopes {
		treeCountProd *= countTreesToWorldBottom(&w, s)
	}
	return treeCountProd, nil
}

type slope struct {
	right int
	down  int
}

type world struct {
	features []string
}

func (w *world) getWidth() int {
	line := w.features[0]
	return len(line)
}

func (w *world) getHeight() int {
	return len(w.features)
}

func (w *world) getFeatureAtPosition(x int, y int) rune {
	line := w.features[y]
	xProjected := x % w.getWidth()
	return rune(line[xProjected])
}

func (w *world) isTreeLocatedAtPosition(x int, y int) bool {
	const tree = '#'
	return w.getFeatureAtPosition(x, y) == tree
}
