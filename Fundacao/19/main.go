package main

// correlato a aula 18 / 19 / 20 

import (
	"fmt"
	"curso-go/matematica"
	// "github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}

	fmt.Println(carro.Andar())
	fmt.Println("Resultado: ", s)
	fmt.Println(matematica.A)
}