package warehouse

import (
	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, s *Service, v *validator.Validate, l *logger.Logger) {
	h := NewHandler(s, v, l)

	group := r.Group("/warehouse")
	group.Post("/", h.CreateWarehouse)
	group.Get("/", h.ListWarehouse)
	group.Patch("/:id", h.UpdateWarehouse)
	group.Delete("/:id", h.DeleteWarehouse)
	group.Post("/zone", h.CreateWarehouseZone)
	group.Post("/location", h.CreateWarehouseLocation)
}
