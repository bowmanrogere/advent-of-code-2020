package main

import (
	"fmt"
	"github.com/bowmanrogere/advent-of-code-2020/internal"
	"log"
	"sort"
)

func main() {
	lines, err := internal.ReadFile("code.txt")
	if err != nil {
		log.Fatal(err)
	}

	codes := internal.ConvertToInts(lines)
	invalidCode := findInvalidCode(codes)

	println(fmt.Sprintf("Puzzle 1: %d", invalidCode))
	println(fmt.Sprintf("Puzzle 2: %d", findEncryptionWeakness(codes, invalidCode)))
}

func findInvalidCode(codes []int) int {
	preamble := 25
	invalidCode := 0
	for i := preamble; i < len(codes); i++ {
		foundSumInPrevious5 := false

		for j := i - preamble; j <= i; j++ {
			for k := i - preamble; k <= i; k++ {
				if codes[j]+codes[k] == codes[i] {
					foundSumInPrevious5 = true
					break
				}
			}
			if foundSumInPrevious5 {
				break
			}
		}

		if !foundSumInPrevious5 {
			invalidCode = codes[i]
			break
		}
	}
	return invalidCode
}

func findEncryptionWeakness(codes []int, invalidCode int) int {
	encryptionWeakness := -1

	for i := 0; i < len(codes); i++ {
		sum := 0
		ints := make([]int, 0)
		for j := i; j < len(codes); j++ {
			sum += codes[j]
			ints = append(ints, codes[j])
			sort.Ints(ints)
			if sum < invalidCode {
				continue
			}
			if sum > invalidCode {
				break
			}
			encryptionWeakness = ints[0] + ints[len(ints) - 1]
			break
		}
		if encryptionWeakness != -1 {
			break
		}
	}

	return encryptionWeakness
}