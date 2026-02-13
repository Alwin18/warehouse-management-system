package purchaseorder

import (
	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, s *Service, v *validator.Validate, l *logger.Logger) {
	h := NewHandler(s, v, l)

	group := r.Group("/purchase-order")
	group.Get("/", h.ListPurchaseOrder)
	group.Post("/", h.CreatePurchaseOrder)
	group.Patch("/:id", h.UpdatePurchaseOrder)
	group.Delete("/:id", h.DeletePurchaseOrder)
}
