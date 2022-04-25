package kata

// kyu 8
func CountSheeps(numbers []bool) (total int) {
	for _, v := range numbers {
		if v {
			total++
		}
	}
	return
}
