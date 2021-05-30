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

func main() {

	cola.Add(3)

	go rutina("rutina 1: ")
	go rutina("rutina 2: ")
	go rutina("rutina 3: ")

	cola.Wait()
}
