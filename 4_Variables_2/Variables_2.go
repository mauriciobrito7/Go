package main;
import "fmt";

func main(){
	//Declaración multiple
	var nombre,nombre2 string;
	nombre="Mauricio";
	nombre2="Eduardo";	
	//Permutación sin necesidad de declar una tercera variables 
	nombre,nombre2= nombre2,nombre;
	fmt.Println(nombre, nombre2);
	var(
		Dios="Goku";
		enemigo1, enemigo2="Freezer", "Cell";
		);
	//Se usa el duck typing y no es necesario decir el tipo de variable
	fmt.Println(Dios,enemigo1,enemigo2);

	//Scope
	var razaSaiyajin="Vegeta";
	fmt.println("Un saiyajin es "+ razaSaiyajin);
}