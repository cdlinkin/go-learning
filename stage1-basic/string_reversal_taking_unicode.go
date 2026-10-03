package stage1

import "fmt"

func StringReversalTakingUnicode(str string) {
	stringToRunes := []rune(str)
	var temp rune
	left := 0
	right := len(stringToRunes) - 1
	for left < right {
		temp = stringToRunes[left]
		stringToRunes[left] = stringToRunes[right]
		stringToRunes[right] = temp

		left++
		right--
	}

	runesToString := string(stringToRunes)
	fmt.Println(runesToString)
}
