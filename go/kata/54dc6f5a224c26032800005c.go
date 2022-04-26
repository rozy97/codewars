package kata

import (
	"fmt"
	"strconv"
	"strings"
)

//kyu 6
func StockList(books []string, categories []string) string {
	var result string
	if len(categories) == 0 || len(books) == 0 {
		return result
	}

	for i, category := range categories {
		total := 0
		for _, book := range books {
			if string(book[0]) == category {
				stock, _ := strconv.Atoi(strings.Split(book, " ")[1])
				total += stock
			}
		}

		result += fmt.Sprintf("(%v : %v)", category, total)
		if i != len(categories)-1 {
			result += " - "
		}

	}

	return result
}
