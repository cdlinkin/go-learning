package stage1

import (
	"strings"
)

func CountWords(text string) map[string]int {
	counter := make(map[string]int)

	fieldText := strings.Fields(text)
	for _, v := range fieldText {
		counter[v]++
	}

	return counter
}
