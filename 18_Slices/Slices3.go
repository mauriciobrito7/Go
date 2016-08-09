package main

import "fmt"

func main() {
	//Declaramos el Slice x
	x := make([]byte, 4, 10)
	//Imprimimos x
	fmt.Println(x)
	//Inicializamos x
	//Cuando se pasa con comillas simples se pasa es el valor en bytes de cada letra
	x = []byte{'H', 'O', 'L', 'A'}
	fmt.Println(x)
	//Imprimimos x dandole formato UTF-8
	fmt.Printf("Slice x: %s\n", x)
	//Imprimimos cada elemento del slice x por separado
	//utilizando un bucle for
	for i := 0; i < len(x); i++ {
		fmt.Printf("%c\n", x[i])
	}

	//Asignamos un espacio en blanco en el indice 5 del slice x
	//A pesar que se tiene una capacidad de 10 no se puede acceder más alla
	//de la longitud declarada para es existe la función append
	//x[5] = ' ' esto es un  error
	//Se le pasa el slice a modificar y lo que se le va agregar
	//La función append devuelve el mismo slice si tiene la suficiente longitud para agregar
	//O devuelve un nuevo slice si no lo tiene por eso es que se escribe asi
	x = append(x, ' ')
	fmt.Println(x)
	//Aqui podemos notar que la capicidad a cambiado de 10 a 8 porque es un nuev slice
	//Dado que no tenia sufciente longitud genera una capacidad del doble de la longitud anterior
	fmt.Println(cap(x))

	//Agregamos los byte "MUNDO" al slice x
	x = append(x, 'M', 'U', 'N', 'D', 'O')
	fmt.Println(x)
	fmt.Printf("%s\n", x)
	//Capacidad x
	fmt.Println(cap(x))
	//Longitud x
	fmt.Println(len(x))

	//Declaramos el slice y
	var y []int
	for i := 1; i <= 15; i++ {
		//Agregamos el valor de i al slice y
		y = append(y, i)
		//Imprime el slice y
		fmt.Println("Slice y:", y)
		//Imprimimos la longitud y la capacidad de y
		fmt.Printf("Len y: %d -Cap y: %d - Elementos:%d \n", len(y), cap(y), i)
	}

}
