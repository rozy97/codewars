package kata

// kyu 7
func MxDifLg(a1 []string, a2 []string) (maxLength int) {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	for _, x := range a1 {
		for _, y := range a2 {
			length := abs(len(x) - len(y))
			if length > maxLength {
				maxLength = length
			}
		}
	}
	return
}

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return num * -1
}
