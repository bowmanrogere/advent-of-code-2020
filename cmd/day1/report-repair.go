package main

import (
	"github.com/bowmanrogere/advent-of-code-2020/internal"
	"log"
	"strconv"
)

func readExpenseReport() []int {
	lines, err := internal.ReadFile("expense-report.txt")
	if err != nil {
		log.Fatal(err)
	}

	values := make([]int, 0)
	for _, line := range lines {
		val, _ := strconv.Atoi(line)
		values = append(values, val)
	}

	return values
}

func findTwo(expenseValues []int) {
	valuesFound := false
	for _, val1 := range expenseValues {
		for _, val2 := range expenseValues {
			if val1 + val2 == 2020 {
				println(val1 * val2)
				valuesFound = true
				break
			}
		}
		if valuesFound {
			break
		}
	}
}

func findThree(expenseValues []int) {
	valuesFound := false
	for _, val1 := range expenseValues {
		for _, val2 := range expenseValues {
			for _, val3 := range expenseValues {
				if val1+val2+val3 == 2020 {
					println(val1 * val2 * val3)
					valuesFound = true
					break
				}
			}
			if valuesFound {
				break
			}
		}
		if valuesFound {
			break
		}
	}
}

func main() {
	expenseValues := readExpenseReport()

	findTwo(expenseValues)

	findThree(expenseValues)
}
