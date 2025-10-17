package math

var X string = "Hello, World!"

type math struct{
	a int
	b int
	C int // essa é uma propriedade exportada
}
// dessa forma eu estou exportando o struct Math
// e as propriedades A e B
// para acessar elas fora do pacote, eu preciso usar a letra maiuscula
func NewMath(a, b int) math {
	return math{a: a, b: b}
}

func (m math) Sum() int {
	return m.a + m.b
}