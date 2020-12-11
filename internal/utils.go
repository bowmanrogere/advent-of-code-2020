package internal

import "strconv"

func ContainsString(list []string, s string) bool {
	for _, l := range list {
		if l == s {
			return true
		}
	}
	return false
}

func ContainsInt(list []int, i int) bool {
	for _, l := range list {
		if l == i {
			return true
		}
	}
	return false
}

func ConvertToInts(list []string) []int {
	ints := make([]int, 0)
	for _, str := range list {
		i, _ := strconv.Atoi(str)
		ints = append(ints, i)
	}
	return ints
}