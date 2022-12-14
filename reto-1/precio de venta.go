package main

import "fmt"

//funcion para calcular igv

func calcularigv(vv, igv float64) float64 {
	return (igv * vv)

}

//funcion para calcular el precio de venta

func precio(vv, igv float64) float64 {
	return vv + igv

}

// funcion principal

func main() {

	//variables

	var vv, igv, pr float64

	//entrada de datos

	fmt.Println(" ingrese el valor de venta:")
	fmt.Scanln(&vv)

	// calculo de igv
	igv = calcularigv(vv, 0.18)

	// calcular precio de venta
	pr = precio(vv, igv)

	//salida de datos

	fmt.Println("valor de venta:", vv)
	fmt.Println("valor con igv:", igv)
	fmt.Println("precio total:", pr)

}
