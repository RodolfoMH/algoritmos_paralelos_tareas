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
