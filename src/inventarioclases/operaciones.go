package inventarioclases

import "fmt"

func operaciones() {
	const pi float64 = 3.1415
	fmt.Println("")
	//Calcular le area de un cuadrado
	fmt.Println("Calcular le area de un cuadrado")
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado: ", areaCuadrado)

	fmt.Println("")
	fmt.Println("Operaciones Aritmeticas")
	x := 10
	y := 51

	//Suma
	result := x + y
	fmt.Println("Suma: ", result)

	//Resta
	result = x - y
	fmt.Println("Resta: ", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion", result)

	//Division
	result = y / x
	fmt.Println("Division", result)

	//Modulo (Residuo)
	result = y % x
	fmt.Println("Modulo", result)

	//Incremental
	x++
	fmt.Println("Incremental", x)

	//Decremental
	x--
	fmt.Println("Decremental", x)

	fmt.Println("")
	fmt.Println("Retos")
	//Retos

	//Area de un rectangulo
	base := 12
	altura := 14
	result = base * altura
	fmt.Println("Area del rectangulo: ", result)

	//Area de un trapecio
	var baseMenor float64 = 10
	var baseMayor float64 = 15
	var alturaTrapecio float64 = 4
	var areaTrapecio float64 = ((baseMayor + baseMenor) / 2.0) * alturaTrapecio
	fmt.Println("Area del trapecio: ", areaTrapecio)

	//Area de un circulo
	var radioCiculo float64 = 7
	var areaCirculo float64 = pi * pi * radioCiculo
	fmt.Println("Area circulo:", areaCirculo)

}
