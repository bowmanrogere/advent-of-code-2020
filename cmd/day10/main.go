package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/bowmanrogere/advent-of-code-2020/internal"
	"github.com/bowmanrogere/advent-of-code-2020/internal/graph"
)

func main() {
	lines, err := internal.ReadFile(fmt.Sprintf("%s/Development/advent-of-code-2020/cmd/day10/adapters.txt", os.Getenv("HOME")))
	if err != nil {
		log.Fatal(err)
	}

	adapters := internal.ConvertToInts(lines)
	sort.Ints(adapters)
	puzzle1(adapters)
	puzzle2(adapters)
}

func puzzle1(adapters []int) {
	// add built in adapter
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	differences := make(map[int][]int)

	previous := 0

	for _, adapter := range adapters {
		diff := adapter - previous
		if _, ok := differences[diff]; !ok {
			differences[diff] = make([]int, 0)
		}

		differences[diff] = append(differences[diff], adapter)
		previous = adapter
	}

	println(fmt.Sprintf("Puzzle #1: %d", len(differences[1])*len(differences[3])))
}

func puzzle2(adapters []int) {
	adapters = append(adapters, 0)

	sort.Ints(adapters)

	graphs := make([]*graph.Graph, 0)
	continuous := []int{0}
	previous := 0

	for i := 1; i < len(adapters); i++ {
		if adapters[i]-previous != 3 {
			continuous = append(continuous, adapters[i])
		} else {
			if len(continuous) > 1 {
				graphs = append(graphs, createGraph(continuous))
			}
			// start continuous over
			continuous = []int{adapters[i]}
		}

		previous = adapters[i]
	}

	// if we have continous left over, create a graph for it
	if len(continuous) > 1 {
		graphs = append(graphs, createGraph(continuous))
	}

	numPaths := 1

	for _, g := range graphs {
		if len(g.Paths) > 0 {
			numPaths *= len(g.Paths)
		}
	}

	println(fmt.Sprintf("Puzzle Two: %d", numPaths))
}

func createGraph(continuous []int) *graph.Graph {
	g := graph.NewGraph(len(continuous))

	for j := 0; j < len(continuous); j++ {
		for k := j + 1; k <= j+3 && k < len(continuous); k++ {
			g.AddEdge(continuous[j], continuous[k])
		}
	}

	g.CreatePaths(continuous[0], continuous[len(continuous)-1])

	return g
}
