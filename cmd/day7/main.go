package main

import (
	"fmt"
	"github.com/bowmanrogere/advent-of-code-2020/internal"
	"log"
	"regexp"
	"strings"
)

var (
	regex = regexp.MustCompile("^([0-9]+) ([a-zA-z\\s]+) bags?$")
)

func main() {
	lines, err := internal.ReadFile("rules.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle1(lines)
}

func puzzle1(lines []string) {
	bags := make(map[string][]string)

	for _, line := range lines {
		parts := strings.Split(line, " bags contain ")
		bags[parts[0]] = strings.Split(strings.ReplaceAll(strings.ReplaceAll(parts[1], ".", ""), "no other bags", ""), ", ")
	}

	contains := make([]string, 0)
	doesntContain := make([]string, 0)
	for bag := range bags {
		if canContainShinyGoldBag(bags, bag) {
			contains = append(contains, bag)
		} else {
			doesntContain = append(doesntContain, bag)
		}
	}

	println(fmt.Sprintf("Puzzle #1: %d", len(contains)))
}

func canContainShinyGoldBag(bags map[string][]string, color string) bool {
	var canContain bool

	for _, bag := range bags[color] {
		if bag == "" {
			canContain = false
		} else {
			matches := regex.FindStringSubmatch(bag)

			if len(matches) == 0 {
				canContain = false
			} else if matches[2] == "shiny gold" {
				canContain = true
				break
			} else {
				canContain = canContainShinyGoldBag(bags, matches[2])
				if canContain {
					break
				}
			}
		}

	}

	return canContain
}
