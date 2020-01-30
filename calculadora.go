package main

import (
	"fmt"
)

var metrosporLitro float64

func main() {
	metrosporLitro = 10
	fmt.Printf("%.2f litros son los que se necesitan\n", pintar(4.2, 3.0))
}

//la funciones se declaran con el camel Case no pude inicar con un número y debe empezar con una ltra minuscula
//con el paso de parametros se tienen que declarar las variables con su tipo de dato

func pintar(width float64, height float64) float64 {
	area := width * height
	return area / metrosporLitro
}
