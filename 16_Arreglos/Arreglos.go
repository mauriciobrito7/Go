package main;
import "fmt";

func main(){
	//Declaración de un arreglo
	var a[5] int;
	//imprime el vector que por defecto esta inicializado en 0
	fmt.Println("emp:",a);

	a[4]=100;
	fmt.Println("set:",a);
	fmt.Println("get",a[4]);

	//Dimension del arreglo
	fmt.Println("len:",len(a));
	//Declarar un arreglo inicializado
	b:=[5] int{1,2,3,4,5}
	fmt.Println("dcl:",b);

	var twoD[2][3]int;
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			twoD[i][j] = i + j;
		}
	}

	fmt.Println("2D:",twoD);
}