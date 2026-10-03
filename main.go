package main

import (
	stage1 "cdlinkin/go-learning/stage1-basic"
	"fmt"
)

func main() {
	fmt.Println(stage1.CountWords("go go go Golang")) // CountWords

	f, err := stage1.ConverterToFahrenheit() // ConverterToFahrenheit
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("temperature in Fahrenheit:", f)
}
