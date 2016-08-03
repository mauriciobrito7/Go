package main;
import(
	"fmt"
	"bufio"
	"os"	
);

func main(){
	edad:=22;
	nombre:="Mauricio";
	activo:=true;
	sueldo:=50000.48;
	fmt.Printf("Nombre: %s \n",nombre);
	fmt.Printf("Edad: %d \n", edad);
	fmt.Printf("Activo: %t \n", activo);
	fmt.Printf("Sueldo: %.2f \n", sueldo);

	fmt.Println("*****Nuevos datos*****");
	fmt.Println("Ingrese el nombre");
	fmt.Scanf("%s\n",&nombre);
	fmt.Println("Ingrese la edad");
	fmt.Scanf("%d\n",&edad);
	fmt.Println("Ingrese el sueldo");
	fmt.Scanf("%f\n",&sueldo);

	fmt.Printf("Nombre: %s \n",nombre);
	fmt.Printf("Edad: %d \n", edad);
	fmt.Printf("Activo: %t \n", activo);
	fmt.Printf("Sueldo: %.2f \n", sueldo);

	//Otra forma de leer datos
	reader:= bufio.newReader(os.Stdin);
	fmt.Println("Ingresa tu nombre :");
	nombre,err= reader.ReadString('\n');
	if err==nil{
		fmt.Println(err);
	}else{
		fmt.Println("Hola"+nombre);
	}

}