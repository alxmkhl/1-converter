package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	USD = "USD"
	EUR = "EUR"
	RUB = "RUB"
)

var rates = map[string]map[string]float64{
	USD: {USD: 1.0, EUR: 0.90, RUB: 30.0},
	EUR: {EUR: 1.0, USD: 1.10, RUB: 40.0},
	RUB: {RUB: 1.0, USD: 1.0 / 30.0, EUR: 1.0 / 40.0},
}

var allowed = []string{USD, EUR, RUB}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("-------- Конвертер валют --------")

		src := promptCurrency(reader, "Введите исходную валюту (USD, EUR, RUB) или q для выхода:")
		if src == "Q" {
			fmt.Println("Выход.")
			return
		}

		amount := promptAmount(reader, "Введите количество (положительное число):")

		tgt := promptCurrency(reader, "Введите целевую валюту (USD, EUR, RUB) или q для выхода:")
		if tgt == "Q" {
			fmt.Println("Выход.")
			return
		}

		res, err := convert(amount, src, tgt)
		if err != nil {
			fmt.Println("Ошибка конвертации:", err)
			continue
		}

		fmt.Printf("%.2f %s = %.2f %s\n\n", amount, src, res, tgt)
	}
}

func promptCurrency(r *bufio.Reader, msg string) string {
	for {
		fmt.Println(msg)
		raw, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода, попробуйте ещё раз.")
			continue
		}
		val := strings.ToUpper(strings.TrimSpace(raw))
		if val == "Q" {
			return "Q"
		}
		if validCurrency(val) {
			return val
		}
		fmt.Printf("Неверная валюта. Допустимые варианты: %s\n", strings.Join(allowed, ", "))
	}
}

func promptAmount(r *bufio.Reader, msg string) float64 {
	for {
		fmt.Println(msg)
		var num float64
		_, err := fmt.Fscanln(r, &num)
		if err != nil {
			// очистить плохую строку из буфера
			_, _ = r.ReadString('\n')
			fmt.Println("Неверный формат числа. Повторите ввод.")
			continue
		}
		if num <= 0 {
			fmt.Println("Число должно быть положительным.")
			continue
		}
		// съесть остаток строки (если есть)
		_, _ = r.ReadString('\n')
		return num
	}
}

func validCurrency(s string) bool {
	for _, v := range allowed {
		if v == s {
			return true
		}
	}
	return false
}

func convert(howMuch float64, source, target string) (float64, error) {
	srcMap, ok := rates[source]
	if !ok {
		return 0, fmt.Errorf("неизвестная исходная валюта: %s", source)
	}
	rate, ok := srcMap[target]
	if !ok {
		return 0, fmt.Errorf("конвертация из %s в %s не поддерживается", source, target)
	}
	return howMuch * rate, nil
}
