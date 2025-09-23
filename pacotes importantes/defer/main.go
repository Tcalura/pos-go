package main

import "fmt"

func main() {
	// basicamente atrasa a execuçao para a ultima expressao da funçao
	// ideal para fechar requisiçao e fechar arquivos que esta sendo lido
	defer fmt.Println("Primeira Linha")
	fmt.Println("Segunda Linha")
	fmt.Println("Terceira Linha")
}