package main

import (
	"fmt"
	"github.com/devfullcycle/goexpert/packaging/1/math"
)

func main() {
	fmt.Println(math.X)

	m := math.NewMath(1, 2)
	m.C = 5
	fmt.Println("Sum:", m.Sum())
	fmt.Println("o valor de C:", m.C)
}