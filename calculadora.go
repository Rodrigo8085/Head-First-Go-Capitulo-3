package main

import (
	"fmt"
	"log"
)

func main() {
	var cantidad, total float64
	cantidad, err := pintar(4.3, 3.0)
	//validacion si es que hay un error
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%0.2f litros son los que se necesitan\n", cantidad)
		total += cantidad
		cantidad, err = pintar(5.3, 3.5)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("%0.2f litros son los que se necesitan\n", cantidad)
			total += cantidad
			fmt.Printf("Total: %0.2f litros\n", total)
		}
	}
}

//la funciones se declaran con el camel Case no pude inicar con un número y debe empezar con una ltra minuscula
//con el paso de parametros se tienen que declarar las variables con su tipo de dato

func pintar(width float64, height float64) (float64, error) {
	//Validación de que deber el valor es mas chico que el cero en ambas
	if width < 0 {
		//regresa el valor en cero el float y el error en una impersion de pantalla del error
		return 0, fmt.Errorf("El valor de ancho es %0.2f es negativo: ", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("El valor del largo es %0.2f es negativo: ", height)
	}
	area := width * height
	//nil indica que no existe un error durante las validaciónes
	return area / 10, nil
}
