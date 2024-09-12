package main

import "fmt"

//Operação com inteiros

func main() {
	var A, B, C, D, oper int
	fmt.Println("Inteiros A, B, C e D: ")
	fmt.Scan(&A, &B, &C, &D)
	oper = (A*B - C*D)
	fmt.Println("A diferença do produto de C e D é", oper)
}
