package main

import (
	"fmt"
	"github.com/bowmanrogere/advent-of-code-2020/internal"
	"log"
)

type Slope struct {
	x int
	y int
}

func main() {
	lines, err := internal.ReadFile("map.txt")
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	puzzle1(lines)
	puzzle2(lines)
}

func puzzle1(lines []string) {
	slope := &Slope{
		x: 3,
		y: 1,
	}
	treesEncountered := findTreesForSlope(lines, slope)
	println(fmt.Sprintf("# of trees encountered: %d", treesEncountered))
}

func puzzle2(lines []string) {
	slopes := []*Slope{
		{x: 1, y: 1},
		{x: 3, y: 1},
		{x: 5, y: 1},
		{x: 7, y: 1},
		{x: 1, y: 2},
	}

	answer := 1
	for _, slope := range slopes {
		answer = answer * findTreesForSlope(lines, slope)
	}

	println(fmt.Sprintf("Number of trees multiplied together: %d", answer))
}

func findTreesForSlope(lines []string, slope *Slope) int {
	// this will be what we use to determine when to repeat the line.
	length := len(lines[0])

	idx := 0
	treesEncountered := 0
	for lineIndex, line := range lines {
		if lineIndex%slope.y == 0 {
			if lineIndex > 0 {
				if string(line[idx%length]) == "#" {
					treesEncountered++
				}
			}

			idx += slope.x
		}
	}

	return treesEncountered
}
