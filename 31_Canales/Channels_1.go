package main

import (
	"fmt"
	"time"
)

/*CANALES
Si las goroutines son las actividades concurrentes en Go los channels
son las comunicacione sentre ellas
*/

func imprimirPing(c chan string) {
	var contador int
	for {
		//Recibiendo valores a través del canal
		contador++
		fmt.Println(<-c, " ", contador)
		time.Sleep(time.Second * 1)
	}
}

func enviarPing(c chan string) {
	for {
		//Enviando valores a través del canal
		c <- "Ping" //Channel de tipo unbuffered
		/*Detiene la goroutina hasta que el canal este listo para recibir*/
	}
}

func main() {
	//Creamos un canal hola mundo
	/*Para crear un channel se hace con la función make al igual que los slice
	el channel que se crea hace referencia a la estructura de datos que creo la
	función make, por lo cual cuando se esta pasando una channel por parametro
	se esta pasando por referencia y no por valor*/
	/*Los channel tienen dos operaciones principales que son la de mandar y recibir
	que en conjunto se conoce como comunicaciones*/
	ch := make(chan string) //Un channel le permite mandar el tipo de dato cuando lo
	//construyo con make

	//Llamamos las funciones como Goroutines
	go enviarPing(ch)
	go imprimirPing(ch)

	//Escaneamos la entrada de datos para que no finalice la goroutine "main"
	var input string
	fmt.Scanln(&input)

}
