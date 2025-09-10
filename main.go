package main

import (
	"fmt"
)

const USD = 30
const EUR = 40
const RUB = 1

func main() {

	for {
		howMuch, from, to, err := userInput()
		if err != nil {
			break
		}
		res := convert(howMuch, from)
		fmt.Println(res)

	}
}

func convert(howMuch float64, from string) float64 {
	switch from {
	case "EUR":
		howMuch *= 40
	case "USD":
		howMuch *= 30
	}
	return howMuch
}

func userInput() (float64, string, string, error) {
	fmt.Println("Введите количество валюты  и ее тип (USD/UER)")
	var number float64
	var source, target string
	_, err := fmt.Scan(&number, &source, &target)
	if err != nil {
		return number, source, target, err
	}
	return number, source, target, err
}
