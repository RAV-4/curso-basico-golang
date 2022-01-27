package inventariocases

import "fmt"

func menssage(text string, c chan<- string) {
	c <- text
}
func recorrerCanales() {

	c := make(chan string, 2)
	c <- "Mensaje 1"
	c <- "Mensaje 2"

	fmt.Println(len(c), cap(c))

	//Range y close
	close(c) //Cierra el canal para que no se le pueda insertar mas datos

	for menssage := range c {
		fmt.Println(menssage)
	}

	//Select
	email1 := make(chan string)
	email2 := make(chan string)

	go menssage("Mensaje por el email 1", email1)
	go menssage("Mensaje por el email 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <- email1:
			fmt.Println("Email :", m1)
		case m2 := <- email2:
			fmt.Println("Email :", m2)
		}
	}
}