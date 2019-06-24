package piscine

func Compact(ptr *[]string , leng int) int {
	count:=0
	for i:=0;i<leng;i++{
		if (*ptr)[i]!=" "{
			count++
		}
	}
	return count
}
