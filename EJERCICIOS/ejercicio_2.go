package main

import (
	"fmt"
	"sync"
)

var cola sync.WaitGroup

func rutina(str string) {

	for i := 0; i < 3; i++ {
		fmt.Println(str, i)
	}
	cola.Done()
}

//Ejecuta 3 tareas de forma concurrente pero una vez que una de ellas inicia las otras se bloan hasta que esta termine imprmiendo un mensaje desde el punto de contro de cada rutina
func main() {

	cola.Add(3)

	go rutina("rutina 1: ")
	go rutina("rutina 2: ")
	go rutina("rutina 3: ")

	cola.Wait()
}
