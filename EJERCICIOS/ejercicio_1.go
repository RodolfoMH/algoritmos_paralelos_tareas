package main

import (
	"fmt"
)

func print(str string) {
	fmt.Println(str)
}

//Imprime un grupo de mensajes de forma asincrona
func main() {

	go print("Hola 1")
	go print("Hola 2")
	go print("Hola 3")
	go print("Hola 4")
	go print("Hola 5")
	go print("Hola 6")

	var wait string
	fmt.Scanln(&wait)
}
