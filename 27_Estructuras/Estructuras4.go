package main

import (
	"fmt"
)

//Punto contiene los valores x,y
type Punto struct {
	x, y int
}

type Punto3D struct {
	x, y, z int
	/*Una estructura no se puede tener como campo asi misma, lo que si
	permite go es que una estructura contenga un puntero hacia otra estructura
	de su mismo tipo o de si misma*/
	//Punto3D
	*Punto3D
}

//Se pueden declarar estructuras sin campos
/*Por lo general cuando se declara una estructura sin campos es porque
se le van a agregar métodos que van a ser utilizados desde esa estructura*/
//OpPunto tiene los métodos para realizar operaciones con puntos
type OpPunto struct {
}

func main() {
	p := Punto{}
	fmt.Println(p.x, p.y)
	p2 := Punto3D{
		5,
		6,
		4,
		&Punto3D{
			6,
			4,
			6,
			nil,
		},
	}
	fmt.Println("p2: ", p2)

	//Comparando estructuras
	a := Punto{5, 6}
	b := Punto{7, 4}
	/*Para que dos estructuras sean true cuando se esten comparando cada uno
	de sus respectivos campos deben ser iguales */
	fmt.Println("a==b:", a == b)
	c := Punto{7, 4}
	fmt.Println("b==c: ", b == c)

	/*Como las estructuras son comparables pueden entrar como indices en los
	mapas*/
	figuras := make(map[Punto]string)
	figuras[a] = "Hola mundo"
	fmt.Println("figuras[a]: ", figuras[a])
}
