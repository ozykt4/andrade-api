package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ozykt4/andrade-api/internal/api/router"
	"github.com/ozykt4/andrade-api/internal/model"
	"github.com/ozykt4/andrade-api/internal/service"
)

type ProductHandler struct {
	service service.IProductService
}

func NewProductHandler(service service.IProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (ph *ProductHandler) Routes() router.Router {
	return func(route fiber.Router) {
		product := route.Group("/products")
		product.Post("/", ph.CreateProductHandler)
		product.Get("/", ph.FindAllProductsHandler)
		product.Get("/:id", ph.FindProductByIDHandler)
		product.Put("/:id", ph.UpdateProductHandler)
		product.Delete("/:id", ph.DeleteProductHandler)
	}
}

func (ph *ProductHandler) CreateProductHandler(ctx *fiber.Ctx) error {
	productReq := new(model.ProductReq)

	if err := ctx.BodyParser(productReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse(err, fiber.ErrBadRequest))
	}

	res := ph.service.CreateProduct(productReq)
	return ctx.Status(res.Status).JSON(res)
}

func (ph *ProductHandler) FindAllProductsHandler(ctx *fiber.Ctx) error {
	res := ph.service.FindAllProducts()
	return ctx.Status(res.Status).JSON(res)
}

func (ph *ProductHandler) FindProductByIDHandler(ctx *fiber.Ctx) error {
	productID, _ := strconv.Atoi(ctx.Params("id", "0"))

	res := ph.service.FindProductByID(uint(productID))
	return ctx.Status(res.Status).JSON(res)
}

func (ph *ProductHandler) UpdateProductHandler(ctx *fiber.Ctx) error {
	productReq := new(model.ProductReq)
	if err := ctx.BodyParser(productReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse(err, fiber.ErrBadRequest))
	}

	productID, _ := strconv.Atoi(ctx.Params("id", "0"))

	res := ph.service.UpdateProduct(uint(productID), productReq)
	return ctx.Status(res.Status).JSON(res)
}

func (ph *ProductHandler) DeleteProductHandler(ctx *fiber.Ctx) error {
	productID, _ := strconv.Atoi(ctx.Params("id", "0"))

	res := ph.service.DeleteProduct(uint(productID))
	return ctx.Status(res.Status).JSON(res)
}
