package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:product_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:product_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3307)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&Product{}, &Category{})

	// create category
	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)

	// category2 := Category{Name: "Cozinha"}
	// db.Create(&category2)

	// // create product
	// db.Create(&Product{
	// 	Name:       "Microondas",
	// 	Price:      2999.0,
	// 	Categories: []Category{category, category2},
	// })

	var categories []Category
	// para conseguir vizualizar os produtos de cada categoria, precisa usar o Preload
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name)
		}
	}
}
