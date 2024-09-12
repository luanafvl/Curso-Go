package main

import "fmt"

//Guardar números e printar o maior deles quando o 0 for escaneado

func main() {
	var num int
	var maior int
	
	num=1
	maior=-1000000000000000
	
	for {
	    fmt.Println("Digite um número")
		fmt.Scan(&num)
	    if num==0 {
            break
	} else {
	    if num>maior {
	    maior = num    
	        }
	    }
    }
    fmt.Println(maior)
}
