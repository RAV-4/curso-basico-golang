package inventariocases

import "fmt"

type figuras2D interface{
	area() float64
}
type cuadrado struct {
	lado float64
}

type rectangulo struct {
	base float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.lado * c.lado
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area: ", f.area())
}

func flujoTrabajo() {
	myCuadrado := cuadrado{lado: 3}
	myRectangulo := rectangulo{base: 4, altura: 2}
	
	calcular(myCuadrado)
	calcular(myRectangulo)

	//Lista de interfaces
	myInterface := []interface{}{"hola", 12, 5.9}
	fmt.Println(myInterface...)
}