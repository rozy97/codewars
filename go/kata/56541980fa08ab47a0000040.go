package kata

import "fmt"

func PrinterError(s string) string {
	var totalErr uint64
	for _, v := range s {
		if !(v >= 97 && v <= 109) {
			totalErr++
		}
	}
	return fmt.Sprintf("%v/%v", totalErr, len(s))
}
