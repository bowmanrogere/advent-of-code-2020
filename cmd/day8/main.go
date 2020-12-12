package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bowmanrogere/advent-of-code-2020/internal"
)

var (
	acc = 0
)

func main() {
	lines, err := internal.ReadFile("instructions.txt")
	if err != nil {
		log.Fatal(err)
	}

	puzzle1(lines)
	puzzle2(lines)
}

func puzzle1(lines []string) {
	executed := make([]int, 0)

	index := 0
	acc = 0

	for index < len(lines) {
		if internal.ContainsInt(executed, index) {
			break
		}

		newIndex := executeInstruction(lines[index], index)
		executed = append(executed, index)
		index = newIndex
	}

	println(fmt.Sprintf("Puzzle #1 - Acc: %d", acc))
}

func puzzle2(lines []string) {
	successful := false

	for idx, line := range lines {
		if strings.HasPrefix(line, "jmp") || strings.HasPrefix(line, "nop") {
			instructionsCopy := make([]string, len(lines))
			copy(instructionsCopy, lines)

			if strings.HasPrefix(line, "jmp") {
				instructionsCopy[idx] = strings.ReplaceAll(line, "jmp", "nop")
			}

			if strings.HasPrefix(line, "nop") {
				instructionsCopy[idx] = strings.ReplaceAll(line, "nop", "jmp")
			}

			history := make([]int, 0)
			index := 0
			loopDetected := false
			acc = 0

			for index < len(instructionsCopy) {
				if internal.ContainsInt(history, index) {
					loopDetected = true
					break
				}

				newIndex := executeInstruction(instructionsCopy[index], index)
				history = append(history, index)
				index = newIndex
			}

			successful = !loopDetected
		}

		if successful {
			break
		}
	}

	println(fmt.Sprintf("Puzzle #2 - Acc: %d", acc))
}

func executeInstruction(instruction string, index int) int {
	parts := strings.Split(instruction, " ")

	switch parts[0] {
	case "acc":
		num, _ := strconv.Atoi(parts[1])
		acc += num
		index++
	case "jmp":
		num, _ := strconv.Atoi(parts[1])
		index += num
	case "nop":
		index++
	}

	return index
}
