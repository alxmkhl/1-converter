package main

import (
	"errors"
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
		fmt.Println(howMuch, from, to)

	}
}

func convert(howMuch float64, from, to string) {
	switch from {
	case "EUR":

	}
}

func userInput() (float64, string, string, error) {
	fmt.Println("Введите количество валюты  и ее тип (USD/UER)")
	var a float64
	var b, c string
	_, err := fmt.Scan(&a, &b, &c)
	if err != nil {
		return _, _, _, errors.New("Ошибка ввода")
	}
	return a, b, c, _
}
