package main

import (
	"fmt"
)

type User struct {
	edad             int
	nombre, apellido string
}

func (this User) nombre_completo() string {
	return this.nombre + " " + this.apellido
}

func (this *User) set_name(n string) {
	this.nombre = n
}
func (this *User) set_last_name(n string) {
	this.apellido = n
}
func main() {
	/*new() nos permite declarar una estructura y retorna
	un puntero*/
	Mauricio := new(User)
	Mauricio.set_name("Mauricio")
	Mauricio.set_last_name("Brito")
	fmt.Println("Bienvenido ", Mauricio.nombre_completo())

}
