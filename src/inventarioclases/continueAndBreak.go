package inventarioclases

import "fmt"

func continueBreak() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue //Se usa para que aun si sudece un error el flujo continue, se recomienda controlar el error igualmente
		}

		//Break
		if i == 4 {
			fmt.Println("Es 4")
			break //Termina el flujo, en este caso termina el for
		}
	}
}
