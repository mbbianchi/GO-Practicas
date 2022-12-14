package main

import "fmt"

// funcion de sumar dos numeros
func sumar(a, b int) int {
	return a + b

}

//funcion principal

func main() {

	//variables
	var a, b, suma int

	//entrada de datos

	fmt.Print("INgrese N01:")
	fmt.Scanln(&a)

	fmt.Print("ingrese N02:")
	fmt.Scanln(&b)

	//llamar a la funcion sumar

	suma = sumar(a, b)

	//salida de datos

	fmt.Println("la suma es:", suma)

}
