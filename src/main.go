package main

import (
	"fmt"
)

func main() {

	//Declaracion de constantes
	const pi float64 = 3.1415
	const pi2 = 3.14

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	//Declaracion de variables
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero value
	var a int //Valor por defecto 0
	var b float64 //Valor por defecto 0
	var c string //Valor por defecto vacio
	var d bool //Valor por defecto false

	fmt.Println(a, b, c, d)

	//Calcular le area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado: ", areaCuadrado)
}
