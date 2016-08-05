//Escriba un algoritmo que permita recorrer la diagonal secundaria
//de una matriz cuadrada
package main;
import "fmt";

func main(){
	m:=[4][4]int{{1,2,3,4},{5,6,7,8},{9,10,11,12},{13,14,15,16}};
	for i:=0;i<4;i++{
		for j:=0;j<4;j++{
			if i==len(m)-1 && j==0 || j==len(m)-1 && i==0 || i+j==len(m)-1 {
				fmt.Println(m[i][j]);
			}
		}
	}
}