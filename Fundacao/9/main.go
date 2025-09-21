package main

import (
	"fmt"
)


func main(){
	// soma deve retornar 15
	// funçao variatica, recebe como parametro uma lista do mesmos tipos
	// de tamanho indefinido
	fmt.Println(sum(1,2,3,4,5))
}

func sum(numeros ...int) (int) {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}