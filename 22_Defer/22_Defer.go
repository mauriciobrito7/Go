package main

import (
	"fmt"
	"os" //tiene acceso a funciones del sistema operativo
	//no importa en que S.O se ejecute
)

//Defer
func primera() {
	fmt.Println("Primera")
}

func segunda() {
	fmt.Println("Segunda")
}

func main() {
	//Hace que se ejecute si o si despues que termine todo el hilo
	//de ejecución de la función main
	//Es muy util para cuando se abren archivos o recursos externos
	defer primera()
	segunda()

	//Abrimos el archivo
	f, err := os.Open("texto.txt")
	//Verificamos que no haya ocurrido ninguún error
	if err != nil {
		panic(err)
	}
	//No importa lo que pase con el resto de código defer va a asegurar
	//de ejecutar lo que esta definido en este caso se asegura de que se
	//cierre el archivo.
	defer f.Close()
	//Creamos un Slice para almacenar loque leemos del archivo
	data := make([]byte, 11)
	c, err := f.Read(data)
	//verificamos que no haya ocurrido ninguún error
	if err != nil {
		panic(err)
	}
	//Imprimos lo leido
	fmt.Printf("Cantidad de byte leidos: %d\n Texto leido: \n %q \n error: %v", c, data, err)
	//Los defer son en forma de pila
	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
