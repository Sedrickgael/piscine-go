package piscine

func NRune(str string) rune {
	lastrune := []rune(str)
  lens := len([]rune(str))
	return lastrune[lens-1]
}

