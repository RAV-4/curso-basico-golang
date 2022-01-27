package inventariocases

import "fmt"

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