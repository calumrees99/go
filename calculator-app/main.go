package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Input Calculation:")

	var number1 int
	var operator string
	var number2 int

	if _, err := fmt.Scan(&number1, &operator, &number2); err != nil {
		log.Fatalf("Invalid input: %v", err)
	}

	switch operator {
	case "+":
		add(number1, number2)
	case "-":
		subtract(number1, number2)
	case "*":
		multiply(number1, number2)
	case "/":
		divide(number1, number2)
	default:
		fmt.Println("Invalid operator!")
	}
}

func add(number1 int, number2 int) {
	result := number1 + number2
	fmt.Printf("Result: %d\n", result)
}

func subtract(number1 int, number2 int) {
	result := number1 - number2
	fmt.Printf("Result: %d\n", result)
}

func multiply(number1 int, number2 int) {
	result := number1 * number2
	fmt.Printf("Result: %d\n", result)
}

func divide(number1 int, number2 int) {
	result := number1 / number2
	fmt.Printf("Result: %d\n", result)
}
