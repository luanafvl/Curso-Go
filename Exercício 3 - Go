package main

import "fmt"

//Cálculos com switch case

func main() {
    var opt int
	var A, B, C float64
	pi:=3.14159
	
	fmt.Println("Selecione uma opção:\n1.A área do triângulo retângulo que tem A por base e C por altura\n2.A área do círculo de raio C\n3.A área do trapézio que tem A e B por bases e C por altura\n4.A área do quadrado que tem lado B\n5.A área do retângulo que tem lados A e B")
	fmt.Scan(&opt)
	fmt.Println("Agora, informe os valores para as variáveis A, B e C:")
	fmt.Scan(&A, &B, &C)
	
	switch opt {
	    case 1:
	        fmt.Println("Resultado = ", A*C)
        case 2:
	        fmt.Println("Resultado = ", (C*C)*pi)
	    case 3:
	        fmt.Println("Resultado = ", (A+B)*C/2)
	    case 4:
	        fmt.Println("Resultado = ", B*B)
	    case 5:
	        fmt.Println("Resultado = ", A*B)
    }
}
