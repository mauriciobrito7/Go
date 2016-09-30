package main

import (
	"fmt"
	"time"
)

func main() {
	//Se ejecuta la animación hasta que se imprima el resultado
	go animacion(100 * +time.Millisecond)
	const N = 45
	resultado := fib(N)
	fmt.Printf("\r Fibonacci(%d)=%d\n", N, resultado)
}

func animacion(retraso time.Duration) {
	//Bucle infinito
	for {
		//for each
		//Usamos range para recorrer el string infinitamente
		for _, r := range `\|/-` {
			fmt.Printf("\r %c", r)
			time.Sleep(retraso)
		}
	}
}

func fib(x int) int {

	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
