package main

import (
	"fmt"
	"time"
)

//Mejor funcionalidad de los channels con un range
func main() {
	numero := make(chan int)
	cuadrado := make(chan int)

	//Generamos los numeros
	go func() {
		for x := 0; x < 5; x++ {
			numero <- x
		}
		//CERRAR EL CANAL PARA QUE NO ENVIE MÁS VALORES
		close(numero)
		/*Cuando se cierra un canal luego de el ultimo envio se empieza
		a enviar el valor 0 del tipo que se pasa a traves del chanel*/
	}()
	//Lo elevamos al cuadrado

	go func() {
		/*No ponemos <- para recibir valores del channel porque ya el range se encarga de eso
		le asigna el valor a x que es enviado a traves de ese channel*/
		for x := range numero {
			cuadrado <- x * x
		}
		close(cuadrado)
	}()

	//Imprimimos en main

	for x := range cuadrado {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}
