package main

import "fmt"

//Múltiplos de 3 de 0 a 30

func main() {
	for i:=0; i<=30; i++ {
		if i%3==0 {
			fmt.Println(i)
		}
	}
}
