package main

import "fmt"

//Números ímpares de 1 a 19

func main() {
	for i := 1; i <= 19; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}
