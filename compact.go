package piscine

func Compact(ptr *[]string, int, length int) int {
	count := 0
	for i := 0; i < length; i++ {
		if (*ptr)[i] != " " {
			count++
		}
	}
	return count
}
