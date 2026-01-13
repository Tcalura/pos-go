package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type ExchangeRate struct {
	Value       float64   `json:"value"`
	Create_date time.Time `json:"create_date"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Println("Erro na requisiçao com server")
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Erro na requisiçao com server")
		panic(err)
	}
	defer res.Body.Close()

	if ctx.Err() == context.DeadlineExceeded {
		log.Println("Timeout ao buscar cotação: limite de 300ms excedido")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		println(err)
	}

	var exchangeRate ExchangeRate
	err = json.Unmarshal(body, &exchangeRate)
	if err != nil {
		println(err)
	}

	// escreve no arquivo
	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write([]byte(fmt.Sprintf("Dólar: {%f}", exchangeRate.Value)))
	if err != nil {
		panic(err)
	}
}
