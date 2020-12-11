package main

import (
	"fmt"
	"github.com/bowmanrogere/advent-of-code-2020/internal"
	"log"
	"os"
	"sort"
)

func main() {
	lines, err := internal.ReadFile(fmt.Sprintf("%s/Development/advent-of-code-2020/cmd/day10/adapters.txt", os.Getenv("HOME")))
	if err != nil {
		log.Fatal(err)
	}

	adapters := internal.ConvertToInts(lines)
	sort.Ints(adapters)
	// add built in adapter
	adapters = append(adapters, adapters[len(adapters)-1] + 3)
	puzzle1(adapters)
}

func puzzle1(adapters []int) {
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
	println(fmt.Sprintf("Puzzle #1: %d", len(differences[1]) * len(differences[3])))
}


