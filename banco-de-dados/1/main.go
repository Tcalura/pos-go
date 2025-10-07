package main

import (
	"database/sql"
	"fmt"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "root"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3307"
	cfg.DBName = "goexpert"

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Notebook", 9999.0)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	p, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: %v, possui o preço de %.2f", p.Name, p.Price)

	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("Product: %v, possui o preço de %.2f\n", p.Name, p.Price)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func selectAllProducts(db *sql.DB) ([]*Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	fmt.Println("\n Produto deletado com sucesso id: ", id)
	return nil
}
