package main

import (
	"fmt"
)

//Interface
/*Son unas herramientas de programación que nos permite separar la definición
de los objetos de su implementación en Go las interfaces son fundamentales
se utilizan mucho ya que no existen clases, se trabaja con composición
*/
//Persona...
type Persona struct {
	nombre string
	email  string
	edad   int
}

//Moderador..
type Moderardor struct {
	Persona
	Foro string
}
type Administrador struct {
	Persona
	Seccion string
}

//Metodos................

//Nombre..
func (p Persona) Nombre() string {
	return p.nombre
}

//Email..
func (p Persona) Email() string {
	return p.email
}

//AbrirForo..
func (m Moderardor) AbrirForo() {
	fmt.Println("Abrir foro")
}

//CerrarForo..
func (m Moderardor) CerrarForo() {
	fmt.Println("Cerrar foro")
}

//CrearForo
func (a Administrador) CrearForo() {
	fmt.Println("Creando foro")
}

//EliminarForo
func (a Administrador) EliminarForo() {
	fmt.Println("Eliminar foro")
}

//Metodo de la interface...
func Presentarse(u Usuario) {
	fmt.Println("Nombre: ", u.Nombre())
	fmt.Println("Email: ", u.Email())
}

//Declaración de una interface
/*Se declaran igual que las estructuras y se ponen los metodos que
va a implementar esa interface, los metodos no se codifican, solo se
deben indicar los metodos que debe tener un tipo para que sea incluido
como esa interface*/
type Usuario interface {
	Nombre() string
	Email() string
}

/*Cuando una estructura en este caso (Persona) implementa los metodos
declarados en una interface ya automaticamente esa estructura esta aplicando
a esa interface*/
func main() {
	mauricio := Persona{"Mauricio", "mauricio_brito@hotmail.com", 21}
	Presentarse(mauricio)
	juan := Moderardor{Persona{"Juan", "xxxxx@hotmail.com", 34}, "Juegos"}
	pedro := Administrador{Persona{"Pedro", "yyyyyy@gmail.com", 25}, "Pagina"}
	Presentarse(juan)
	Presentarse(pedro)

	//Podemos declarar una variable de tipo de la interface
	var i Usuario
	i = mauricio
	fmt.Println("i: ", i)
	fmt.Println("i :", i.Email())
	fmt.Println("i:", i.Nombre())

	/*Pero solo se tiene acceso a los metodos que son declarados
	dentro de la interface no del tipo de la estructura*/
	i = juan

	fmt.Println("i: ", i)
	fmt.Println("i :", i.Email())
	fmt.Println("i:", i.Nombre())

}
