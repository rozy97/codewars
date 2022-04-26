package kata

// kyu 5
func alphanumeric(str string) (result bool) {
	if len(str) == 0 {
		return
	}

	for _, v := range str {
		if !((v >= 48 && v <= 57) || (v >= 65 && v <= 90) || (v >= 97 && v <= 122)) {
			return
		}
	}

	return true
}
