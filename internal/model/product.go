package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
}

type ProductReq struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductRes struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAT   time.Time `json:"created_at"`
	UpdatedAT   time.Time `json:"updated_at"`
}

func (p *ProductReq) ToProduct() *Product {
	return &Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}
}

func (pr *Product) ToProductRes() *ProductRes {
	return &ProductRes{
		ID:          pr.ID,
		Name:        pr.Name,
		Description: pr.Description,
		Price:       pr.Price,
		CreatedAT:   pr.CreatedAt,
		UpdatedAT:   pr.UpdatedAt,
	}
}
