package main

import "fmt"

func main() {
	salarios := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	delete(salarios, "Wesley")
	salarios["Wes"] = 5000

	// print com ambos os dados de chave e valor
	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	// print ignorando o primeiro valor que se refere a cada chave do map
	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}

	// print ignorando o segundo valor que se refere a cada valor do map
	for nome, _ := range salarios {
		fmt.Printf("O nome é %s\n", nome)
	}
}