package inventariocases

import "fmt"

type car struct {
	brand string
	year int
}

func creacionStructs(){
	//Creando una instancia con parametros
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	//Creado una instancia vacia
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar) 
}