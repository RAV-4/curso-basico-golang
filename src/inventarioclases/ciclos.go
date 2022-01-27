package inventarioclases

import (
	"fmt"
	"strings"
)

func ciclos() {
	//For condicional
	for i := 0; i < 10; i++ {
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
