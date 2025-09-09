package main

import "fmt"

func main() {
	const USD = 30
	const EUR = 40
	const RUB = 20
	fmt.Println(EUR * RUB)
}

func convert(a float64, b, c string) {

}

func userInput() (float64, string, string) {
	var a float64
	var b, c string
	fmt.Scan(&a, &b, &c)
	return a, b, c
}
