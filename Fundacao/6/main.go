package main

import "fmt"

func main() {
  // cria o slice
	s := []int{10, 20, 30, 50, 60, 70, 80, 90, 100}
  // printa o tamanho a capacidade e os itens
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
  // printa o tamanho a capacidade e os itens sao substituidos por vazio
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
  // printa o tamanho 4 a capacidade 9 e os itens nas 4 primeiras posiçoes 
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])
  // printa o tamanho a capacidade e os itens todos a partir do segundo (nao inclusivo)  
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

  // quando se da um append em um slice ja cheio ele cria um novo slice com
  // o dobro do tamanho independente da quantidade de itens adicionados,
  // isso pq por de baixo todo slice é um array 
	s = append(s, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])
}