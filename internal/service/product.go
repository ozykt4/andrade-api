package service

import (
	"github.com/ozykt4/andrade-api/internal/model"
	"github.com/ozykt4/andrade-api/internal/repository"
)

type IProductService interface {
	CreateProduct(productReq *model.ProductReq) *model.Response
	FindAllProducts() *model.Response
	FindProductByID(id uint) *model.Response
	UpdateProduct(id uint, productReq *model.ProductReq) *model.Response
	DeleteProduct(id uint) *model.Response
}

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) IProductService {
	return &ProductService{
		repo: repo,
	}
}

func (ps *ProductService) CreateProduct(productReq *model.ProductReq) *model.Response {
	product := productReq.ToProduct()

	createProduct, err := ps.repo.Create(product)
	if err != nil {
		return model.NewErrorResponse(err, 500)
	}

	return model.NewSuccessResponse(createProduct.ToProductRes())
}

func (ps *ProductService) FindAllProducts() *model.Response {
	products, err := ps.repo.FindAll()
	if err != nil {
		return model.NewErrorResponse(err, 404)
	}

	productsRes := []*model.ProductRes{}
	for _, p := range products {
		productsRes = append(productsRes, p.ToProductRes())
	}

	return model.NewSuccessResponse(productsRes)
}

func (ps *ProductService) FindProductByID(id uint) *model.Response {
	product, err := ps.repo.FindByID(id)
	if err != nil {
		return model.NewErrorResponse(err, 404)
	}

	return model.NewSuccessResponse(product.ToProductRes())
}

func (ps *ProductService) UpdateProduct(id uint, productReq *model.ProductReq) *model.Response {
	product, err := ps.repo.FindByID(id)
	if err != nil {
		return model.NewErrorResponse(err, 404)
	}

	product.Name = productReq.Name
	product.Description = productReq.Description
	product.Price = productReq.Price

	productUpdate, err := ps.repo.Update(product)
	if err != nil {
		return model.NewErrorResponse(err, 500)
	}

	return model.NewSuccessResponse(productUpdate.ToProductRes())
}

func (ps *ProductService) DeleteProduct(id uint) *model.Response {
	_, err := ps.repo.Delete(id)
	if err != nil {
		return model.NewErrorResponse(err, 404)
	}

	return model.NewSuccessResponse(nil)
}
