package main

import (
	"bufio"
	"fmt"
	"os"
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например, 3 + 2 или II + I):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Разделяем входные данные
	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Некорректный ввод. Используйте формат: a OP b.")
	}
	aStr, op, bStr := parts[0], parts[1], parts[2]

	var a, b int
	var isRomanA, isRomanB bool

	// Определяем первое число
	if num, err := strconv.Atoi(aStr); err == nil {
		a = num
	} else if num, exists := romanNumerals[aStr]; exists {
		a = num
		isRomanA = true
	} else {
		panic("Недопустимое число для a.")
	}

	// Определяем второе число
	if num, err := strconv.Atoi(bStr); err == nil {
		b = num
	} else if num, exists := romanNumerals[bStr]; exists {
		b = num
		isRomanB = true
	} else {
		panic("Недопустимое число для b.")
	}

	// Проверка смешанных систем счисления
	if isRomanA != isRomanB {
		panic("Используются одновременно разные системы счисления.")
	}

	// Проверка допустимых чисел (1-10 для римских и положительных для арабских)
	if (isRomanA && (a < 1 || a > 10 || b < 1 || b > 10)) || (!isRomanA && (a < 1 || b < 1)) {
		panic("Числа должны быть от 1 до 10 включительно для римских чисел.")
	}

	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
		if result < 1 && isRomanA {
			panic("Результат не может быть отрицательным в римской системе.")
		}
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

	// Вывод результата
	if isRomanA {
		if result < 1 {
			panic("Результат не может быть отрицательным в римской системе.")
		}
		if result > 10 {
			panic("Результат больше 10 и не может быть представлен римскими числами.")
		}
		fmt.Printf("Результат: %s\n", toRoman(result))
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}
