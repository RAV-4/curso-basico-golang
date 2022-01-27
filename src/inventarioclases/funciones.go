package inventarioclases

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func dobleReturn(a int) (b, c int) {
	return returnValue(a), a + 4
}

func introduccionFunciones() {
	normalFunction("Hola Mundo")
	tripleArgument(1, 2, "Hola")

	value := returnValue(2)
	fmt.Println("value", value)

	value1, value2 := dobleReturn(2)
	//Podemos descartar un de los resultados con el simbolo piso _
	//Ejemplo: value1, _ := dobleReturn(2)
	fmt.Println("value1, value2", value1, value2)
}
