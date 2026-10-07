/*
exercise:

	First, write the tests, then the code for the ParseDuration("1h30m")
	function (a simplified parser: hours, minutes, seconds)
	At least 12 table cases, including erroneous input
*/
package duration

import (
	"fmt"
	"strconv"
	"time"
)

func ParseDuration(input string) (time.Duration, error) {
	if input == "" {
		return 0, fmt.Errorf("duration is empty")
	}

	var res time.Duration
	var num string

	for i := 0; i < len(input); i++ {
		char := input[i]

		if char >= '0' && char <= '9' {
			num += string(char)
			continue
		}

		if num == "" {
			return 0, fmt.Errorf("expected is number")
		}

		v, err := strconv.Atoi(num)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %w", err)
		}

		switch char {
		case 'h':
			res += time.Duration(v) * time.Hour

		case 'm':
			res += time.Duration(v) * time.Minute

		case 's':
			res += time.Duration(v) * time.Second

		default:
			return 0, fmt.Errorf("unknown: %c", char)
		}

		num = ""
	}

	if num != "" {
		return 0, fmt.Errorf("expected unit")
	}

	return res, nil
}
