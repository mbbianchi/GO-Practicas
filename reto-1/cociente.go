package main

import "fmt"

//funcion cociente de dos numeros

func cociente(a, b int) int {

	return (a / b)
}

//funcion que calcula el resto

func resto(a, b int) int {
	return a % b

}

// funcion principal
func main() {

	var (
		a, b, coc, res int
	)

	// ingresar valores

	fmt.Println("Ingrese un valor")
	fmt.Scanln(&a)

	fmt.Println("Ingrese otro valor")
	fmt.Scanln(&b)

	//llamar a la funcion cociente

	coc = cociente(a, b)

	// llamar a la funcion resto

	res = resto(a, b)

	//salida de datos

	fmt.Println("el resultado es:", coc)
	fmt.Println("el resto es:", res)
}
