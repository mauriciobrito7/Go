package main

import (
	"fmt"
	"time"
)

func main() {
	numero := make(chan int)
	cuadrado := make(chan int)

	//Generamos los numeros de manera finita
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
		for {
			//El channel devuelve un segundo valor para saber si el canal esta cerrado
			x, ok := <-numero
			if !ok {
				break
			}
			cuadrado <- x * x
		}
		close(cuadrado)
	}()

	//Imprimimos en main

	for {
		x, ok := <-cuadrado
		if !ok {
			break
		}
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}
