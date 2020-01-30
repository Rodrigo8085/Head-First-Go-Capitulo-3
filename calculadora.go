package main

import (
	"fmt"
)

var metrosporLitro float64

func main() {
	var cantidad, total float64
	cantidad = pintar(4.3, 3.0)
	fmt.Printf("%0.2f litros son los que se necesitan\n", cantidad)
	total += cantidad
	cantidad = pintar(5.3, 3.5)
	fmt.Printf("%0.2f litros son los que se necesitan\n", cantidad)
	total += cantidad
	fmt.Printf("Total: %0.2f litros\n", total)
}

//la funciones se declaran con el camel Case no pude inicar con un número y debe empezar con una ltra minuscula
//con el paso de parametros se tienen que declarar las variables con su tipo de dato

func pintar(width float64, height float64) float64 {
	area := width * height
	return area / 10
}
