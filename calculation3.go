package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"I":     1,
	"II":    2,
	"III":   3,
	"IV":    4,
	"V":     5,
	"VI":    6,
	"VII":   7,
	"VIII":  8,
	"IX":    9,
	"X":     10,
	"XI":    11,
	"XII":   12,
	"XIII":  13,
	"XIV":   14,
	"XV":    15,
	"XVI":   16,
	"XVII":  17,
	"XVIII": 18,
	"XIX":   19,
	"XX":    20,
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
		fmt.Println("Некорректный ввод. Используйте формат: a OP b.")
		return
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
		fmt.Println("Недопустимое число для a.")
		return
	}

	// Определяем второе число
	if num, err := strconv.Atoi(bStr); err == nil {
		b = num
	} else if num, exists := romanNumerals[bStr]; exists {
		b = num
		isRoman = true
	} else {
		fmt.Println("Недопустимое число для b.")
		return
	}

	if a < 1 || a > 20 || b < 1 || b > 20 {
		fmt.Println("Числа должны быть от 1 до 20 включительно.")
		return
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
			fmt.Println("Деление на ноль.")
			return
		}
		result = a / b
	default:
		fmt.Println("Недопустимая операция.")
		return
	}

	if isRoman && result < 1 {
		fmt.Println("Результат римского числа не может быть меньше 1.")
		return
	}

	if isRoman {
		fmt.Printf("Результат: %s\n", toRoman(result))
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}
