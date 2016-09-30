package main

import (
	"fmt"
)

func main() {
	channel := make(chan string)
	//Siempre se pide información
	go func(channel chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}(channel)
	//Siempre se procesa la información
	for {
		msg := <-channel
		fmt.Println(msg)
	}
}
