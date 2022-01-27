package inventarioclases

import "fmt"

func funcionSecundariaDefer() {
	fmt.Println("Primera metodo")
	defer fmt.Println("Ultima medoto")
	fmt.Println("Segunda Metodo")
}

func funcionDefer() {
	defer fmt.Println("Ultima principal")
	//Defer nos sirve para que una sentencia se ejecute al final de todo el codigo
	funcionSecundariaDefer()
	fmt.Println("Primera principal")
	fmt.Println("Segunda Principal")
}
