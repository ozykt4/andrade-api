package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/ozykt4/andrade-api/config"
	"github.com/ozykt4/andrade-api/config/db"
	"github.com/ozykt4/andrade-api/internal/api/handler"
	"github.com/ozykt4/andrade-api/internal/api/router"
	"github.com/ozykt4/andrade-api/internal/model"
	"github.com/ozykt4/andrade-api/internal/repository"
	"github.com/ozykt4/andrade-api/internal/service"
)

func Run(host, port string) error {
	address := fmt.Sprintf("%s:%s", host, port)
	log.Println("Listen app in port ", address)

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		Prefork:     config.GetConfig().Prefork,
		ProxyHeader: fiber.HeaderXForwardedFor,
	})

	db, err := db.ConnectDB(config.GetConfig().DBAPPURL)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db = db.WithContext(ctx)

	if err := db.AutoMigrate(
		&model.Product{},
		&model.User{},
	); err != nil {
		return err
	}

	// Loads all repositories
	productRepo := repository.NewProductRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Loads all services

	productService := service.NewProductService(productRepo)
	userService := service.NewUserService(userRepo)

	// Loads all handlers
	productHandler := handler.NewProductHandler(productService)
	userHandler := handler.NewUserHandler(userService)

	// // Setup middlewares
	// auth := middleware.NewAuthMiddlware(authConfig,
	// 	"/zwa",
	// 	"/api/v1/status",
	// 	"/api/v1/accounts/signup",
	// 	"/api/v1/accounts/login",
	// 	"/api/v1/accounts/confirmation/send",
	// 	"/api/v1/accounts/confirmation",
	// 	"/api/v1/accounts/recovery/send",
	// 	"/api/v1/accounts/recovery",
	// )

	// middleware.SetupMiddleware(app, auth.Authorize)

	// Setup routes
	router.SetupRouter(app,
		productHandler.Routes(),
		userHandler.Routes(),
	)

	c := make(chan os.Signal, 1)
	errc := make(chan error, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		cancel()
		errc <- app.Shutdown()
	}()

	if err := app.Listen(address); err != nil {
		return err
	}

	err = <-errc

	return err
}
