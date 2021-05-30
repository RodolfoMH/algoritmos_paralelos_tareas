package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func print(str string) {
	fmt.Println(str)
}

var cola sync.WaitGroup

func generarRandon(posicion string, rango int) {
	fmt.Println(posicion, rand.Intn(rango))
	cola.Done()
}

//generar cada numero del loto de forma asincrona usando numeros randon
func main() {

	fmt.Println("Los 7 numeros ganadores son: ")

	cola.Add(7)

	go generarRandon("Primer numero => ", 38)
	go generarRandon("segundo numero => ", 38)
	go generarRandon("tercero numero => ", 38)
	go generarRandon("cuarto numero => ", 38)
	go generarRandon("quinto numero => ", 38)
	go generarRandon("sexto numero => ", 38)

	go generarRandon("loto mas => ", 10)

	cola.Wait()

	// var wait string
	// fmt.Scanln(&wait)

}
