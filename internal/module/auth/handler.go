package auth

import "github.com/gofiber/fiber/v2"

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s}
}

func (h *Handler) Login(ctx *fiber.Ctx) error {

	var body LoginRequest
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	loginResponse, err := h.service.Login(ctx.UserContext(), body)
	if err != nil {
		return err
	}
	return ctx.JSON(loginResponse)
}
