package main

import (
	"fmt"
)

func print(cadena string) {
	fmt.Println(cadena)
}

func print2(cadena, cadena2 string) {
	fmt.Println(cadena, cadena2)
}

/*print3 esta recibiendo la función print la cual recibe un string
el cual se esta pasando por fprint*/
func print3(fprint func(string)) {
	fprint("Hola Mundo desde Print3")
}

/*Función incrementar que retorna una función que no recibe parametro y devuelve un int*/

func incrementar() func() int {
	i := 0
	//Se retona inmediatamente una función que retorna un int.
	return func() (r int) {
		r = i
		i += 2
		return
	}
}

/*
En Go no se puede declarar una función dentro de otra de manera explicita
*/
func main() {

	cadena := "Hola Mundo"
	/*Una manera de tener una función dentro de otra es guardandola en una variable
	Aquí ya la variable imprimir va a contener a la función print y va a ser de ese tipo*/
	imprimir := print /*Obtiene la función guardandola sin los parentesis*/
	imprimir(cadena)

	//Otra forma de tener una función de otra función
	/*La función imprimir tiene acceso a las variables dentro de main*/
	imprimir2 := func() {
		fmt.Println(cadena)
	}
	imprimir2()
	//Esto ocaciona un error porque ya imprimmir tiene guardada una firma de la función anterior
	//que es distinta a print3
	//imprimir = print2 ERROR

	/*Le paso la primera función que es para imprimir una cadena*/
	print3(print)

	//Las funciones son comparables con nil
	//Las funciones son comparables con el valor nil
	var fb func()

	if fb == nil {
		fmt.Println("fb es igual a nil")
	}
	/*El valor que utiliza la función no se resetea porque queda almacenada
	en la variable y como es un closure cada vez que se llama el valor que utiliza
	la función interna de la funcion hace que se guarde */
	inc := incrementar()

	fmt.Println("Valor de i:", inc())
	fmt.Println("Valor de i:", inc())
	fmt.Println("Valor de i:", inc())
	fmt.Println("Valor de i:", inc())
}
