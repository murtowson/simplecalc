package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func romanToInt(s string) (int, error) {
	roman := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	value, exists := roman[s]
	if !exists {
		return 0, errors.New("Выдача паники введена цифра за диапазоном римских чисел с которыми я могу работать")
	}
	return value, nil
}

func intToRoman(num int) (string, error) {
	if num < 1 {
		return "", errors.New("Выдача паники цифра меньше I")
	}
	roman := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	if num <= 10 {
		return roman[num-1], nil
	}
	return "", errors.New("Выдача паники результат за прделом диапозона римских цифр с котрыми я могу раьотать ")
}


func isRoman(s string) bool {
	_, err := romanToInt(s)
	return err == nil
}

func main() {
	var input string
	fmt.Println("Введите выражение (например, 1 + 2 или VI / III):")
	fmt.Scanln(&input)

	input = strings.ReplaceAll(input, " ", "")

	var operation string
	if strings.Contains(input, "+") {
		operation = "+"
	} else if strings.Contains(input, "-") {
		operation = "-"
	} else if strings.Contains(input, "*") {
		operation = "*"
	} else if strings.Contains(input, "/") {
		operation = "/"
	} else {
		panic("Некорректная операция")
	}

	parts := strings.Split(input, operation)
	if len(parts) != 2 {
		panic("Некорректный формат ввода")
	}

	operand1, operand2 := parts[0], parts[1]


	isRoman1 := isRoman(operand1)
	isRoman2 := isRoman(operand2)

	if isRoman1 != isRoman2 {
		panic("Выдача паники: используются одновременно разные системы счисления")
	}

	var num1, num2 int
	var err error

	if isRoman1 {
		num1, err = romanToInt(operand1)
		if err != nil {
			panic(err)
		}
		num2, err = romanToInt(operand2)
		if err != nil {
			panic(err)
		}
	} else {
		num1, err = strconv.Atoi(operand1)
		if err != nil {
			panic("Некорректное арабское число")
		}
		num2, err = strconv.Atoi(operand2)
		if err != nil {
			panic("Некорректное арабское число")
		}
		if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
			panic("Число выходит за пределы допустимого диапазона (1-10)")
		}
	}


	var result int
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			panic("Выдача паники: Деление на ноль")
		}
		result = num1 / num2
	}

	if isRoman1 {
		romanResult, err := intToRoman(result)
		if err != nil {
			panic(err)
		}
		fmt.Println("Результат:", romanResult)
	} else {
		fmt.Println("Результат:", result)
	}
}
