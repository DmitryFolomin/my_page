package main

import (
	"fmt"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func toRoman(num int) string {
	for k, v := range romanNumerals {
		if v == num {
			return k
		}
	}
	return ""
}

func main() {
	var input string
	fmt.Println("Введите выражение (например, 3 + 2 или II + I):")
	fmt.Scanln(input)

	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Некорректный ввод.")
	}

	aStr, op, bStr := parts[0], parts[1], parts[2]

	var a, b int

	if num, err := strconv.Atoi(aStr); err == nil {
		a = num
	} else if num, exists := romanNumerals[aStr]; exists {
		a = num
	} else {
		panic("Недопустимое число.")
	}

	if num, err := strconv.Atoi(bStr); err == nil {
		b = num
	} else if num, exists := romanNumerals[bStr]; exists {
		b = num
	} else {
		panic("Недопустимое число.")
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Числа должны быть от 1 до 10 включительно.")
	}

	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			panic("Деление на ноль.")
		}
		result = a / b
	default:
		panic("Недопустимая операция.")
	}

	if result <= 0 && (aStr != "" && bStr != "" && strings.ToUpper(aStr[0:1]) != "I" && strings.ToUpper(bStr[0:1]) != "I") {
		fmt.Printf("Результат: %d\n", result)
	} else if result <= 0 {
		panic("Результат римского числа не может быть меньше 1.")
	} else {
		fmt.Printf("Результат: %s\n", toRoman(result))
	}
}
