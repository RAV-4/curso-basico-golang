package inventariocases

import (
	"fmt"
	"time"
)

func say(text string) {
	fmt.Println(text)
}

func flujo() {
	say("Hello")
	
	go say("World")

	go func (text string )  {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second) //Funcion no recomendada para esperar que ejecute la funcion concurrente
}