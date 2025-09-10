package main

import (
	"fmt"
)

const USD = 30
const EUR = 40
const RUB = 1

func main() {

	for {
		number, source, target := userInput()

		res := convert(number, source, target)
		fmt.Println(res)

	}
}

func convert(howMuch float64, source, target string) float64 {
	if source == "EUR" && target == "RUB" {
		return howMuch * 40
	}
	if source == "USD" && target == "RUB" {
		return howMuch * 30
	}
	if source == "USD" && target == "EUR" {
		return howMuch * 1.10
	}
	if source == "EUR" && target == "USD" {
		return howMuch * 0.90
	}
	return howMuch
}

func userInput() (float64, string, string) {
	var number float64
	var source, target string
	fmt.Println("Введите исходную валюту:")
	fmt.Scan(&source)
	fmt.Println("Введите количество:")
	fmt.Scan(&number)
	fmt.Println("Введите целевую валюту:")
	fmt.Scan(&target)
	return number, source, target
}
