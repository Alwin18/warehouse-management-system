package http

import (
	"github.com/Alwin18/golang-modular-template/internal/module/auth"
	"github.com/Alwin18/golang-modular-template/internal/module/carrier"
	"github.com/Alwin18/golang-modular-template/internal/module/customer"
	"github.com/Alwin18/golang-modular-template/internal/module/product"
	"github.com/Alwin18/golang-modular-template/internal/module/role"
	"github.com/Alwin18/golang-modular-template/internal/module/supplier"
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
	AuthService      auth.Service
	UserService      *user.Service
	WarehouseService *warehouse.Service
	RoleService      *role.Service
	ProductService   *product.Service
	SupplierService  *supplier.Service
	CustomerService  *customer.Service
	CarrierService   *carrier.Service
}

func RegisterRoutes(app *fiber.App, d Deps) {
	api := app.Group("/api/v1")

	user.RegisterRoutes(api, d.UserService, d.Validator, &d.Logger)
	auth.RegisterRoutes(api, d.AuthService, d.Validator, &d.Logger)
	warehouse.RegisterRoutes(api, d.WarehouseService, d.Validator, &d.Logger)
	role.RegisterRoutes(api, d.RoleService, d.Validator, &d.Logger)
	product.RegisterRoutes(api, d.ProductService, d.Validator, &d.Logger)
	supplier.RegisterRoutes(api, d.SupplierService, d.Validator, &d.Logger)
	customer.RegisterRoutes(api, d.CustomerService, d.Validator, &d.Logger)
	carrier.RegisterRoutes(api, d.CarrierService, d.Validator, &d.Logger)
}
