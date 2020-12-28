package day11

const (
	seatEmpty    = 'L'
	seatOccupied = '#'
	seatFloor    = '.'
)

var (
	directions = [][]int{
		{-1, 0},
		{-1, 1},
		{-1, -1},
		{1, 0},
		{1, 1},
		{1, -1},
		{0, 1},
		{0, -1},
	}
)

// Solution implements the day 11 solution
type Solution struct{}

// PartOne answers part 1 of the day 11 puzzle
func (s *Solution) PartOne(input *[]string) (int, error) {
	var rule seatingRule = adjacentSeatingRule{}
	return simulateSeating(input, &rule), nil
}

func simulateSeating(input *[]string, r *seatingRule) int {
	seats := seatLayout{input}
	occupied := seats.countOccupancy()
	for {
		seats.applyRule(r)
		occupiedPrev := occupied
		occupied = seats.countOccupancy()
		if occupied == occupiedPrev {
			break
		}
	}
	return occupied
}

// PartTwo answers part 2 of the day 11 puzzle
func (s *Solution) PartTwo(input *[]string) (int, error) {
	var rule seatingRule = allDirectionsSeatingRule{}
	return simulateSeating(input, &rule), nil
}

type seatingRule interface {
	apply(s *seatLayout, x, y int) rune
}

type adjacentSeatingRule struct{}
type allDirectionsSeatingRule struct{}

func (r adjacentSeatingRule) apply(s *seatLayout, x, y int) rune {
	seatState := s.getAtPosition(x, y)
	occupied := s.countAdjacentOccupancy(x, y)
	switch {
	case seatState == seatEmpty && occupied == 0:
		return seatOccupied
	case seatState == seatOccupied && occupied >= 4:
		return seatEmpty
	default:
		return seatState
	}
}

func (r allDirectionsSeatingRule) apply(s *seatLayout, x, y int) rune {
	seatState := s.getAtPosition(x, y)
	occupied := s.countVisibleOccupancy(x, y)
	switch {
	case seatState == seatEmpty && occupied == 0:
		return seatOccupied
	case seatState == seatOccupied && occupied >= 5:
		return seatEmpty
	default:
		return seatState
	}
}

type seatLayout struct {
	seats *[]string
}

func (s *seatLayout) getAtPosition(x, y int) rune {
	if x < 0 || x >= s.getWidth() {
		return rune(0)
	}
	if y < 0 || y >= s.getHeight() {
		return rune(0)
	}
	return rune((*s.seats)[y][x])
}

func (s *seatLayout) getWidth() int {
	row := (*s.seats)[0]
	return len(row)
}

func (s *seatLayout) getHeight() int {
	return len(*s.seats)
}

func (s *seatLayout) applyRule(r *seatingRule) {
	seatsNew := make([]string, s.getHeight())
	for y := 0; y < s.getHeight(); y++ {
		rowNew := []rune((*s.seats)[y])
		for x := 0; x < s.getWidth(); x++ {
			rowNew[x] = (*r).apply(s, x, y)
		}
		seatsNew[y] = string(rowNew)
	}
	s.seats = &seatsNew
}

func (s *seatLayout) countAdjacentOccupancy(x, y int) int {
	count := 0
	for _, dir := range directions {
		dx, dy := dir[0], dir[1]
		if s.getAtPosition(x+dx, y+dy) == seatOccupied {
			count++
		}
	}
	return count
}

func (s *seatLayout) countVisibleOccupancy(x, y int) int {
	count := 0
	for _, dir := range directions {
		dx, dy := dir[0], dir[1]
		xTemp, yTemp := x, y
		for {
			xTemp, yTemp = xTemp+dx, yTemp+dy
			seatState := s.getAtPosition(xTemp, yTemp)
			if seatState == seatOccupied {
				count++
			}
			if seatState == rune(0) || seatState == seatOccupied || seatState == seatEmpty {
				break
			}
		}
	}
	return count
}

func (s *seatLayout) countOccupancy() int {
	count := 0
	for _, row := range *s.seats {
		for _, str := range row {
			if str == seatOccupied {
				count++
			}
		}
	}
	return count
}
