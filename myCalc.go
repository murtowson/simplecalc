package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{
	"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10,
	"XL": 40, "L": 50, "XC": 90, "C": 100,
	"CD": 400, "D": 500, "CM": 900, "M": 1000,
}

var arabicToRoman = []struct {
	value  int
	symbol string
}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите выражение (например, 2 + 3 или VI*III), или 'E' для выхода:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ToUpper(input)

		if input == "E" {
			fmt.Println("Программа завершена.")
			break
		}

		result, err := calculate(input)
		if err != nil {
			fmt.Printf("Выдача паники: %s\n", err)
		} else {
			fmt.Println("Результат:", result)
		}
	}
}

func calculate(input string) (string, error) {
	re := regexp.MustCompile(`(?i)^([IVXLCDM]+|\d+(?:\.\d+)?)\s*([\+\-\*/])\s*([IVXLCDM]+|\d+(?:\.\d+)?)$`)
	matches := re.FindStringSubmatch(input)
	if len(matches) != 4 {
		return "", errors.New("Некорректный формат ввода")
	}

	num1, num2 := matches[1], matches[3]
	operator := matches[2]

	isRoman := isRomanNumber(num1) && isRomanNumber(num2)
	isArabic := isArabicNumber(num1) && isArabicNumber(num2)

	if !isRoman && !isArabic {
		return "", errors.New("Некорректный формат чисел")
	}

	if isRoman {
		a, err := romanToInt(num1)
		if err != nil {
			return "", err
		}
		b, err := romanToInt(num2)
		if err != nil {
			return "", err
		}
		result, err := performOperation(float64(a), float64(b), operator)
		if err != nil {
			return "", err
		}
		if result < 1 {
			return "", errors.New("В римской системе нет отрицательных чисел")
		}
		return intToRoman(int(math.Round(result))), nil
	}

	a, err := strconv.ParseFloat(num1, 64)
	if err != nil {
		return "", err
	}
	b, err := strconv.ParseFloat(num2, 64)
	if err != nil {
		return "", err
	}
	result, err := performOperation(a, b, operator)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.2f", result), nil
}

func isRomanNumber(s string) bool {
	_, err := romanToInt(s)
	return err == nil
}

func isArabicNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func romanToInt(s string) (int, error) {
	result := 0
	for i := 0; i < len(s); i++ {
		if i > 0 && romanToArabic[s[i:i+1]] > romanToArabic[s[i-1:i]] {
			result += romanToArabic[s[i:i+1]] - 2*romanToArabic[s[i-1:i]]
		} else {
			result += romanToArabic[s[i:i+1]]
		}
	}
	if result == 0 {
		return 0, errors.New("некорректное римское число")
	}
	return result, nil
}

func intToRoman(num int) string {
	var result strings.Builder
	for _, conversion := range arabicToRoman {
		for num >= conversion.value {
			result.WriteString(conversion.symbol)
			num -= conversion.value
		}
	}
	return result.String()
}

func performOperation(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("некорректный оператор")
	}
}

