package main

import "fmt"


// Nao é obrigatorio para trabalhar com channels
// a seta no param hello informa que essa funçao apenas trabalha com receive-only
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// a seta no param data informa que essa funçao apenas trabalha com send-only
func ler(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	ler(hello)
}
