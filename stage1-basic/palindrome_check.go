package stage1

import "unicode"

func PalindromCheck(str string) bool {
	stringToRunes := []rune(str)

	cleaned := []rune{}

	for _, r := range stringToRunes {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, unicode.ToLower(r))
		}
	}

	left := 0
	right := len(cleaned) - 1

	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}

		left++
		right--
	}

	return true
}
