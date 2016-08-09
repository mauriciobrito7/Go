package main

import "fmt"

func main() {
	nombres := []string{
		"Mauricio",
		"Maria",
		"Pedro",
		"Carlos",
	}
	//Indice va ser la posición en la que este el nombre y nombre
	//va a tomar un valor de nombres
	//El range devuelve dos valores
	for index, nombre := range nombres {
		//Index = 0
		// nombre nombres[index]
		fmt.Printf("El nombre %s esta en el index %d\n", nombre, index)
	}
	//si no vamos a necesitar el indice lo desechamos el indice
	//con el subrayado sustituye al indice
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}
	//si solo se quiere usar el indice
	for index, _ := range nombres {
		fmt.Println(index)
	}

	cadena := "Hola gente, yo soy Mauricio Brito"
	for index, letra := range cadena {
		fmt.Printf("La letra %c esta en el index %d \n", letra, index)
	}
}
