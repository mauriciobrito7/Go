package main

import "fmt"

//Funciones vacias
func imprimirNombre(nombre string) {
	fmt.Println("Función fuera de main")
	fmt.Println("El nombre es:", nombre)
}

//Funciones con retorno
func suma(n1 int, n2 int) int {
	return n1 + n2
}

//Otra forma de declarar funciones
func resta(n1, n2 int) (r int) {
	r = n1 - n2
	return
}

func main() {

	//Funciones
	//Llamar funciones
	imprimirNombre("Mauricio")
	fmt.Println(suma(2, 3))
	fmt.Println(resta(2, 3))

	//Firma de una función
	//Si dos o más funciones tienen el mismo cuerpo y devuelven el mismo
	//tipo se dice que tienen la misma firma.
	//%T imprime la firma de la función
	fmt.Printf("Función suma:%T\n", suma)
	fmt.Printf("Función resta:%T\n", resta)

}
