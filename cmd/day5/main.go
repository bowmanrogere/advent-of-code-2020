package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/bowmanrogere/advent-of-code-2020/internal"
)

func main() {
	lines, err := internal.ReadFile("boarding-passes.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle1(lines)
	puzzle2(lines)
}

func puzzle1(lines []string) {
	seatIDs := seatIds(lines)

	maxSeatID := 0

	for _, id := range seatIDs {
		if id > maxSeatID {
			maxSeatID = id
		}
	}

	println(fmt.Sprintf("Max Seat ID: %d", maxSeatID))
}

func puzzle2(lines []string) {
	seatIds := seatIds(lines)

	mySeat := 0

	sort.Ints(seatIds)

	for index, id := range seatIds {
		if index == 0 || index >= len(seatIds) {
			continue
		}

		seatBefore := seatIds[index-1]

		if id-1 != seatBefore {
			mySeat = id - 1
			break
		}
	}

	println(fmt.Sprintf("My Seat ID: %d", mySeat))
}

func seatIds(lines []string) []int {
	seatIds := make([]int, 0)

	for _, line := range lines {
		bp := internal.NewBoardingPass(line)
		seatIds = append(seatIds, bp.SeatID())
	}

	return seatIds
}
