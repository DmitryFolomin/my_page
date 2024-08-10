package main

import (
"fmt"
"os"
"bufio"
"strings"
"strconv"
)

func romanToArabic(roman string) int {
romanMap := map[string]int{"I":1, "II":2, "III":3, "IV":4, "V":5, "VI":6, "VII":7, "VIII":8, "IX":9, "X":10}
return romanMap[roman]
}

func arabicToRoman(arabic int) string {
if arabic < 1 || arabic > 3999 {
panic("Invalid input")
}

go
Copy code
vals := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
romans := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}

romanNum := ""

for i := 0; i < 13; i++ {
	for arabic >= vals[i] {
		romanNum += romans[i]
		arabic -= vals[i]
	}
}

return romanNum
}

type Calculator struct {
}

func (c Calculator) Add(a, b int) int {
return a + b
}

func (c Calculator) Subtract(a, b int) int {
return a - b
}

func (c Calculator) Multiply(a, b int) int {
return a * b
}

func (c Calculator) Divide(a, b int) int {
if b == 0 {
panic("Division by zero")
}
return a / b
}

func main() {
calc := Calculator{}

go
Copy code
reader := bufio.NewReader(os.Stdin)
fmt.Println("Enter expression (e.g. 2 + 3):")
expression, _ := reader.ReadString('\n')

expression = strings.TrimSpace(expression)
parts := strings.Split(expression, " ")

if len(parts) != 3 {
	panic("Invalid input")
}

a, err1 := strconv.Atoi(parts[0])
b, err2 := strconv.Atoi(parts[2])

if err1 != nil || err2 != nil {
	panic("Invalid input")
}

result := 0

switch parts[1] {
case "+":
	result = calc.Add(a, b)
case "-":
	result = calc.Subtract(a, b)
case "*":
	result = calc.Multiply(a, b)
case "/":
	result = calc.Divide(a, b)
default:
	panic("Invalid operation")
}

fmt.Println("Result:", result)
}

