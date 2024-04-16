package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsValidOperator(operator string) bool {
	switch operator {
	case
		"+",
		"-",
		"*",
		"/":
		return true
	}
	return false
}

func IsValidArabicNumber(number string) bool {
	for index, character := range number {
		if index == 0 && character == '0' {
			return false
		}

		switch character {
		case
			'0',
			'1',
			'2',
			'3',
			'4',
			'5',
			'6',
			'7',
			'8',
			'9':
			continue
		}
		return false
	}
	return true
}

func IsValidINumber(number string) bool {
	for _, character := range number {
		if character == 'I' || character == 'X' || character == 'V' {
			continue
		}
		return false
	}
	return true
}

func ConvertItoArabic(Inumber string) int {
	switch Inumber {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	}
	panic("Wrong Roman number")
}

func ConvertArabicToI(ArabicNumber int) string {
	if ArabicNumber < 1 {
		panic("Result is less than 1")
	}
	delimiters := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	numerals := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result string
	for i, val := range delimiters {
		for ArabicNumber >= val {
			result += numerals[i]
			ArabicNumber -= val
		}
	}
	return result

}

func complete_int_operation(operator string, num1 int, num2 int) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "/":
		return int(num1 / num2)
	case "*":
		return num1 * num2
	default:
		panic("Unknown operation!")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Calculator")

	for {
		fmt.Print("Enter an operation -> ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-2]
		var operands_array = strings.Split(text, " ")

		if len(operands_array) != 3 {
			panic("Wrong operands amount")
		}

		if !IsValidOperator(operands_array[1]) {
			panic("Wrong mathematical operator")
		}

		if !(IsValidArabicNumber(operands_array[0]) || IsValidINumber(operands_array[0])) {
			panic("Wrong type of the first number")
		}

		if !(IsValidArabicNumber(operands_array[2]) || IsValidINumber(operands_array[2])) {
			panic("Wrong type of the second number")
		}

		if IsValidArabicNumber(operands_array[0]) && IsValidArabicNumber(operands_array[2]) {
			num1, _ := strconv.Atoi(operands_array[0])
			num2, _ := strconv.Atoi(operands_array[2])
			fmt.Println("Result: ", complete_int_operation(operands_array[1], num1, num2))
		} else {
			if IsValidINumber(operands_array[0]) && IsValidINumber(operands_array[2]) {
				num1 := ConvertItoArabic(operands_array[0])
				num2 := ConvertItoArabic(operands_array[2])
				fmt.Println("Result: ", ConvertArabicToI(complete_int_operation(operands_array[1], num1, num2)))
			} else {
				panic("Both numbers must be the same type")
			}
		}
	}
}
