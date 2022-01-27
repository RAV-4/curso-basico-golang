package inventarioclases

import "fmt"

//chan<- canal de entrada de datos
//<-chan canal de salida de datos
func sayT(text string, c chan<- string) {
	c <- text
}

func usoChannels() {
	c := make(chan string, 1)

	fmt.Println("Hello")

	go sayT("Bye", c)

	fmt.Println(<-c)
}
