package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	QtdWorkers := 1000000

	// inicializa os workers
	for i := 0; i < QtdWorkers; i++ {
		go worker(i, data)
	}

  // envia as informaçoes para o channel
  // de forma serial basicamente
	for i := 0; i < 10000000; i++ {
		data <- i
	}
}
