package inventarioclases

import "fmt"

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
