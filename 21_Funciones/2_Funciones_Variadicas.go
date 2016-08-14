package main

import "fmt"

//Variadic Functions
//Recibe una cantidad indeterminada de enteros
//Al colocar ... decimos que vamos a recibir valores enteros pero no sabemos que cantidad
//Convierte la variable número en un slice y se recorre con range
func sumar(numeros ...int) int {
	resultado := 0
	for _, numero := range numeros {
		resultado += numero
	}
	return resultado
}

//Solo el ultimo parametro puede recibir variables indefinidas ninguna otra del mismo tipo
func print(cadena string, cadenas ...string) {
	for _, c := range cadenas {
		cadena += " "
		cadena += c
	}

	fmt.Println(cadena)
}

func main() {
	fmt.Println(sumar(1, 2, 3, 4))
	//Sino se le pasa parametros devuelve el valor por defecto
	fmt.Println(sumar())
	//Sin embargo a crear un slice y pasarselo por parametros a la funcion suma
	//nos da un error dado que la función no recibe un slice de enteros
	//Si no que recibe enteros en si
	print("Hola", "Mundo", "que", "es", "la", "que", "hay", "?")
	numeros := []int{
		25,
		56,
		32,
	}

	//fmt.Println(sumar(numeros)) ERROR
	//Para indicarle a Go que es un slice y que no lo tome como parametro
	//Si no que el contenido de ese slice es que lo estamos pasando como parametros
	fmt.Println(sumar(numeros...))

}
