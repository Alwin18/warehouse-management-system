package examplemodule

import (
	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	validate *validator.Validate
	logger   *logger.Logger
	service  *Service
}

func NewHandler(s *Service, v *validator.Validate, l *logger.Logger) *Handler {
	return &Handler{
		service:  s,
		validate: v,
		logger:   l,
	}
}

func (h *Handler) ListExampleModule(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}
