package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ExchangeRate struct {
	ID          int64     `json:"-" gorm:"id, primaryKey"`
	Bid         float64   `json:"bid" gorm:"value"`
	Create_date time.Time `json:"create_date" gorm:"verification_date"`
}

func main() {
	http.HandleFunc("/cotacao", HandlerExchangeRate)
	http.ListenAndServe(":8080", nil)
}

func HandlerExchangeRate(w http.ResponseWriter, r *http.Request) {
	// buscaCotacao
	exRate, err := GetExchangeRate()
	if err != nil {
		log.Println(err)
	}
	// SalvaCotacao
	err = saveExchangeRate(exRate)
	if err != nil {
		log.Println(err)
	}
	// retorna json
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exRate)

}

// ta funcionando
func GetExchangeRate() (*ExchangeRate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	//cria request com context
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	//executa request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Println("Timeout ao buscar cotação: limite de 200ms excedido")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var exRate ExchangeRate
	if err := json.Unmarshal(body, &exRate); err != nil {
		return nil, err
	}

	return &exRate, nil
}

// ta funcionando
func (q *ExchangeRate) UnmarshalJSON(data []byte) error {
	// temp struct
	var raw struct {
		USDBRL struct {
			Bid        string `json:"bid"`
			CreateDate string `json:"create_date"`
		} `json:"USDBRL"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("failed unmarshal to JSON: %w", err)
	}

	bid, err := strconv.ParseFloat(raw.USDBRL.Bid, 64)
	if err != nil {
		return fmt.Errorf("failed to convert to float64: %w", err)
	}

	createDate, err := time.Parse("2006-01-02 15:04:05", raw.USDBRL.CreateDate)
	if err != nil {
		return fmt.Errorf("failed to convert date: %w", err)
	}

	q.Bid = bid
	q.Create_date = createDate

	return nil
}

func saveExchangeRate(exchangeRate *ExchangeRate) error {
	db, _ := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})

	db.AutoMigrate(&ExchangeRate{})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	result := db.WithContext(ctx).Create(&exchangeRate)

	if result.Error != nil {
		if errors.Is(result.Error, context.DeadlineExceeded) {
			log.Println("Timeout ao salvar cotação: limite de 10ms excedido")
		} else {
			log.Println("Problema ao salvar no banco de dados")
		}
    log.Printf("Erro técnico: %v\n", result.Error)
		return result.Error
	}
	return nil
}
