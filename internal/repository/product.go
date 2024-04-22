package repository

import (
	"errors"

	"github.com/ozykt4/andrade-api/internal/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Product) (*model.Product, error)
	FindAll() ([]model.Product, error)
	FindByID(id uint) (*model.Product, error)
	Update(product *model.Product) (*model.Product, error)
	Delete(id uint) (*model.Product, error)
}

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepo{
		db: db,
	}
}

func (p *ProductRepo) Create(product *model.Product) (*model.Product, error) {
	var count int64
	if err := p.db.Model(&model.Product{}).Where("name = ?", product.Name).Count(&count).Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, errors.New("product with the same NAME already exists")
	}

	if err := p.db.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductRepo) FindAll() (products []model.Product, err error) {
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductRepo) FindByID(id uint) (product *model.Product, err error) {
	if err := p.db.First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductRepo) Update(product *model.Product) (*model.Product, error) {
	if err := p.db.Save(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductRepo) Delete(id uint) (product *model.Product, err error) {
	if err := p.db.Where("id = ?", id).Delete(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}
