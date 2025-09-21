package matematica

// variaveis, funçoes e structs devem começar com maiusculas
// caso contratio nao serao exportadas para quem usa
// os pacotes

func Soma[T int | float64](a T, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
}

func (c Carro) Andar() string {
	return "Carro andando"
}