package inventarioclases

import "fmt"

//Las clases publicas deben tener mayuscula inicial
type CarPublic struct {
	Brand string //Un atributo publico tiene mayuscula inicial
	year  int    //Un atributo privado tiene minuscula inicial
}

//Calse privada
type carPrivate struct {
	brand string
	year  int
}

//Las funciones publicas deben tener mayuscula inicial
func PrintMessage() {
	fmt.Println("Hola desde una funcion publica")
}
