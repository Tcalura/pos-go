package main

const a = "Hello, Word!"

var (
  b bool
  c int 
  d string
  e float64
)

func main() {
  // Go infere valores iniciais para qualquer tipo de variavel
  println(b)
  println(c)
  println(d)
  println(e)
}