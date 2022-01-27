package inventarioclases

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (pc pc) ping() {
	fmt.Println(pc.brand, "Pong")
}

func (pc *pc) duplicateRam() {
	pc.ram = pc.ram * 2
}

//Este metodo nos sirve para sobrecargar el toString
func (pc pc) String() string {
	return fmt.Sprintf("Tengo %d GB Ram, %d GB Disco y es una %s", pc.ram, pc.disk, pc.brand)
}

func usoPunteros() {
	a := 50
	b := &a //Se le asigna a B la direccion en memoria de a, osea que b es el puntero de a
	fmt.Println(a)
	fmt.Println(b)  //Asi obtenemos la direccion de memoria
	fmt.Println(*b) //Asi obtenemos el valor en esa direccion de memoria

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 250, brand: "MSI"}
	myPC.ping()

	fmt.Println(myPC)

	myPC.duplicateRam()
	fmt.Println(myPC)

	myPC.duplicateRam()
	fmt.Println(myPC)
}
