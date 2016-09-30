package main

import (
	"errors"
	"fmt"
)

//Errores
//Creando variables de tipo error
var (
	//ErrorUsuarioNoValido...
	ErrorUsuarioNoValido = errors.New("El usuario no es valido")
	//ErrorUsuarioEnProceso...
	ErrorUsuarioEnProceso = errors.New("El usuario esta en proceso de registro")
	//ErrorPorDefecto...
	ErrorPorDefecto = errors.New("Algo paso aqui..")
)

type error interface {
	Error() string
}

//También se puede declarar una estructura de tipo error
type PathError struct {
	Op   string //"open", "unlink", etc.
	Path string //The associated file.
	Err  error  //Returned by the system call.
}

func (e *PathError) error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func baneado(usuario string) (err error) {
	ban := false
	switch usuario {
	case "miguel":
		ban = true
	case "carlos":
		ban = false
	case "juan":
		/*Devuelve el error formateado con el string que le pasemos*/
		return ErrorUsuarioNoValido
	case "pedro":
		return ErrorUsuarioEnProceso
	default:
		return ErrorPorDefecto
	}

	if !ban {
		fmt.Printf("El usuario %s no esta baneado \n", usuario)
	} else {
		fmt.Printf("El usuario %s esta baneado \n", usuario)
	}

	return nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error :", err)
	}
}

func main() {
	err := baneado("miguel")
	checkError(err)
	err = baneado("juan")
	checkError(err)
	err = baneado("carlos")
	checkError(err)
	err = baneado("pedro")
	checkError(err)
	err = baneado("pololo")
	checkError(err)

}
