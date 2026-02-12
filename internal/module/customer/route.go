package customer

import (
	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, s *Service, v *validator.Validate, l *logger.Logger) {
	h := NewHandler(s, v, l)

	group := r.Group("/customer")
	group.Get("/", h.ListCustomer)
	group.Post("/", h.CreateCustomer)
	group.Put("/:id", h.UpdateCustomer)
	group.Delete("/:id", h.DeleteCustomer)
}
