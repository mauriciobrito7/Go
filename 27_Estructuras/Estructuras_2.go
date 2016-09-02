package main

import (
	"fmt"
)

//Persona contiene los campos nombre y apellido
type Persona struct {
	Nombre   string
	Apellido string
}

//Estudiante contiene el campo persona y carrera
type Estudiante struct {
	Persona
	Carrera string
}

//Profesor tiene los campos Estudiante y Carrera
type Profesor struct {
	Estudiante
	Carrera string
}

func main() {
	mauricio := Estudiante{
		Persona{
			Nombre:   "Mauricio",
			Apellido: "Brito",
		},
		"Informatica",
	}

	fmt.Println(mauricio.Persona.Nombre, mauricio.Persona.Apellido, mauricio.Carrera)
	pedro := Profesor{
		Estudiante{
			Persona{
				"Pedro",
				"Almonte",
			},
			"Contabilidad",
		},
		"Informatica",
	}

	fmt.Println(pedro.Nombre, pedro.Apellido, pedro.Carrera, pedro.Estudiante.Carrera)
	manuel := Profesor{Estudiante{Persona{"Manuel", "Peralta"}, "Ing.Industrial"}, "Informatica"}
	fmt.Println("Manuel:", manuel)

	var jose Profesor
	jose.Nombre = "José"
	jose.Apellido = "Contreras"
	jose.Estudiante.Carrera = "Eduacación"
	jose.Carrera = "Mercadeo"
	fmt.Println(jose)
}
