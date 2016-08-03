package main;
import "fmt";

func main(){
	encabezado:=`
	*****************************
	 Contador de numeros impares
	*****************************
	`;	
	fmt.Println(encabezado);
	fmt.Println("Ingrese el primer número");
	var numero1 int;
	fmt.Scanln(&numero1);

	fmt.Println("Ingrese el segundo número");
	var numero2 int;
	fmt.Scanln(&numero2);
	var contador int;
	for i:=numero1;i<numero2;i++{
		if(i%2!=0){
			contador++;
			fmt.Println("El número",i," es impar ");
		}
	}
	encabezado=`
	**********************
	 Resultado del conteo
	**********************
	`;
	fmt.Println(encabezado);
	fmt.Println("La cantidad de impares es:",contador);	
}

