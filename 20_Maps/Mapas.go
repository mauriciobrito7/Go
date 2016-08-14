/*Los maps en go son los que en otros lenguajes se conoce como diccionarios o tablas hash
Los maps nos permiten indicar un indice y un valor y mediante el índice podemos obtener, modificar o borrar el valor
¿Cómo se crea un map en Go?*/

/*Utilizamos la función make nativa de Go que declara un vector dinamico*/
package main

import (
	//Librería para impresión de Go
	"fmt"
)

func main() {
	//Declarar un map
	//Le pasamos entre [] el tipo del indice y luego le pasamos el tipo de valor
	//También se le puede pasar otro parametro indicando la capacidad
	x := make(map[string]string)
	fmt.Println(x)
	//Agregar valores
	x["nombre"] = "Mauricio"
	x["edad"] = "29"
	fmt.Println(x)

	//Acceder a los valores
	fmt.Println(x["nombre"])
	fmt.Println(x["edad"])

	//Otra forma de declarar Maps
	edades := map[string]int{
		"Karla":    28,
		"Mauricio": 20,
		"Manuel":   21,
		"Maria":    19,
	}
	fmt.Println(edades["Karla"])
	//Para el índice podemos utilizar cualquier tipo que sea comparable con el operador ==
	//Para el valor podemos utilizar cualquier tipo
	dias := map[int]string{
		1: "Lunes",
		2: "Martes",
		3: "Miercoles",
	}

	fmt.Println(dias[2])

	//Eliminar elementos de un Map
	delete(edades, "Maria")
	fmt.Println(edades)

	//Si el elemento no existe los mapas devuelven
	//el valor 0 del tipo de dato sea el elemento
	fmt.Println("La edad de pedro es:", edades["pedro"])
	fmt.Println(edades)

	//Podemos utilizar el operador ++
	//Aquí el elemento pedro no existe y le estamos aplicando el valor de
	//incremento por lo cual como no existe pedro  en el map va a devolver el valor 0
	// y luego se incrementa en 1 entonces ya pedro va existir
	edades["pedro"]++
	fmt.Println(edades)
	//Podemos utilizar los operadores de tipo +=, -=, *=, etc
	edades["pedro"]++
	fmt.Println(edades)

	//Igual aqui con Carlos no existe pero al aumentarlo en 2 ya existiria
	edades["Carlos"] += 2
	fmt.Println(edades)
	//Saber si un elemento de un Map existe
	//Si no existe se devuelve 0
	//El map devuelve otro valor booleano si ese map existe o no y ese valor se guarda en ok
	if edad, ok := edades["Manuel"]; ok {
		if edad < 18 {
			fmt.Println("Es menor de edad")
		} else {
			fmt.Println("Es mayor de edad")
		}
	} else {
		fmt.Println("El valor no existe")
	}

	//Saber el tamaño de un Map
	fmt.Println(len(edades))

	//Los elementos de un Map no son variables
	//por lo cual no podemos obtener sus direcciones
	// puntero := &edades["Carlos"] ERROR

	//Recorrer un map
	//Los Maps son desordenados
	//range en Go es el for each en otros lenguajes
	//Los maps son dinamicos no siempre vamos a obtener los valores en mismo orden
	for nombre, edad := range edades {
		fmt.Printf("La edad de %s es: %d \n", nombre, edad)
	}
}
