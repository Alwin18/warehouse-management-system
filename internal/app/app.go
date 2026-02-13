package app

import (
	httpDelivery "github.com/Alwin18/golang-modular-template/internal/http"
	"github.com/gofiber/fiber/v2"
)

func NewApp(c *Container) *fiber.App {
	app := fiber.New()

	httpDelivery.RegisterRoutes(app, httpDelivery.Deps{
		Logger:    c.Logger,
		Validator: c.Validator,

		// Module Services
		UserService:          c.UserService,
		AuthService:          c.AuthService,
		WarehouseService:     c.WarehouseService,
		RoleService:          c.RoleService,
		ProductService:       c.ProductService,
		SupplierService:      c.SupplierService,
		CustomerService:      c.CustomerService,
		CarrierService:       c.CarrierService,
		PurchaseOrderService: c.PurchaseOrderService,
	})

	return app
}
