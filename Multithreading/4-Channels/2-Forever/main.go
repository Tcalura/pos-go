package main

// Thread 1
func main() {
	forever := make(chan bool)

	// Thread 2
	go func() {
		for i := range 10 {
			println(i)
		}
		// preenche o channel pela thread 2 (esse cara nao pode ser preenchido pela thread 1)
		forever <- true
	}()

	// esvazia a thread
	<-forever
}
