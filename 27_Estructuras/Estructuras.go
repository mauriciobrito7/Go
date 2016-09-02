package main

import (
	"fmt"
)

//Estructuras

//Persona : estructura que representa a una persona
type Persona struct {
	Nombre string
	Edad   int
}

//Older determina cual es mayor y la diferencia de edad
func Older(p1, p2 Persona) (Persona, int) {
	if p1.Edad > p2.Edad {
		return p1, p1.Edad - p2.Edad
	}
	return p2, p2.Edad - p1.Edad
}

func main() {
	//Declarar una variable de tipo Persona #1
	var p Persona
	p.Nombre = "Mauricio"
	p.Edad = 21
	fmt.Println("Estructura p de tipo Persona: ", p)

	fmt.Println("Nombre p:", p.Nombre)
	fmt.Println("Edad p:", p.Edad)

	//Declarar una variable de tipo Persona #2
	p2 := Persona{Nombre: "Rosa", Edad: 60}
	fmt.Println("Nombre p2:", p2.Nombre)
	fmt.Println("Edad p2:", p2.Edad)

	//Declarar una variable de tipo Persona #3
	p3 := Persona{"Miguel", 24}
	fmt.Println("Nombre p3:", p3.Nombre)
	fmt.Println("Nombre p3:", p3.Edad)

	pOlder, pDiff := Older(p, p2)
	fmt.Println("Entre", p.Nombre, " y ", p2.Nombre, ",", pOlder.Nombre, " es mayor por",
		pDiff)

}
