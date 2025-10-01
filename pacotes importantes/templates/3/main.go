package main

import (
	"os"
	"html/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Python", 30},
		{"Java", 50},
	})

	if err != nil {
		panic(err)
	}
}
