package main

import (
	"fmt"
	"math/rand"
	"sync" /*Para trabajar con Gorutinas sincronizar que el programa espere
	que la Gorutina se ejecute antes de cerrar*/
	"time" /*Para trabajar con el tiempo*/
)

//Goroutines
//Un hilo ligero manejado por el runtime de go

//wg es utilizado para indicarle al programa que debe esperar
//que finalicen las gorutinas
/*Nos permite controlar que cantidad de Gorutinas debe esperar el programa que se
ejecuten para salir del programa*/
var wg sync.WaitGroup

func main() {
	//Añadimos 2 wg para que espere que finalicen 2 gorutinas
	wg.Add(2)

	fmt.Println("Iniciamos las gorutinas..")
	//Lanzamos la gorutina con la etiqueta "A"
	/*Cuando le ponemos go alante a la función le estamos diciendo al programa
	que es una gorutina*/
	go imprimirCantidad("A")
	//Lanzamos la gorutina con la etiqueta "B"
	go imprimirCantidad("B")
	//Esperamos a que las gorutinas finalicen.
	fmt.Println("Esperando que Finalicen..")
	wg.Wait()
	fmt.Println("\n Terminando el programa")
}

func imprimirCantidad(etiqueta string) {
	//Llamamos la función Done() de wg para indicarle que la gorutina termino.
	//Con Done se va restando en 1 la cantidad del waitgroup
	defer wg.Done()
	//Espera aleatoria
	for cantidad := 1; cantidad <= 10; cantidad++ {
		//Indicamos que tenga un número random
		//Devuelve un número random de 0-1000
		sleep := rand.Int63n(60)
		//Sleep detiene la ejecución por un tiempo indicado
		//Pueden ser microsegundos, milisegundos, segundos, minutos..
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("Cantidad", cantidad, " de ", etiqueta)
	}
}
