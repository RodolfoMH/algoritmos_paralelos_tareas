package main

import (
	"fmt"
	"sync"
)

var cola sync.WaitGroup

func rutina(str string) {
	for i := 0; i < 3; i++ {
		go fmt.Println(str, i)
	}
}

//Imprime un secuencias de mensajes de contros desde diferentes rutinas de manera simultanea
func main() {

	go rutina("rutina 1: ")
	go rutina("rutina 2: ")
	go rutina("rutina 3: ")

	var wait string
	fmt.Scanln(&wait)
}
