package main

import "fmt"

func main() {
	a := 20
	b := 10

	// suma
	result := a + b

	fmt.Println("suma", result)

	// resta
	result = a - b
	fmt.Println("resta", result)

	// multiplicacion
	result = a * b
	fmt.Println("multi", result)

	// division
	result = a / b
	fmt.Println("division", result)

	var div float64 = 3 / 2
	fmt.Println("division", div)

	// modulo
	result = a % b
	fmt.Println("modulo", result)

}
