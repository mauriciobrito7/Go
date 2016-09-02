package main

import (
	"fmt"
)

//Type
//Declarar nuestro propio tipo

//Dinero Tipo declarado por nosotros
//Tipo dinero que parte del tipo int
type Dinero int

//Declaramos el metodo String para el tipo Dinero
/*La diferencia de crear tipos con type es que se
puden modificar los tipos agregandole metodos*/
func (d Dinero) String() string {
	//Sprintf nos permite devolver una cadena con formato
	return fmt.Sprintf("$%d", d)
}
func main() {
	/*Con la variable sueldo se puede hacer todas las
	operaciones que se hace con un int*/
	var sueldo Dinero
	sueldo = 25000
	fmt.Println("Sueldo: ", sueldo)

	aumento := 10000
	sueldo += Dinero(aumento)
	fmt.Println("Sueldo + Aumento", sueldo)

}
