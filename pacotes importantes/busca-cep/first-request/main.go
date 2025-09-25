package main 

import "net/http"

func main() {
	// cria uma rota e aponta qual metodo executa a essa rota
	http.HandleFunc("/", BuscaCep)
	
	// disponibiliza o server web 
	// control + c interrompe
	http.ListenAndServe(":8080", nil)
}

func BuscaCep(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}