package piscine

func Join(strs []string, sep string) string {
	concat := ""
	for i, join := range strs {
		concat += join
		if i != len(strs)-1 {
			concat += sep
		}
	}
	return concat
}
