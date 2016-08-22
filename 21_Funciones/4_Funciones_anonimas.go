package main

import (
	"fmt"
	"strings"
)

/*Asi seria la función sin ser anonima*/
func mapear(r rune) rune {
	return r + 1
}

func main() {
	cadena := "123456789"
	//La función del paquete strings lo que hace es que a una cadena de
	//caracteres a cada caracter que la conforma le aplica una función esa
	//función tiene que tener un tipo y recibe una rune y devuelve una rune
	//El código UTF-8 de cada uno de los caracteres en Go es de tipo rune
	//Función de manera anónima
	cadena = strings.Map(func(r rune) rune {
		return r + 1
	}, cadena)

	fmt.Println(cadena)
}
