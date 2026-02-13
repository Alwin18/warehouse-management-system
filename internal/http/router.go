package http

import (
	"github.com/Alwin18/golang-modular-template/internal/http/middleware"
	"github.com/Alwin18/golang-modular-template/internal/module/auth"
	purchaseorder "github.com/Alwin18/golang-modular-template/internal/module/inbound/purchase_order"
	"github.com/Alwin18/golang-modular-template/internal/module/master_data/carrier"
	"github.com/Alwin18/golang-modular-template/internal/module/master_data/customer"
	"github.com/Alwin18/golang-modular-template/internal/module/master_data/product"
	"github.com/Alwin18/golang-modular-template/internal/module/master_data/role"
	"github.com/Alwin18/golang-modular-template/internal/module/master_data/supplier"
	"github.com/Alwin18/golang-modular-template/internal/module/master_data/warehouse"
	"github.com/Alwin18/golang-modular-template/internal/module/user"
	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Deps struct {
	Logger    logger.Logger
	Validator *validator.Validate

	// User
	AuthService auth.Service
	UserService *user.Service

	// Master Data
	WarehouseService *warehouse.Service
	RoleService      *role.Service
	ProductService   *product.Service
	SupplierService  *supplier.Service
	CustomerService  *customer.Service
	CarrierService   *carrier.Service

	// Inbound
	PurchaseOrderService *purchaseorder.Service
}

func RegisterRoutes(app *fiber.App, d Deps) {
	api := app.Group("/api/v1")

	api.Use(middleware.AuthMiddleware())

	// User
	user.RegisterRoutes(api, d.UserService, d.Validator, &d.Logger)
	auth.RegisterRoutes(api, d.AuthService, d.Validator, &d.Logger)

	// Master Data
	warehouse.RegisterRoutes(api, d.WarehouseService, d.Validator, &d.Logger)
	role.RegisterRoutes(api, d.RoleService, d.Validator, &d.Logger)
	product.RegisterRoutes(api, d.ProductService, d.Validator, &d.Logger)
	supplier.RegisterRoutes(api, d.SupplierService, d.Validator, &d.Logger)
	customer.RegisterRoutes(api, d.CustomerService, d.Validator, &d.Logger)
	carrier.RegisterRoutes(api, d.CarrierService, d.Validator, &d.Logger)

	// Inbound
	purchaseorder.RegisterRoutes(api, d.PurchaseOrderService, d.Validator, &d.Logger)
}
