go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"I":  1,
	"II": 2,
	"III": 3,
	"IV":  4,
	"V":   5,
	"VI":  6,
	"VII": 7,
	"VIII": 8,
	"IX":  9,
	"X":   10,
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
	fmt.Scanln(&input) // Изменил на Scanln

	// Убираем лишние пробелы и разделяем входные данные
	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Некорректный ввод. Используйте формат: a OP b.")
	}

	aStr, op, bStr := parts[0], parts[1], parts[2]

	var a, b int
	var isRoman bool

	// Определяем первое число
	if num, err := strconv.Atoi(aStr); err == nil {
		a = num
	} else if num, exists := romanNumerals[aStr]; exists {
		a = num
		isRoman = true
	} else {
		panic("Недопустимое число для a.")
	}

	// Определяем второе число
	if num, err := strconv.Atoi(bStr); err == nil {
		b = num
	} else if num, exists := romanNumerals[bStr]; exists {
		b = num
		isRoman = true
	} else {
		panic("Недопустимое число для b.")
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

	if isRoman && result < 1 {
		panic("Результат римского числа не может быть меньше 1.")
	}

	if isRoman {
		fmt.Printf("Результат: %s\n", toRoman(result))
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}