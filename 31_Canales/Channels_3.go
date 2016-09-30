package main

import (
	"fmt"
	"time"
)

//unbuffered channels
/*Hace que el recibidor y el que envia el mensaje se sincronicen ya que una tiene
que esparar a la otra*/
func main() {
	numero := make(chan int)
	cuadrado := make(chan int)

	//Generamos los numeros
	go func() {
		for x := 0; ; x++ {
			numero <- x
		}
	}()
	//Lo elevamos al cuadrado
	/*Se puede decir que el número antes que se eleve al cuadrado se va
	a generar*/
	go func() {
		for {
			x := <-numero
			cuadrado <- x * x
		}
	}()

	//Imprimimos en main
	/*Y se puede decir que el número antes de que sea impreso se va a elevar al
	cuadrado*/
	for {
		fmt.Println(<-cuadrado)
		time.Sleep(1 * time.Second)
	}
}
