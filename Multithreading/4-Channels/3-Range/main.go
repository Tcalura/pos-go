package main

import "fmt"

// Thread 1
func main() {
  ch := make(chan int)
  go publish(ch)
  reader(ch)
}

func reader(ch chan int) {
  for x := range ch {
    fmt.Printf("Received %d\n", x)
  }
}

func publish(ch chan int) {
  // defer close(ch) // remove o close da linha 23
  for i := range 10 {
    ch <- i
  }
  close(ch)
}
