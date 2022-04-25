package kata

import "fmt"

// kyu 8
func MultiTable(number int) (result string) {

	lastMultiply := 10

	for i := 1; i <= lastMultiply; i++ {
		result += fmt.Sprintf("%v * %v = %v", i, number, i*number)
		if i != lastMultiply {
			result += "\n"
		}
	}

	return result
}
