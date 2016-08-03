package main;

import "fmt";

func main(){
	//for clasico
	for j:=7;j<10;j++{
		fmt.Println(j);
	}
	//simulando un while
	i:=1;
	for i<10{
		fmt.Println(i);	
		i++;
	}

	for{
		fmt.Println("loop");
		break	
	}

}