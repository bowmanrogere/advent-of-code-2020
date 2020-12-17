package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/bowmanrogere/advent-of-code-2020/internal"
)

func main() {
	filename := fmt.Sprintf("%s/Development/advent-of-code-2020/cmd/day14/mask-input.txt", os.Getenv("HOME"))

	instructions, err := internal.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	puzzle1(instructions)
	puzzle2(instructions)
}

func puzzle1(instructions []string) {
	memRegex := regexp.MustCompile(`^mem\[([0-9]+)]$`)

	memory := make(map[int]int)

	mask := ""

	for _, inst := range instructions {
		instParts := strings.Split(inst, " = ")

		if instParts[0] == "mask" {
			mask = instParts[1]
		} else {
			// apply mask to int
			num, _ := strconv.Atoi(instParts[1])
			binary := strconv.FormatInt(int64(num), 2)
			padded := fmt.Sprintf("%036s", binary)

			tmp := ""

			for i := 0; i < len(mask); i++ {
				if string(mask[i]) == "1" {
					tmp += "1"
				} else if string(mask[i]) == "0" {
					tmp += "0"
				} else {
					tmp += string(padded[i])
				}
			}

			matches := memRegex.FindStringSubmatch(instParts[0])
			index, _ := strconv.Atoi(matches[1])
			newNum, _ := strconv.ParseInt(tmp, 2, 0)
			memory[index] = int(newNum)
		}
	}

	sum := 0
	for _, v := range memory {
		sum += v
	}

	println(fmt.Sprintf("Puzzle #1: %d", sum))
}

func puzzle2(instructions []string) {
	memRegex := regexp.MustCompile(`^mem\[([0-9]+)]$`)

	memory := make(map[int]int)

	mask := ""

	for _, inst := range instructions {
		instParts := strings.Split(inst, " = ")

		if instParts[0] == "mask" {
			mask = instParts[1]
		} else {
			num, _ := strconv.Atoi(instParts[1])
			initMemAddrStr := memRegex.FindStringSubmatch(instParts[0])[1]
			initMemAddr, _ := strconv.Atoi(initMemAddrStr)
			binaryMemAddr := strconv.FormatInt(int64(initMemAddr), 2)
			padded := fmt.Sprintf("%036s", binaryMemAddr)

			tmp := ""
			xs := ""
			for i := 0; i < len(mask); i++ {
				if string(mask[i]) == "0" {
					tmp += string(padded[i])
				} else {
					tmp += string(mask[i])
					if string(mask[i]) == "X" {
						xs += "X"
					}
				}
			}

			for i := 0; i < int(math.Pow(2, float64(len(xs)))); i++ {
				b := strconv.FormatInt(int64(i), 2)
				f := fmt.Sprintf("%%0%ds", len(xs))
				p := fmt.Sprintf(f, b)

				xCount := 0

				newMemAddrStr := ""

				for j := 0; j < len(tmp); j++ {
					if string(tmp[j]) == "X" {
						newMemAddrStr += string(p[xCount])
						xCount++
					} else {
						newMemAddrStr += string(tmp[j])
					}
				}

				newNum, _ := strconv.ParseInt(newMemAddrStr, 2, 0)
				memory[int(newNum)] = num
			}
		}
	}

	sum := 0
	for _, v := range memory {
		sum += v
	}

	println(fmt.Sprintf("Puzzle #2: %d", sum))
}
