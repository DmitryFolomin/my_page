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
	fmt.Scanf("%[^\n]", &input) // используем Scanf для считывания строки с пробелами

	// Убираем лишние пробелы и разделяем входные данные
	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Некорректный ввод. Используйте формат: a OP b.")
	}

	aStr, op, bStr := parts[0], parts[1], parts[2]

	var a, b int
	var isRoman bool

	// Определяем числа
	if num, err := strconv.Atoi(aStr); err == nil { // Попытка парсинга арабского числа
		a = num
	} else if num, exists := romanNumerals[aStr]; exists { // Попытка парсинга римского числа
		a = num
		isRoman = true
	} else {
		panic("Недопустимое число для a.")
	}

	if num, err := strconv.Atoi(bStr); err == nil { // Попытка парсинга арабского числа
		b = num
	} else if num, exists := romanNumerals[bStr]; exists { // Попытка парсинга римского числа
		b = num
		isRoman = true
	} else {
		panic("Недопустимое число для b.")
	}

	if a < 1 || a > 10 || b < 1 || b > 10 { // Проверка диапазона
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

	if isRoman && result < 1 { // Проверка результата для римских чисел
		panic("Результат римского числа не может быть меньше 1.")
	}

	if isRoman {
		fmt.Printf("Результат: %s\n", toRoman(result))
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}