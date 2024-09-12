package main

import "fmt"

//Função que retorna a multiplicação de dois valores

func main() {
	var x, y int
	fmt.Println("Informe dois valores: ")
	fmt.Scan(&x, &y)
	fmt.Printf("%d x %d = %d", x, y, soma(x, y))
}

func soma(x int, y int) int {
	return x * y
}
