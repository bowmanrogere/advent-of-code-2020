package main

import (
	"fmt"
	"github.com/bowmanrogere/advent-of-code-2020/internal"
	"log"
	"strconv"
	"strings"
)

func main() {
	puzzle1()
}

func puzzle1() {
	lines, err := internal.ReadFile("passwords.txt")
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	validPasswords := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")

		password := parts[2]
		char := strings.ReplaceAll(parts[1], ":", "")
		min, _ := strconv.Atoi(strings.Split(parts[0], "-")[0])
		max, _ := strconv.Atoi(strings.Split(parts[0], "-")[1])

		num := strings.Count(password, char)

		if min <= num && num <= max {
			validPasswords++
		}
	}

	println(fmt.Sprintf("Valid Passwords: %d", validPasswords))
}