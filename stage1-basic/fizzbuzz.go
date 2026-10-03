package stage1

import "fmt"

type Rule struct {
	Number int
	Word   string
}

func FizzBuzz() {
	rules := []Rule{
		{
			Number: 3,
			Word:   "Fizz",
		},
		{
			Number: 5,
			Word:   "Buzz",
		},
		{
			Number: 7,
			Word:   "Woof",
		},
	}
	for i := 1; i <= 100; i++ {
		var result string
		for _, rule := range rules {
			if i%rule.Number == 0 {
				result += rule.Word
			}
		}
		if result == "" {
			fmt.Println(i)
		} else {
			fmt.Println(result)
		}
	}
}
