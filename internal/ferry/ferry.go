package ferry

import (
	"regexp"
	"strconv"
)

const (
	north   = 0
	east    = 90
	south   = 180
	west    = 270
	forward = 1
	left    = 2
	right   = 3
)

type Ferry struct {
	directionRegex   *regexp.Regexp
	UnitsEast        int
	UnitsNorth       int
	currentDirection int
	directions       []string
}

func NewFerry(directions []string) *Ferry {
	return &Ferry{
		currentDirection: east,
		directions:       directions,
		UnitsEast:        0,
		UnitsNorth:       0,
	}
}

func (f *Ferry) Sail() {
	f.directionRegex = regexp.MustCompile(`^([NSEWRLF])([0-9]+)$`)
	processNEWS := func(d, u int) {
		switch d {
		case north:
			f.UnitsNorth += u
		case south:
			f.UnitsNorth -= u
		case east:
			f.UnitsEast += u
		case west:
			f.UnitsEast -= u
		}
	}

	processLR := func(d, u int) {
		if d == left {
			u *= -1
		}

		direction := f.currentDirection

		direction = (direction + u) % 360

		if direction < 0 {
			direction = direction + 360
		}

		f.currentDirection = direction
	}

	for _, direction := range f.directions {
		d, u := f.parseDirection(direction)

		switch d {
		case north, east, west, south:
			processNEWS(d, u)
		case forward:
			processNEWS(f.currentDirection, u)
		case left, right:
			processLR(d, u)
		}
	}
}

func (f *Ferry) parseDirection(direction string) (int, int) {
	parts := f.directionRegex.FindStringSubmatch(direction)
	d := 0
	switch parts[1] {
	case "N":
		d = north
	case "E":
		d = east
	case "S":
		d = south
	case "W":
		d = west
	case "L":
		d = left
	case "R":
		d = right
	case "F":
		d = forward
	}
	u, _ := strconv.Atoi(parts[2])
	return d, u
}
