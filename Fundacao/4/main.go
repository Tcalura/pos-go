package main

import "fmt"

const a = "Hello, Word!"

type ID int

var (
  b bool
  c int 
  d string
  e float64
  f ID = 1
)

func main() {
  fmt.Printf("O tipo de 'f' é %T\n", f)
}