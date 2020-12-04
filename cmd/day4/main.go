package main

import (
	"fmt"
	"github.com/bowmanrogere/advent-of-code-2020/internal"
	"log"
)

func main() {
	lines, err := internal.ReadFile("passports.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle1(lines)
}

func puzzle1(lines []string) {
	validPassports := 0

	passport := &internal.Passport{}
	for idx, line := range lines {
		if line != "" {
			passport.AddInformation(line)
		}

		if line == "" || idx == len(lines) - 1 {
			if passport.Valid() {
				validPassports++
			}
			passport = &internal.Passport{}
		}
	}

	println(fmt.Sprintf("Puzzle 1 # valid passports: %d", validPassports))
}