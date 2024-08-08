package main

import (
	"fmt"
	"strconv"
)

// Преобразуем римские числа в арабские
func romanToArabic(roman string) (int, error) {
	romanNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10}
	result := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumerals[rune(roman[i])]
		if value == 0 {
			return 0, fmt.Errorf("некорректное римское число")
		}
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}
	if result < 1 || result > 10 {
		return 0, fmt.Errorf("число вне допустимого диапазона 1-10")
	}
	return result, nil
}

// Преобразуем арабские числа в римские
func arabicToRoman(arabic int) string {
	arabicNumerals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNumerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""

	for i, arabicValue := range arabicNumerals {
		for arabic >= arabicValue {
			arabic -= arabicValue
			result += romanNumerals[i]
		}
	}
	return result
}

func main() {
	var input1, input2, operation string

	fmt.Print("Введите выражение (например, 2 + 2 или II + III): ")
	fmt.Scanf("%s %s %s", &input1, &operation, &input2)

	isRoman := func(s string) bool { // функция проверки на римские числа
		for _, char := range s {
			if _, ok := map[rune]int{'I': 1, 'V': 5, 'X': 10}[char]; !ok {
				return false
			}
		}
		return true
	}

	if isRoman(input1) && isRoman(input2) { // проверяем, являются ли оба значения римскими числами
		a, err1 := romanToArabic(input1)
		b, err2 := romanToArabic(input2)
		if err1 != nil || err2 != nil {
			fmt.Println(err1.Error())
			return
		}
		processOperation(a, b, operation, true)
	} else { // обрабатываем арабские числа
		a, err1 := strconv.Atoi(input1)
		b, err2 := strconv.Atoi(input2)
		if err1 != nil || err2 != nil {
			fmt.Println("Ошибка: Некорректное арабское число.")
			return
		}
		if a < 1 || a > 10 || b < 1 || b > 10 {
			fmt.Println("Ошибка: Числа должны быть в диапазоне от 1 до 10.")
			return
		}
		processOperation(a, b, operation, false)
	}
}

func processOperation(a, b int, operation string, isRoman bool) {
	var result int
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя!")
			return
		}
		result = a / b
	default:
		fmt.Println("Неверная операция. Доступные операции: +, -, *, /.")
		return
	}

	if isRoman {
		fmt.Println(arabicToRoman(result))
	} else {
		fmt.Println(result)
	}
}
