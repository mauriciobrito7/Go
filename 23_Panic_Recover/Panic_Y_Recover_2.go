package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ejecucion := readFile()
	/*Esto nunca se va a ejecutar si el archivo no existe*/
	fmt.Println(ejecucion)
}

func readFile() bool {
	archivo, error := os.Open("Holaa.txt")
	defer func() {
		archivo.Close()
		fmt.Println("Defer")

		/*Una forma de detener un panic es usando recover*/
		//recover()
		//Asi se obtiene el panic y sigue ejecutandose el programa
		r := recover()
		fmt.Println(r)
	}()

	if error != nil {
		panic(error)
	}
	scanner := bufio.NewScanner(archivo)
	var i int
	for scanner.Scan() {
		i++
		linea := scanner.Text()
		fmt.Println(i)
		fmt.Println(linea)
	}
	fmt.Println("Nunca me ejecuto")
	return true
}
