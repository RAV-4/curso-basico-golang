package inventarioclases

import (
	"fmt"
	"sync"
)

func sayText(text string, wg *sync.WaitGroup) {
	defer wg.Done() //con este defer podemos controlar cuando termina la ejecucion
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup //Se crea esta variable de este tipo que nos ayuda a tener control sobre las goroutine
	fmt.Println("Hello")
	wg.Add(1) // Indicamos que sigue una ejecucion de una goroutine

	go sayText("World", &wg)

	wg.Wait() // Esperamos que termine a ejecucion, sin el deficit de un sleep
}
