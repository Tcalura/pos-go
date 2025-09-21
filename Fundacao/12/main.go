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

// interface nao permite passar attributos apenas metodos
type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {
	// e.Nome = "XPTO"
}

type Client struct {
	Nome string
	Idade int
	Ativo bool
	//Address Endereco
	Endereco
}

func (c Client) Desativar() {
	c.Ativo = false
	// fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func Desativacao(pessoa Pessoa){
	pessoa.Desativar()
}

func main(){
	thiago := Client{
		Nome: "Thiago",
		Idade: 30,
		Ativo: true,
	}

	thiago.Desativar()

	// minhaEmpresa := Empresa{Nome: "Fake Empresa"}
	minhaEmpresa := Empresa{}
	Desativacao(minhaEmpresa)
	fmt.Println(minhaEmpresa.Nome)
}
