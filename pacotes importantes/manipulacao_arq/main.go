package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	// Cria o arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// formas de escrever no arquivo
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	// tamanho, err := f.WriteString("Hello, world!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes", tamanho)
	f.Close()

	// ler o arquivo
	// Existe o metodo Open
	arq, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("\n"+string(arq))

	// leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// deleta o arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}