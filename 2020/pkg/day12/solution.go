package day12

import (
	"math"
	"strconv"
)

const (
	north = 0
	east  = 1
	south = 2
	west  = 3
)

// Solution implements the day 12 solution
type Solution struct{}

// PartOne answers part 1 of the day 12 puzzle
func (s *Solution) PartOne(input *[]string) (int, error) {
	ship := getShip()
	for _, instruction := range *input {
		interpretPartOneInstructions(&ship, instruction)
	}
	return ship.manhattanDistance(), nil
}

func interpretPartOneInstructions(s *ship, instruction string) {
	direction, unit := parseInstruction(instruction)
	switch direction {
	case "N":
		s.moveNorth(unit)
	case "S":
		s.moveSouth(unit)
	case "E":
		s.moveEast(unit)
	case "W":
		s.moveWest(unit)
	case "L":
		s.turnLeft(unit)
	case "R":
		s.turnRight(unit)
	case "F":
		s.moveForward(unit)
	}
}

func parseInstruction(instruction string) (string, int) {
	unit, _ := strconv.Atoi(instruction[1:])
	direction := string(instruction[0])
	return direction, unit
}

// PartTwo answers part 2 of the day 12 puzzle
func (s *Solution) PartTwo(input *[]string) (int, error) {
	waypoint := getWaypoint()
	ship := getShip()
	for _, instruction := range *input {
		interpretPartTwoInstructions(&waypoint, &ship, instruction)
	}
	return ship.manhattanDistance(), nil
}

func interpretPartTwoInstructions(w *waypoint, s *ship, instruction string) {
	direction, unit := parseInstruction(instruction)
	switch direction {
	case "N":
		w.moveNorth(unit)
	case "S":
		w.moveSouth(unit)
	case "E":
		w.moveEast(unit)
	case "W":
		w.moveWest(unit)
	case "L":
		w.rotateLeft(unit)
	case "R":
		w.rotateRight(unit)
	case "F":
		s.moveBy(unit*w.posRelativeToShip.east, unit*w.posRelativeToShip.north)
	}
}

type waypoint struct {
	posRelativeToShip position
}

func getWaypoint() waypoint {
	return waypoint{position{10, 1}}
}

func (w *waypoint) moveNorth(units int) {
	w.posRelativeToShip.north += units
}

func (w *waypoint) moveSouth(units int) {
	w.posRelativeToShip.north -= units
}

func (w *waypoint) moveEast(units int) {
	w.posRelativeToShip.east += units
}

func (w *waypoint) moveWest(units int) {
	w.posRelativeToShip.east -= units
}

func (w *waypoint) rotateRight(deg int) {
	turns := deg / 90
	for i := 0; i < turns; i++ {
		oldEast := w.posRelativeToShip.east
		w.posRelativeToShip.east = w.posRelativeToShip.north
		w.posRelativeToShip.north = -oldEast
	}
}

func (w *waypoint) rotateLeft(deg int) {
	w.rotateRight(360 - deg)
}

type ship struct {
	pos       position
	direction int
}

func getShip() ship {
	return ship{position{0, 0}, east}
}

func (s *ship) moveNorth(units int) {
	s.pos.north += units
}

func (s *ship) moveSouth(units int) {
	s.pos.north -= units
}

func (s *ship) moveEast(units int) {
	s.pos.east += units
}

func (s *ship) moveWest(units int) {
	s.pos.east -= units
}

func (s *ship) moveBy(east, north int) {
	s.pos.east += east
	s.pos.north += north
}

func (s *ship) turnRight(deg int) {
	turns := deg / 90
	for i := 0; i < turns; i++ {
		switch s.direction {
		case west:
			s.direction = north
		case north:
			s.direction = east
		case east:
			s.direction = south
		case south:
			s.direction = west
		}
	}
}

func (s *ship) turnLeft(deg int) {
	s.turnRight(360 - deg)
}

func (s *ship) moveForward(units int) {
	switch s.direction {
	case west:
		s.moveWest(units)
	case east:
		s.moveEast(units)
	case south:
		s.moveSouth(units)
	case north:
		s.moveNorth(units)
	}
}

func (s *ship) manhattanDistance() int {
	return int(math.Abs(float64(s.pos.east)) + math.Abs(float64(s.pos.north)))
}

type position struct {
	east  int
	north int
}
