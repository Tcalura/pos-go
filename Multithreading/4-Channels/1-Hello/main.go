package main

import "fmt"

// Thread 1
func main() {
  canal := make(chan string) // channel vazio

  // Thread 2
  go func() {
    canal <- "Olá Mundo!" // channel preenchido
  }()

  // Thread 1
  msg := <- canal // channel esvazia
  fmt.Println(msg)
}
