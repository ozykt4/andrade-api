package handler

import (
	"strconv"

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
		user.Get("/", uh.FindAllUsersHandler)
		user.Get("/:id", uh.FindUserByID)
		user.Put("/:id", uh.UpdateUserHandler)
		user.Delete("/:id", uh.DeleteUserHandler)
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

func (uh *UserHandler) FindAllUsersHandler(ctx *fiber.Ctx) error {
	res := uh.service.FindAllUsers()
	return ctx.Status(res.Status).JSON(res)
}

func (uh *UserHandler) FindUserByID(ctx *fiber.Ctx) error {
	userID, _ := strconv.Atoi(ctx.Params("id", "0"))

	res := uh.service.FindUserByID(uint(userID))

	return ctx.Status(res.Status).JSON(res)
}

func (uh *UserHandler) UpdateUserHandler(ctx *fiber.Ctx) error {
	userReq := new(model.UserReq)
	if err := ctx.BodyParser(userReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse(err, fiber.ErrBadRequest))
	}

	userID, _ := strconv.Atoi(ctx.Params("id", "0"))

	res := uh.service.UpdateUser(uint(userID), userReq)
	return ctx.Status(res.Status).JSON(res)
}

func (uh *UserHandler) DeleteUserHandler(ctx *fiber.Ctx) error {
	userID, _ := strconv.Atoi(ctx.Params("id", "0"))

	res := uh.service.DeleteUser(uint(userID))
	return ctx.Status(res.Status).JSON(res)
}
