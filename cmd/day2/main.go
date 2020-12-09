package main

import (
	"fmt"
	"github.com/bowmanrogere/advent-of-code-2020/internal"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines, err := internal.ReadFile("passwords.txt")
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	puzzle1(lines)
	puzzle2(lines)
}

func puzzle1(lines []string) {
	validPasswords := 0
	for _, line := range lines {
		password, char, min, max := splitLine(line)
		num := strings.Count(password, char)
		if min <= num && num <= max {
			validPasswords++
		}
	}

	println(fmt.Sprintf("Puzzle 1 # of Valid Passwords: %d", validPasswords))
}

func puzzle2(lines []string) {
	validPasswords := 0
	for _, line := range lines {
		password, char, pos1, pos2 := splitLine(line)
		check1 := string(password[pos1-1])
		check2 := string(password[pos2-1])
		if (char == check1 && char != check2) || (char != check1 && char == check2) {
			validPasswords++
		}
	}

	println(fmt.Sprintf("Puzzle 2 # of Valid Passwords: %d", validPasswords))
}

func splitLine(line string) (string, string, int, int) {
	parts := strings.Split(line, " ")
	password := parts[2]
	char := strings.ReplaceAll(parts[1], ":", "")
	min, _ := strconv.Atoi(strings.Split(parts[0], "-")[0])
	max, _ := strconv.Atoi(strings.Split(parts[0], "-")[1])
	return password, char, min, max
}
