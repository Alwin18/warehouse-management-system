package user

import (
	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, s *Service, v *validator.Validate, l *logger.Logger) {
	h := NewHandler(s, v, l)

	group := r.Group("/user")
	group.Get("/", h.ListUsers)
	group.Post("/", h.CreateUser)
	group.Put("/:id", h.UpdateUser)
	group.Delete("/:id", h.DeleteUser)
}
