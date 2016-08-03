package main;
import (
		"fmt"
		"time"
);

func main(){
	i:=1;
	switch i{
		case 1:
			fmt.Println("Uno");
		case 2:
			fmt.Println("Dos");
		case 3:
			fmt.Println("Tres");
	}

	switch time.Now().Weekday(){
		case time.Saturday, time.Sunday:
			fmt.Println("Es fin de semana");
		default:
			fmt.Println("Es día de la semana");
	}
	t:=time.Now();
	switch{
		case t.Hour()<12:
			fmt.Println("Antes del medio día");
		default:
			fmt.Println("Despues del medio día");
	}
}