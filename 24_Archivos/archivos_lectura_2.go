package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	archivo, err := os.Open("./hola.txt")
	defer func() {
		archivo.Close()
	}()
	if err != nil {
		panic("Error")
	}

	scanner := bufio.NewScanner(archivo)
	/*Para iterar cada una de las lineas del archivo*/
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(linea)
	}
}
