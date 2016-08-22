package main

import (
	"fmt"
	"io/ioutil" //Nos permite leer el archivo
)

func main() {
	/*A readFile le pasamos la ubicación y el nombre del archivo
	./ quiere decir que es en la misma carpeta */
	/*readFile retorna dos cosas una el archivo y otra el error
	por si acaso el archivo no lo haya encontrado*/
	file_data, err := ioutil.ReadFile("./hola.txt")

	if err != nil {
		panic("Error no se encuentra el archivo")
	}

	fmt.Println(string(file_data))

}
