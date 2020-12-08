package main

import (
	"fmt"
	"github.com/bowmanrogere/advent-of-code-2020/internal"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	regex = regexp.MustCompile("^([0-9]+) ([a-zA-z\\s]+) bags?$")
)

func main() {
	lines, err := internal.ReadFile("/Users/t24974a/Development/advent-of-code-2020/cmd/day7/rules.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle1(lines)
	puzzle2(lines)
}

func puzzle1(lines []string) {
	bags := getBags(lines)

	contains := make([]string, 0)
	for bag := range bags {
		if canContainShinyGoldBag(bags, bag) {
			contains = append(contains, bag)
		}
	}

	println(fmt.Sprintf("Puzzle #1: %d", len(contains)))
}

func puzzle2(lines []string) {
	bags := getBags(lines)

	count := countBags(bags, "shiny gold")

	println(fmt.Sprintf("Puzzle 2: %d", count))
}

func getBags(lines []string) map[string][]string {
	bags := make(map[string][]string)

	for _, line := range lines {
		parts := strings.Split(line, " bags contain ")
		bags[parts[0]] = strings.Split(strings.ReplaceAll(strings.ReplaceAll(parts[1], ".", ""), "no other bags", ""), ", ")
	}

	return bags
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

func countBags(bags map[string][]string, color string) int {
	count := 0
	for _, bag := range bags[color] {
		matches := regex.FindStringSubmatch(bag)

		if len(matches) > 0 {
			numBags, _ := strconv.Atoi(matches[1])
			count += numBags + numBags * countBags(bags, matches[2])
		}
	}

	return count
}
