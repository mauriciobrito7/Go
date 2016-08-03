package main;
import (
	"fmt"
	"strconv"
);

func main(){
	edad:=21;
	fmt.Println("Mi edad es ",edad);
	//convertir un entero a string
	edad_str:=strconv.Itoa(edad);
	fmt.Println("Mi edad es "+edad_str);
	//convertir un string a un entero 
	cantidad:="50";
	cantidad_int,_:=strconv.Atoi(cantidad);
	//La función Atoi retorna multiiple valores
	//Se usa el _ para recibir un valor que no se va usar

	fmt.Println(cantidad_int +10);
}