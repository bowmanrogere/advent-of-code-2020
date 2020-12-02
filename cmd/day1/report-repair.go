package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readExpenseReport() []int {
	values := make([]int, 0)

	f, err := os.Open("expense-report.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
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
