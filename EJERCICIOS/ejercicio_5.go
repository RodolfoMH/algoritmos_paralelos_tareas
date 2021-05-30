package main

import (
	"fmt"
)

func print(str string) {
	fmt.Println(str)
}

//Metodo que realiza la sumattoria de todos los numeros dentro de un rango dado
func sumarRango(desde int, hasta int) int {

	var suma int

	for i := desde; i < hasta; i++ {
		fmt.Println("de ", desde, "a", hasta, " => ", suma, " + ", i, " = ", (suma + i))
		suma = suma + i
	}
	return suma
}

func main() {

	var desde int
	var hasta int

	var total int

	//Realiza una suma con todos los numeros desde el rango x hasta el y, de forma secuencial
	print("Ingrese el rango de la suma secuencial. ")
	print(" desde => ")
	fmt.Scanln(&desde)
	print(" hasta => ")
	fmt.Scanln(&hasta)

	total = sumarRango(desde, hasta)

	fmt.Println("total = ", total)
}
