package main

import (
	"fmt"
)

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

func operaciones() {
	const pi float64 = 3.1415
	fmt.Println("")
	//Calcular le area de un cuadrado
	fmt.Println("Calcular le area de un cuadrado")
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado: ", areaCuadrado)

	fmt.Println("")
	fmt.Println("Operaciones Aritmeticas")
	x := 10
	y := 51

	//Suma
	result := x + y
	fmt.Println("Suma: ", result)

	//Resta
	result = x - y
	fmt.Println("Resta: ", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion", result)

	//Division
	result = y / x
	fmt.Println("Division", result)

	//Modulo (Residuo)
	result = y % x
	fmt.Println("Modulo", result)

	//Incremental
	x++
	fmt.Println("Incremental", x)

	//Decremental
	x--
	fmt.Println("Decremental", x)

	fmt.Println("")
	fmt.Println("Retos")
	//Retos

	//Area de un rectangulo
	base := 12
	altura := 14
	result = base * altura
	fmt.Println("Area del rectangulo: ", result)

	//Area de un trapecio
	var baseMenor float64 = 10
	var baseMayor float64 = 15
	var alturaTrapecio float64 = 4
	var areaTrapecio float64 = ((baseMayor + baseMenor) / 2.0) * alturaTrapecio
	fmt.Println("Area del trapecio: ", areaTrapecio)

	//Area de un circulo
	var radioCiculo float64 = 7
	var areaCirculo float64 = pi * pi * radioCiculo
	fmt.Println("Area circulo:", areaCirculo)

}

func paqueteFmt() {
	fmt.Println("")
	//Paquete fmt
	fmt.Println("Paquete fmt")
	//Declaracion de variables
	helloMenssage := "Hello"
	worldMenssage := "world"

	//Println
	fmt.Println(helloMenssage, worldMenssage)
	fmt.Println("")

	//Printf
	nombre := "Platzy"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)
	fmt.Println("")

	message := fmt.Sprintf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Print(message)

	fmt.Println("")
	fmt.Printf("Tipo de dato helloMessage: %T\n", helloMenssage)
	fmt.Printf("Tipo de dato curso: %T\n", cursos)
}

func main() {

}
