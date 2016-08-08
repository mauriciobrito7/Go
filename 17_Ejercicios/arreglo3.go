//Escriba un algoritmo que permita recorrer la tringular superior con respecto a la diagonal principal de una matriz cuadrada.
//Escriba un algoritmo que permita recorrer la tringular inferior con respecto a la diagonal secundaria de una matriz cuadrada.
package main

import (
	"fmt"
)

func main() {
	m := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Println(m)
	//Triangular superior
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j == i || j > i {
				fmt.Println(m[i][j])
			}
		}
	}
	//Triangular inferior
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j == i || j < i {
				fmt.Println(m[i][j])
			}
		}
	}

}
