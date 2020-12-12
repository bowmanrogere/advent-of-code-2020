package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bowmanrogere/advent-of-code-2020/internal"
)

func main() {
	lines, err := internal.ReadFile(fmt.Sprintf("%s/Development/advent-of-code-2020/cmd/day11/seat-map.txt", os.Getenv("HOME")))
	if err != nil {
		log.Fatal(err)
	}

	puzzle1Seats := make([]string, len(lines))
	puzzle2Seats := make([]string, len(lines))
	copy(puzzle1Seats, lines)
	copy(puzzle2Seats, lines)
	puzzle1(puzzle1Seats)
	puzzle2(puzzle2Seats)
}

func puzzle1(seats []string) {
	// run the puzzle until no changes
	changed := true
	for changed {
		occupiedSeats := make([]string, 0)
		for row := 0; row < len(seats); row++ {
			rowVal := seats[row]
			newRowVal := ""
			// check each position
			for col := 0; col < len(rowVal); col++ {
				if string(rowVal[col]) == "." {
					newRowVal += "."
					continue
				}

				occupiedAdjacent := adjacentSeats(row, col, seats)
				if string(rowVal[col]) == "L" && occupiedAdjacent == 0 {
					newRowVal += "#"
				} else if string(rowVal[col]) == "#" && occupiedAdjacent >= 4 {
					newRowVal += "L"
				} else {
					newRowVal += string(rowVal[col])
				}
			}
			occupiedSeats = append(occupiedSeats, newRowVal)
		}
		changed = hasChanged(seats, occupiedSeats)
		copy(seats, occupiedSeats)
	}

	// count number of occupied seats
	occupiedSeats := 0
	for _, row := range seats {
		for _, seat := range row {
			if string(seat) == "#" {
				occupiedSeats++
			}
		}
	}

	println(fmt.Sprintf("Puzzle #1: %d", occupiedSeats))
}

func adjacentSeats(row, col int, seats []string) int {
	above := 0
	current := 0
	below := 0
	if row > 0 {
		above = checkSeats(col, seats[row - 1], true)
	}
	current = checkSeats(col, seats[row], false)
	if row < len(seats) - 1 {
		below = checkSeats(col, seats[row + 1], true)
	}
	return above + current + below
}

func checkSeats(index int, seats string, checkIndex bool) int {
	left := 0
	center := 0
	right := 0
	if index > 0 {
		if seatOccupied(index - 1, seats) {
			left++
		}
	}
	if checkIndex {
		if seatOccupied(index, seats) {
			center++
		}
	}
	if index < len(seats) - 1 {
		if seatOccupied(index + 1, seats) {
			right++
		}
	}
	return left + center + right
}

func hasChanged(original, new []string) bool {
	if len(original) != len(new) {
		return true
	}

	for i := 0; i < len(original); i++ {
		if original[i] != new[i] {
			return true
		}
	}
	return false
}

func puzzle2(seats []string) {
	// run the puzzle until no changes
	changed := true
	for changed {
		occupiedSeats := make([]string, 0)
		for row := 0; row < len(seats); row++ {
			rowVal := seats[row]
			newRowVal := ""
			// check each position
			for col := 0; col < len(rowVal); col++ {
				if string(rowVal[col]) == "." {
					newRowVal += "."
					continue
				}

				occupiedAdjacent := firstAdjacentSeats(row, col, seats)
				if string(rowVal[col]) == "L" && occupiedAdjacent == 0 {
					newRowVal += "#"
				} else if string(rowVal[col]) == "#" && occupiedAdjacent >= 5 {
					newRowVal += "L"
				} else {
					newRowVal += string(rowVal[col])
				}
			}
			occupiedSeats = append(occupiedSeats, newRowVal)
		}
		changed = hasChanged(seats, occupiedSeats)
		copy(seats, occupiedSeats)
	}

	// count number of occupied seats
	occupiedSeats := 0
	for _, row := range seats {
		for _, seat := range row {
			if string(seat) == "#" {
				occupiedSeats++
			}
		}
	}

	println(fmt.Sprintf("Puzzle #2: %d", occupiedSeats))
}

func firstAdjacentSeats(row, col int, seats []string) int {
	aboveLeft := 0
	above := 0
	aboveRight := 0
	left := 0
	right := 0
	belowLeft := 0
	below := 0
	belowRight := 0

	if row > 0 {
		if col > 0 {
			if checkAboveLeft(row-1, col-1, seats) {
				aboveLeft++
			}
		}
		if checkAbove(row - 1, col, seats) {
			above++
		}
		if col < len(seats[row]) - 1 {
			if checkAboveRight(row - 1, col + 1, seats) {
				aboveRight++
			}
		}
	}
	if col > 0 {
		if checkLeft(col - 1, seats[row]) {
			left++
		}
	}
	if col < len(seats[row]) - 1 {
		if checkRight(col + 1, seats[row]) {
			right++
		}
	}
	if row < len(seats) - 1 {
		if col > 0 {
			if checkBelowLeft(row + 1, col - 1, seats) {
				belowLeft++
			}
		}
		if checkBelow(row + 1, col, seats) {
			below++
		}
		if col < len(seats[row]) - 1 {
			if checkBelowRight(row + 1, col + 1, seats) {
				belowRight++
			}
		}
	}

	return aboveLeft + above + aboveRight + left + right + belowLeft + below + belowRight
}

func checkAbove(row, col int, seats []string) bool {
	if row == 0 || string(seats[row][col]) != "." {
		return seatOccupied(col, seats[row])
	}
	return checkAbove(row - 1, col, seats)
}

func checkBelow(row, col int, seats []string) bool {
	if row == len(seats) - 1 || string(seats[row][col]) != "." {
		return seatOccupied(col, seats[row])
	}
	return checkBelow(row + 1, col, seats)
}

func checkLeft(col int, seats string) bool {
	if col == 0 || string(seats[col]) != "." {
		return seatOccupied(col, seats)
	}
	return checkLeft(col - 1, seats)
}

func checkRight(col int, seats string) bool {
	if col == len(seats) - 1 || string(seats[col]) != "." {
		return seatOccupied(col, seats)
	}
	return checkRight(col + 1, seats)
}

func checkAboveLeft(row, col int, seats []string) bool {
	if row == 0 || col == 0 || string(seats[row][col]) != "." {
		return seatOccupied(col, seats[row])
	}
	return checkAboveLeft(row - 1, col - 1, seats)
}

func checkAboveRight(row, col int, seats []string) bool {
	if row == 0 || col == len(seats[row]) - 1 || string(seats[row][col]) != "." {
		return seatOccupied(col, seats[row])
	}
	return checkAboveRight(row - 1, col + 1, seats)
}

func checkBelowLeft(row, col int, seats []string) bool {
	if row == len(seats) - 1 || col == 0 || string(seats[row][col]) != "." {
		return seatOccupied(col, seats[row])
	}
	return checkBelowLeft(row + 1, col - 1, seats)
}

func checkBelowRight(row, col int, seats []string) bool {
	if row == len(seats) - 1 || col == len(seats[row]) - 1 || string(seats[row][col]) != "." {
		return seatOccupied(col, seats[row])
	}
	return checkBelowRight(row + 1, col + 1, seats)
}

func seatOccupied(seatNumber int, seats string) bool {
	if string(seats[seatNumber]) == "#" {
		return true
	}
	return false
}