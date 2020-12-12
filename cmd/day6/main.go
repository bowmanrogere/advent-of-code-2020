package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bowmanrogere/advent-of-code-2020/internal"
)

func main() {
	lines, err := internal.ReadFile("answers.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle1(lines)
	puzzle2(lines)
}

func puzzle1(lines []string) {
	uniqueAnswers := make([]string, 0)
	yeses := ""

	for idx, line := range lines {
		if line != "" {
			for _, c := range line {
				if !strings.Contains(yeses, string(c)) {
					yeses += string(c)
				}
			}
		}

		if line == "" || idx == len(lines)-1 {
			uniqueAnswers = append(uniqueAnswers, yeses)
			yeses = ""
		}
	}

	count := sumAnswers(uniqueAnswers)

	println(fmt.Sprintf("Puzzle 1 sum of counts: %d", count))
}

func puzzle2(lines []string) {
	groupAnswers := make([]string, 0)
	allAnswerYesByGroup := make([]string, 0)

	for idx, line := range lines {
		if line != "" {
			groupAnswers = append(groupAnswers, line)
		}

		if line == "" || idx == len(lines)-1 {
			allAnswerYes := ""

			for _, c := range groupAnswers[0] {
				if allContains(groupAnswers, string(c)) {
					allAnswerYes += string(c)
				}
			}

			allAnswerYesByGroup = append(allAnswerYesByGroup, allAnswerYes)
			groupAnswers = make([]string, 0)
		}
	}

	count := sumAnswers(allAnswerYesByGroup)

	println(fmt.Sprintf("Puzzle2 sum of counts: %d", count))
}

func sumAnswers(answers []string) int {
	count := 0
	for _, answer := range answers {
		count += len(answer)
	}

	return count
}

func allContains(answers []string, char string) bool {
	for _, answer := range answers {
		if !strings.Contains(answer, char) {
			return false
		}
	}

	return true
}
