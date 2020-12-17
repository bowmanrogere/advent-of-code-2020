package waypoint

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

type Waypoint struct {
	X int
	Y int
}

type Ferry struct {
	directionRegex   *regexp.Regexp
	UnitsEast        int
	UnitsNorth       int
	currentDirection int
	directions       []string
	waypoint         *Waypoint
}

func NewWaypointFerry(directions []string) *Ferry {
	return &Ferry{
		UnitsEast:        0,
		UnitsNorth:       0,
		currentDirection: north,
		directions:       directions,
		waypoint: &Waypoint{
			X: 10,
			Y: 1,
		},
	}
}

func (w *Ferry) Sail() {
	w.directionRegex = regexp.MustCompile(`^([NSEWRLF])([0-9]+)$`)

	processNEWS := func(d, u int) {
		switch d {
		case north:
			w.waypoint.Y += u
		case east:
			w.waypoint.X += u
		case west:
			w.waypoint.X -= u
		case south:
			w.waypoint.Y -= u
		}
	}

	processLR := func(d, u int) {
		if d == left {
			if u == 90 {
				u = 270
			} else if u == 270 {
				u = 90
			}
		}

		switch u {
		case 90:
			tmpX := w.waypoint.X
			tmpY := w.waypoint.Y
			w.waypoint.X = tmpY
			w.waypoint.Y = tmpX * -1
		case 180:
			w.waypoint.X *= -1
			w.waypoint.Y *= -1
		case 270:
			tmpX := w.waypoint.X
			tmpY := w.waypoint.Y
			w.waypoint.X = tmpY * -1
			w.waypoint.Y = tmpX
		}
	}

	processForward := func(d, u int) {
		w.UnitsNorth += w.waypoint.Y * u
		w.UnitsEast += w.waypoint.X * u
	}

	for _, direction := range w.directions {
		d, u := w.parseDirection(direction)

		switch d {
		case north, east, west, south:
			processNEWS(d, u)
		case left, right:
			processLR(d, u)
		case forward:
			processForward(d, u)
		}
	}
}

func (w *Ferry) parseDirection(direction string) (int, int) {
	parts := w.directionRegex.FindStringSubmatch(direction)
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
