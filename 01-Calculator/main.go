package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var operator string
	fmt.Println("Calculator")
	fmt.Println("Choose an operation:")
	fmt.Println(` + - Plus
 - - Minus
 * - Multiply
 / - Divide`)
	fmt.Scanln(&operator)
	if operator == "+"{
		result := num1 + num2
		fmt.Println(result)
}
}