package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	go mi_nombre_lento("Mauricio")
	fmt.Println("Que aburrido...")
	go func() {
		var wait string
		fmt.Scanln(&wait)
	}()

}

func mi_nombre_lento(name string) {
	/*Split permite partir las cadenas en otras cadenas
	Ejemplo:
	"hola_mundo quisieramos separarlo .Split("holamundo","_")
	lo separamos con el guión bajo
	"*/
	letras := strings.Split(name, "") //si le pasamos una cadena vacia
	//se entiende que va a separar cada letra del string
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)

	}
}
