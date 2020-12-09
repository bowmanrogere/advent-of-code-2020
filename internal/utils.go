package internal

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