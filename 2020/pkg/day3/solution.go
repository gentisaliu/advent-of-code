package day3

// Solution implements the day 3 puzzle
type Solution struct{}

// PartOne answers part 1 of the day 3 puzzle
func (d *Solution) PartOne(input *[]string) (int, error) {
	slope := slope{right: 3, down: 1}
	world := world{*input}
	return countTreesToWorldBottom(&world, slope), nil
}

func countTreesToWorldBottom(w *world, s slope) int {
	x, y := 0, 0
	worldH := w.getHeight()
	treeCount := 0
	for y < worldH {
		if w.isTreeAtPosition(x, y) {
			treeCount++
		}
		x += s.right
		y += s.down
	}
	return treeCount
}

// PartTwo answers part 2 of the day 3 puzzle
func (d *Solution) PartTwo(input *[]string) (int, error) {
	slopes := []slope{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}
	world := world{*input}
	treeCountProd := 1
	for _, slope := range slopes {
		treeCountProd *= countTreesToWorldBottom(&world, slope)
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

func (w *world) getAtPosition(x int, y int) rune {
	line := w.features[y]
	xProjected := x % w.getWidth()
	return rune(line[xProjected])
}

func (w *world) isTreeAtPosition(x int, y int) bool {
	const tree = '#'
	return w.getAtPosition(x, y) == tree
}
