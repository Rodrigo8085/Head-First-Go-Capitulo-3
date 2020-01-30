package main

import (
	"fmt"
)

func main() {
	pintar(4.2, 3.0)
	pintar(5.2, 3.5)
	pintar(5.0, 3.3)
}

//la funciones se declaran con el camel Case no pude inicar con un número y debe empezar con una ltra minuscula
//con el paso de parametros se tienen que declarar las variables con su tipo de dato

func pintar(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f litros son los que se necesitan\n", area/10)
}
