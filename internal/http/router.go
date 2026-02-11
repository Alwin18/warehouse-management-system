package http

import (
	"github.com/Alwin18/golang-modular-template/internal/module/auth"
	"github.com/Alwin18/golang-modular-template/internal/module/product"
	"github.com/Alwin18/golang-modular-template/internal/module/role"
	"github.com/Alwin18/golang-modular-template/internal/module/user"
	"github.com/Alwin18/golang-modular-template/internal/module/warehouse"
	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Deps struct {
	Logger    logger.Logger
	Validator *validator.Validate

	// Module Services
	UserService      user.Service
	AuthService      auth.Service
	WarehouseService *warehouse.Service
	RoleService      *role.Service
	ProductService   *product.Service
}

func RegisterRoutes(app *fiber.App, d Deps) {
	api := app.Group("/api/v1")

	user.RegisterRoutes(api, d.UserService)
	auth.RegisterRoutes(api, d.AuthService, d.Validator, &d.Logger)
	warehouse.RegisterRoutes(api, d.WarehouseService, d.Validator, &d.Logger)
	role.RegisterRoutes(api, d.RoleService, d.Validator, &d.Logger)
	product.RegisterRoutes(api, d.ProductService, d.Validator, &d.Logger)
}
