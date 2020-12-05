package main

import (
	"fmt"
	"github.com/bowmanrogere/advent-of-code-2020/internal"
	"log"
)

func main() {
	//boardingPass, err := internal.NewBoardingPass("FBFBBFFRLR")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//println(fmt.Sprintf("Boarding Pass: %+v", *boardingPass))
	//println(fmt.Sprintf("Seat ID: %d", boardingPass.SeatID()))

	lines, err := internal.ReadFile("boarding-passes.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle1(lines)
}

func puzzle1(lines []string) {
	seatIds := make([]int, 0)
	for _, line := range lines {
		bp := internal.NewBoardingPass(line)
		seatIds = append(seatIds, bp.SeatID())
	}

	maxSeatId := 0
	for _, id := range seatIds {
		if id > maxSeatId {
			maxSeatId = id
		}
	}

	println(fmt.Sprintf("Max Seat ID: %d", maxSeatId))
}


