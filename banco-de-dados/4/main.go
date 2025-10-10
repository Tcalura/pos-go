package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3307)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// create category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	// create product
	db.Create(&Product{
		Name:       "Mouse",
		Price:      1000.00,
		CategoryID: 1,
	})

	// delete product
	// var p2 Product
	// db.First(&p2, "name = ?", "Mouse")
	// fmt.Println(p2.Name)
	// db.Delete(&p2)

	// exibe o nome do produto, nome da categoria e preço.
	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.Price)
	}
}
