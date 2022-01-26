package main

import (
	"fmt"
	"strings"
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

//Primeras funciones
func normalFunction(message string){
	fmt.Println(message)
}

func tripleArgument(a, b int, c string){
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func dobleReturn(a int) (b, c int) {
	return returnValue(a), a + 4
}

func introduccionFunciones(){
	normalFunction("Hola Mundo")
	tripleArgument(1, 2, "Hola")
	
	value := returnValue(2)
	fmt.Println("value", value)

	value1, value2 := dobleReturn(2)
	//Podemos descartar un de los resultados con el simbolo piso _
	//Ejemplo: value1, _ := dobleReturn(2)
	fmt.Println("value1, value2", value1, value2)
}

//Ciclos
func ciclos() {
	//For condicional
	for i :=0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("")
	//For While
	counter := 10
	for counter < 20 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("")
	//For forever
	counterForever := 30
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}

func operadorSwitch(){
	//Switch
	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//Switch Sin condicion
	numero := 135
	switch {
	case numero > 150:
		fmt.Println("Numero mayor a 150")
	case numero >= 100 && numero <= 150:
		fmt.Println("Numero entre 100 y 150")
	default:
		fmt.Println("Numero menor a 100")
	}
}


//Defer
func funcionSecundariaDefer(){
	fmt.Println("Primera metodo")
	defer fmt.Println("Ultima medoto")
	fmt.Println("Segunda Metodo")
}

func funcionDefer(){
	defer fmt.Println("Ultima principal")
	//Defer nos sirve para que una sentencia se ejecute al final de todo el codigo
	funcionSecundariaDefer()
	fmt.Println("Primera principal")
	fmt.Println("Segunda Principal")
}

//Continue y break
func continueBreak(){
	for i := 0; i < 10; i++{
		fmt.Println(i)

		//Continue
		if(i == 2) {
			fmt.Println("Es 2")
			continue //Se usa para que aun si sudece un error el flujo continue, se recomienda controlar el error igualmente
		}

		//Break
		if (i == 4){
			fmt.Println("Es 4")
			break //Termina el flujo, en este caso termina el for
		}
	}
}

//Construccion de arrys y Slice
	func arreglos(){
	//Array
	var array [4]int
	array[0] = 10
	array[1] = 15
	fmt.Println(array, len(array), cap(array))

	//Slice
	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice, len(slice), cap(slice))

	//Metodos slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[2:])

	//Append
	slice = append(slice, 16)
	fmt.Println(slice)

	//Append de lista
	newSlice := []int{17, 18 , 19}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}

func funcionRange(){
	slice :=[]string{"Hola", "que", "hace"}

	//Range
	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	for _, valor := range slice {
		fmt.Println(valor)
	}
	
	for i := range slice {
		fmt.Println(i)
	}
}

func isPalingromo(text string) {
	text = strings.ToLower(text)
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es un palindromo")
	} else {
		fmt.Println("No es un palindromo")
	}
}

func mapas() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	//Recorrer un mao
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar un valor
	value := m["Jose"]
	fmt.Println(value)

	value1, ok1 := m["Jose"]
	fmt.Println(value1, ok1)

	value2, ok2 := m["jose"]
	fmt.Println(value2, ok2)
}

func main() {
	

}
