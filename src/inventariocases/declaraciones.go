package inventariocases

import "fmt"

func declaraciones() {
	//Declaracion de constantes
	fmt.Println("Declaracion de constantes")
	const pi float64 = 3.1415
	const pi2 = 3.14

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	fmt.Println("")
	//Declaracion de variables
	fmt.Println("Declaracion de variables")
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	fmt.Println("")
	//Zero value
	fmt.Println("Zero value")
	var a int     //Valor por defecto 0
	var b float64 //Valor por defecto 0
	var c string  //Valor por defecto vacio
	var d bool    //Valor por defecto false

	fmt.Println(a, b, c, d)
}