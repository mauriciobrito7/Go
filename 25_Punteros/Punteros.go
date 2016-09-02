package main

import "fmt"

func incrementar(numero *int) {
	*numero++
	fmt.Println("El número de la función incrementar:", *numero)
}

func main() {
	//a es de tipo int
	a := 25
	fmt.Println("Valor de a:", a)
	fmt.Println("Dirección de a:", &a)
	//b es un puntero de a, por lo que es de tipo *int
	b := &a

	//Se le imprime el contenido de b
	fmt.Println(b)

	//b es del tipo *int no int
	//b=25 error
	//le asignamos un nuevo valr a (a) a traves de b
	*b = 50
	fmt.Println("valor nuevo de a:", a)
	a++
	fmt.Println("Valor que apunta b:", *b)

	//Valor por defecto de los punteros es nil
	if b != nil {
		fmt.Println("b es diferente de nil")
	}
	//Los punteros son comparables
	if &a == b {
		fmt.Println("Es la misma memoria")
	}

	//Utilizar new()
	//nos devuelve el puntero de tipo que le pasamos por parametro
	d := new(int) //*int
	fmt.Println("Direccion de d", d)
	fmt.Println("Valor de d", *d)

	d = b
	fmt.Println("Dirección nueva de d:", d)
	fmt.Println("Nuevo valor de d:", *d)

	//Usar punteros en funciones
	numero := 2
	fmt.Println("Número antes de incrementar:", numero)
	incrementar(&numero)
	fmt.Println("Número despues de incrementar:", numero)

}
