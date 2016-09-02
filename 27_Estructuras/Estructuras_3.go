package main

import (
	"fmt"
	"math"
)

//Rectangulo tiene los campos ancho y alto
type Rectangulo struct {
	ancho, alto float64
}
type Circulo struct {
	radio float64
}

//Métodos
/*Creamos como si fuera una función lo unico es que antes del nombre de
la función agregamos un recibidor el cual recibe una variable del tipo
que se va asociar este método*/
func (r Rectangulo) area() float64 {
	return r.ancho * r.alto
}

/*Pueden haber métodos con nombres iguales pero con distintos ambitos*/
func (c Circulo) area() float64 {
	return c.radio * c.radio * math.Pi
}

//Retorna un rectangulo modificado
/*func (r Rectangulo) inc(i float64) Rectangulo {
	return Rectangulo{
		r.ancho * i,
		r.alto * i,
	}
}*/

//Función mejorada de incrementar
func (r *Rectangulo) inc(i float64) {
	r.ancho *= i
	r.alto *= i
}

func main() {
	//Declaramos dos rectangulos
	r1 := Rectangulo{12, 2}
	r2 := Rectangulo{9, 4}
	//Calculamos e imprimimos sus areas
	fmt.Println("Area de r1 es: ", r1.area())
	fmt.Println("Area de r2 es: ", r2.area())
	//Declaramos dos circulos
	c1 := Circulo{10}
	c2 := Circulo{25}
	//Calculamos e imprimimos sus areas
	fmt.Println("Area de c1 es: ", c1.area())
	fmt.Println("Area de c2 es: ", c2.area())

	//Incrementar rectangulo
	fmt.Println("r1:", r1)
	r1.inc(10)
	fmt.Println("Nuevo r1:", r1)

}
