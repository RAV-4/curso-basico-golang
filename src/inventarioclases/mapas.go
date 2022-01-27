package inventarioclases

import "fmt"

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
