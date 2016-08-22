package main

import (
	"fmt"
)

//Panic y Recover
/*Panic es el error que detiene la ejecución de un programa*/
func imprimir() {
	fmt.Println("Hola Mauricio")
	//panic("Error")
	//Recover va a recuperar lo que sea que haya mandado panic como
	//error y lo va a almacenar en la cadena
	//cadena := recover()
	//Pero se usa con defer
	defer func() {

		cadena := recover()
		fmt.Println(cadena)

	}()
	panic("Error")
}

func main() {
	imprimir()
	fmt.Println("Hola main")
}
