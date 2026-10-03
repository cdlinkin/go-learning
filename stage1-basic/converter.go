package stage1

import (
	"fmt"
	"os"
	"strconv"
)

func ConverterToFahrenheit() (int, error) {
	t := os.Args

	if len(t) < 2 {
		return 0, fmt.Errorf("the temperature is not specified.")
	}

	fmt.Println("enter the temperature in Celsius.:", t[1])

	temp, err := strconv.Atoi(t[1])
	if err != nil {
		return 0, fmt.Errorf("failed to convert the temperature: %v", err)
	}

	f := (temp * 9 / 5) + 32

	return f, nil
}
