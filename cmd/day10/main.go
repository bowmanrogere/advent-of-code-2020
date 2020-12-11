package main

import (
	"fmt"
	"github.com/bowmanrogere/advent-of-code-2020/internal"
	"log"
	"sort"
)

func main() {
	lines, err := internal.ReadFile("/Users/t24974a/Development/advent-of-code-2020/cmd/day10/adapters.txt")
	if err != nil {
		log.Fatal(err)
	}

	adapters := internal.ConvertToInts(lines)
	puzzle1(adapters)
}

func puzzle1(adapters []int) {
	sort.Ints(adapters)
	differences := make(map[int][]int)

	previous := 0
	for idx, adapter := range adapters {
		diff := adapter - previous
		if _, ok := differences[diff]; !ok {
			differences[diff] = make([]int, 0)
		}
		differences[diff] = append(differences[diff], adapter)

		// built-in adapter
		if idx == len(adapters) - 1 {
			differences[3] = append(differences[3], adapter + 3)
		}

		previous = adapter
	}

	println(fmt.Sprintf("1 jolt differences: %d", len(differences[1])))
	println(fmt.Sprintf("3 jolt differences: %d", len(differences[3])))
	println(fmt.Sprintf("Puzzle #1: %d", len(differences[1]) * len(differences[3])))
}


