package main

import (
	"fmt"
)

//Funcion que imprime un texto pasado por parametro
func print(str string) { fmt.Println(str) }

//Metodo que realiza la sumattoria de todos los numeros dentro de un rango y retorna el total
func sumarRango(desde int, hasta int) int {

	var suma int
	for i := desde; i < hasta; i++ {
		fmt.Println("de ", desde, "a", hasta, " => ", suma, " + ", i, " = ", (suma + i))
		suma = suma + i
	}
	return suma
}

//Metodo principal
func main() {

	//Realiza la suma de un rango de numeros desde x hasta y, para esto divide el rango
	//en 4 partes las cuales se suman de forma separadas en diferentes rutinas para luego
	//unir los resultados agilizando asi por 4 la velocidad en la realizacion de dicha suma.

	var desde int
	var hasta int
	var total int

	//Le pido al usuario los rangos desde y hasta
	print("Ingrese el rango de la suma concurrente. ")
	print(" desde => ")
	fmt.Scanln(&desde)
	print(" hasta => ")
	fmt.Scanln(&hasta)

	//Divido el total de numeros dentro del rango en 4 partes para sumarlas por separado
	offset := (hasta - desde) / 4

	//Realizo la sumatoria en diferentes rangos, por ejemplo si el rango es de 1 a 100,
	//realizo la suma en 4 rangos de a 25 numeros cada uno en diferentes rutinas para
	//acelerar el tiempo x 4.
	go func() { total = total + sumarRango(desde, desde+offset) }()
	go func() { total = total + sumarRango(desde+offset, desde+(offset*2)) }()
	go func() { total = total + sumarRango(desde+(offset*2), desde+(offset*3)) }()
	go func() { total = total + sumarRango(desde+(offset*3), hasta) }()

	//Hago un scan para bloaquear el programa y este no finalize hasta que el usuario
	//inserte algun caracter
	var wait string
	fmt.Scanln(&wait)

	//Imprimo el total de la sumatoria
	fmt.Println("total = ", total)
}
