package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ozykt4/andrade-api/internal/api/router"
	"github.com/ozykt4/andrade-api/internal/model"
	"github.com/ozykt4/andrade-api/internal/service"
)

type UserHandler struct {
	service service.IUserService
}

func NewUserHandler(service service.IUserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) Routes() router.Router {
	return func(router fiber.Router) {
		user := router.Group("users")
		user.Post("/", uh.CreateUserHandler)
	}
}

func (uh *UserHandler) CreateUserHandler(ctx *fiber.Ctx) error {
	userReq := new(model.UserReq)

	if err := ctx.BodyParser(userReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse(err, fiber.ErrBadRequest))
	}

	res := uh.service.CreateUser(userReq)
	return ctx.Status(res.Status).JSON(res)
}
