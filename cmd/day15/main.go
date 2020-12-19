package main

import "fmt"

func main() {
	p1 := memoryGame(2020, 14, 8, 16, 0, 1, 17)
	println(fmt.Sprintf("Puzzle #1: %d", p1))

	p2 := memoryGame(30000000, 14, 8, 16, 0, 1, 17)
	println(fmt.Sprintf("Puzzle #2: %d", p2))
}

func memoryGame(numGuesses int, initial ...int) int {
	spoken := make(map[int][]int)

	for idx, i := range initial {
		spoken[i] = []int{idx}
	}

	lastSpoken := initial[len(initial)-1]

	for i := len(initial); i < numGuesses; i++ {
		// never been spoken before
		if idxs := spoken[lastSpoken]; len(idxs) == 1 {
			lastSpoken = 0
		} else {
			lastSpoken = idxs[len(idxs)-1] - idxs[len(idxs)-2]
		}

		if _, ok := spoken[lastSpoken]; !ok {
			spoken[lastSpoken] = []int{i}
		} else {
			spoken[lastSpoken] = append(spoken[lastSpoken], i)
		}
	}

	return lastSpoken
}
