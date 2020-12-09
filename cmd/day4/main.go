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

	println(fmt.Sprintf("Puzzle 1 # valid passports: %d", countValidPassports(lines, false)))
	println(fmt.Sprintf("Puzzle 2 # valid passports: %d", countValidPassports(lines, true)))
}

func countValidPassports(lines []string, validate bool) int {
	validPassports := 0

	passport := &internal.Passport{}
	for idx, line := range lines {
		if line != "" {
			passport.AddInformation(line, validate)
		}

		if line == "" || idx == len(lines)-1 {
			if passport.Valid() {
				validPassports++
			}
			passport = &internal.Passport{}
		}
	}

	return validPassports
}
