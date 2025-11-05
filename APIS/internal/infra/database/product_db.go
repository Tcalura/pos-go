package database

import (
	"mod-apis/internal/entity"

	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) (*entity.Product, error) {
	result := p.DB.Create(product)
	return product, result.Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]*entity.Product, error) {
	var products []*entity.Product
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		offset := (page - 1) * limit
		err = p.DB.Limit(limit).Offset(offset).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}
	return products, err
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product //TESTAR ESSE CASO product := entity.Product{}
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *Product) Update(product *entity.Product) error {
	product, err := p.FindByID(string(product.ID))
	if err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}
