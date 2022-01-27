package inventariocases

import "fmt"

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