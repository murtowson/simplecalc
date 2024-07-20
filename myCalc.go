package main

import( 
"fmt"
"strings"
)

func main() {
	var a, b int
	var operator string

	fmt.Println("Добро подаловать в калькулятор для пейджера!\n Т.к. мои знания ограничены, выберите оператора из списка:\n +, -, /, * ")

	fmt.Scanln(&operator)

	fmt.Println("Теперь введите числа с которыми мне предстоит произвести вычисления:")
	fmt.Scanln(&a)

	fmt.Scanln(&b)
	
	switch operator {
	case "+":
		fmt.Printf("Результат: %d\n", a+b)
	case "-":
		fmt.Printf("Результат: %d\n", a-b)
	case "/":
		fmt.Printf("Результат: %d\n", a/b)
	case "*":
		fmt.Printf("Результат: %d\n", a*b)	

	}

romanToArabic:=map[string]int{
	"I":1,
	"IV":4,
	"V":5,
	"IX":9,
	"X":10,
	"XL":40,
	"L":50,
	"XC":90,
	"C":100,
	"CD":400,
	"D":500,
	"CM":900,
	"M":1000,
}	
// создаю функцтю форматирования того что ввел юзер в читаемый формат: преобр в заглавные и удаление пробелов
func preProcRomNum(input string) string{

}	


}





