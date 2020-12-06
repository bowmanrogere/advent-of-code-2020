package main

import (
	"fmt"
	"github.com/bowmanrogere/advent-of-code-2020/internal"
	"log"
	"strings"
)

func main() {
	lines, err := internal.ReadFile("answers.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle1(lines)
}

func puzzle1(lines []string) {
	groupYeses := make([]string, 0)
	yeses := ""
	for idx, line := range lines {
		if line != "" {
			for _, c := range line {
				if !strings.Contains(yeses, string(c)) {
					yeses += string(c)
				}
			}
		}

		if line == "" || idx == len(lines) - 1 {
			groupYeses = append(groupYeses, yeses)
			yeses = ""
		}
	}

	count := 0
	for _, answer := range groupYeses {
		count += len(answer)
	}

	println(fmt.Sprintf("Sum of counts: %d", count))
}