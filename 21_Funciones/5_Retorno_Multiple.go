package main

import "fmt"

//Funciones que retornan multiples valores
//Lo unico que hay que hacer para que una funcion retorne multiple valores
//es indicarle los valores que va a retornar
func multiplicar(numero int) (n1, n2, n3 int) {
	n1 = numero * 10
	n2 = numero * 20
	n3 = numero * 30
	return
}

func multiplicar2(numero int) (int, int, int) {
	n1 := numero * 10
	n2 := numero * 20
	n3 := numero * 30
	return n1, n2, n3
}

func retorno() (string, string) {
	return "Hola", "Mundo"
}

func main() {
	fmt.Println(multiplicar(1))
	//Recibir los retornos
	c1, c2, c3 := multiplicar(10)
	fmt.Println(c1, c2, c3)
	//Si solo quiero tomar uno de los retornos
	_, d2, _ := multiplicar(66)
	fmt.Println(d2)
	fmt.Println(retorno())

}
