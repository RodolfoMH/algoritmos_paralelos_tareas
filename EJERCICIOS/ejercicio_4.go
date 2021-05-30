package main

import (
	"fmt"
	//"sync"
)

func print(str string) {
	fmt.Println(str)
}

//var cola sync.WaitGroup

func sumarRango(desde int, hasta int) int {

	var suma int

	for i := desde; i < hasta; i++ {
		fmt.Println("de ", desde, "a", hasta, " => ", suma, " + ", i, " = ", (suma + i))
		suma = suma + i
	}
	return suma
}

//Realiza la suma de un rango de numeros desde x hasta y, para esto divide el rango en 4 partes
//las cuales se suman de forma separadas en diferentes rutinas para luego unir los resultados agilizando asi por 4 la velocidad
// en la realizacion de dicha suma.
func main() {

	var desde int
	var hasta int

	var total int

	print("Ingrese el rango de la suma concurrente. ")
	print(" desde => ")
	fmt.Scanln(&desde)
	print(" hasta => ")
	fmt.Scanln(&hasta)

	offset := (hasta - desde) / 4

	go func() {
		total = total + sumarRango(desde, desde+offset)
	}()

	go func() {
		total = total + sumarRango(desde+offset, desde+offset+offset)
	}()

	go func() {
		total = total + sumarRango(desde+offset+offset, desde+offset+offset+offset)
	}()

	go func() {
		total = total + sumarRango(desde+offset+offset+offset, hasta)
	}()

	var wait string
	fmt.Scanln(&wait)

	fmt.Println("total = ", total)
}
