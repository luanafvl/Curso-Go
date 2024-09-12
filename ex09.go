package main

import "fmt"

//Divisores de um número inteiro positivo

func main() {

	var n int

	fmt.Println("Digite um inteiro positivo:")
	fmt.Scan(&n)
	fmt.Println("Divisores:")
	for i := n; i >= 1; i-- {
		if n%i == 0 {
			fmt.Println(i)
		}
	}
}
