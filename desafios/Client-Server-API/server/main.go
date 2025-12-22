package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type ExchangeRate struct {
	ID          int64     `json:"-" db:"id"`
	Value       float64   `json:"bid" db:"value"`
	Create_date time.Time `json:"create_date" db:"verification_date"`
}

func main() {
	db, err := ConfigDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/cotacao", HandlerExchangeRate)
	http.ListenAndServe(":8080", nil)
}

func ConfigDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS cotacoes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		value REAL NOT NULL,
		verification_date DATETIME NOT NULL
	); `

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	fmt.Println("DB created!")
	return db, nil
}

func HandlerExchangeRate(w http.ResponseWriter, r *http.Request) {

	// buscaCotacao
	exRate, err := GetExchangeRate()
  
	// SalvaCotacao

	// retorna json

}

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

	q.Value = bid
	q.Create_date = createDate

	return nil
}
