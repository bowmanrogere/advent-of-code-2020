package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/bowmanrogere/advent-of-code-2020/internal"
	"github.com/bowmanrogere/advent-of-code-2020/internal/ferry"
	"github.com/bowmanrogere/advent-of-code-2020/internal/waypoint"
)

func main() {
	directions, err := internal.ReadFile(fmt.Sprintf("%s/Development/advent-of-code-2020/cmd/day12/directions.txt", os.Getenv("HOME")))
	if err != nil {
		log.Fatal(err)
	}

	puzzle1(directions)
	puzzle2(directions)
}

func puzzle1(directions []string) {
	f := ferry.NewFerry(directions)
	f.Sail()
	println(fmt.Sprintf("Puzzle #1: %d", manhattanDistance(f.UnitsNorth, f.UnitsEast)))
}

func puzzle2(directions []string) {
	w := waypoint.NewWaypointFerry(directions)
	w.Sail()
	println(fmt.Sprintf("Puzzle #2: %d", manhattanDistance(w.UnitsNorth, w.UnitsEast)))
}

func manhattanDistance(unitsNorth, unitsEast int) int {
	return int(math.Abs(float64(unitsNorth)) + math.Abs(float64(unitsEast)))
}
