package main

import (
	"fmt"
)

func main() {
	//Slices
	//Declarar Slices
	//También se pueden incializar los Slices pero no se indica la capacidad
	var j []int
	fmt.Println(j)

	//Declarar Slices inicializados
	x := []int{1, 2, 3}
	fmt.Println(x)
	//Declarar Slices usando make, indicando la Longitud
	//make viene integrada
	//Se le pasa el tipo y la longitud y un tercer parametro para la capacidad
	y := make([]int, 5, 10)
	fmt.Println(y)
	fmt.Println("Longitud :", len(y))
	fmt.Println("Capacidad :", cap(y))
	//Definimos un array con 10 elementos de tipo int
	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Array ar:", ar)
	var a, b []int
	fmt.Println("Slice a:", a)
	fmt.Println("Slice b:", b)
	// 'a' apunta del 3ro al 5to elementos en el array ar.
	a = ar[2:5]
	//ahora 'a' tiene los elementos ar[2], ar[3] y ar[4]
	fmt.Println("Slice a:", a)

	//'b' es otro slice del array ar
	b = ar[3:5]
	//Ahora 'b' tiene ar[3] y ar[4]
	fmt.Println("Slice b:", b)
	b[0] = 25
	fmt.Println("Asignemos b[0]=25")
	fmt.Println("Slice b:", b)
	fmt.Println("Slice a:", a)
	fmt.Println("Array ar:", ar)
	fmt.Println("Cap(a):", cap(a))
	fmt.Println("Cap(b):", cap(b))

}
