package main

import (
	"fmt"
)
type Endereco struct {
	Logradouro string
	Numero int
	Cidade string
	Estado string
}

// composiçao de endereco em Ciente
type Client struct {
	Nome string
	Idade int
	Ativo bool
	//Address Endereco
	Endereco
}

func (c Client) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func main(){
	thiago := Client{
		Nome: "Thiago",
		Idade: 30,
		Ativo: true,
	}

	thiago.Desativar()
}
