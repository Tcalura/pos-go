package main

// correlato a aula 17

type MyNumber int

type Number interface {
	~int | ~float64
}

// o que esta entre [] é chamado de generics
func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m{
		soma += v
	}
	
	return soma
}

// comparable é uma constraints vide doc do go sobre constraints
func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Wesley": 1000.2, "João": 2000.2, "Maria": 3000.2}
	m3 := map[string]MyNumber{"Wesley": 1000, "João": 2000, "Maria": 3000}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))
}