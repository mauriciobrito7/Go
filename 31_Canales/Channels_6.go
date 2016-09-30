package main

import (
	"fmt"
	"time"
)

/*Indicar si un channel es de entrada o de salida
channel undireccional de salida
out chan<- int
*/
func generarNumeros(out chan<- int) {
	for x := 0; x < 5; x++ {
		out <- x
	}
	close(out)
}

func elevarAlcuadrado(in chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

/*channel de entrada
in <-chan int
*/
func imprimir(in <-chan int) {
	for x := range in {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
	/*No se puede cerrar el channel cuando recibe datos
	solo puede ser cerrado cuando envia datos*/
}

func main() {
	//Tenemos el channel numero y cuadrado
	numero := make(chan int)
	cuadrado := make(chan int)
	/*Cuando se pasa como parametro Go hace el cambio de doble via
	a un channel de una sola via ya sea para recibir o para enviar datos*/
	go generarNumeros(numero)
	go elevarAlcuadrado(numero, cuadrado)

	imprimir(cuadrado)

}
