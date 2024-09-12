package main

import "fmt"

//Localização  no plano cartesiano

func main() {
	var x int
	var y int
	fmt.Println("Digite a sua localização em coordenadas: ")
	fmt.Scan(&x, &y)

	if x > 0 && y > 0 {
		fmt.Println("Você está no primeiro quadrante.")
	} else if x < 0 && y > 0 {
		fmt.Println("Você está no segundo quadrante.")
	} else if x < 0 && y < 0 {
		fmt.Println("Você está no terceiro quadrante.")
	} else if x > 0 && y < 0 {
		fmt.Println("Você está no quarto quadrante.")
	} else if x == 0 && y != 0 {
		fmt.Println("Você está sobre o eixo Y.")
	} else if x != 0 && y == 0 {
		fmt.Println("Você está sobre o eixo X.")
	} else {
		fmt.Println("Você está sobre a origem.")
	}
}
