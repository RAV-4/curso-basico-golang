package inventariocases

import "fmt"

func arreglos(){
	//Array
	var array [4]int
	array[0] = 10
	array[1] = 15
	fmt.Println(array, len(array), cap(array))

	//Slice
	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice, len(slice), cap(slice))

	//Metodos slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[2:])

	//Append
	slice = append(slice, 16)
	fmt.Println(slice)

	//Append de lista
	newSlice := []int{17, 18 , 19}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}