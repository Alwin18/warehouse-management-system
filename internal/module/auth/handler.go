package auth

import (
	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
	"github.com/Alwin18/golang-modular-template/internal/shared/response"
	"github.com/Alwin18/golang-modular-template/internal/shared/validation"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	validate *validator.Validate
	logger   *logger.Logger
	service  Service
}

func NewHandler(s Service, v *validator.Validate, l *logger.Logger) *Handler {
	return &Handler{
		service:  s,
		validate: v,
		logger:   l,
	}
}

func (h *Handler) Login(ctx *fiber.Ctx) error {
	var body LoginRequest
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.NewErrorResponse(response.ResponseError{
			Message: err.Error(),
			Code:    fiber.StatusBadRequest,
		}))
	}

	if err := h.validate.Struct(body); err != nil {
		errors := validation.FormatValidationErrors(err, body)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.NewErrorResponse(response.ResponseError{
			Message: "body request tidak sesuai",
			Code:    fiber.StatusBadRequest,
			Errors:  errors,
		}))
	}

	resp, err := h.service.Login(ctx, body)
	if err != nil {
		return ctx.JSON(response.NewErrorResponse(response.ResponseError{
			Message: err.Error(),
			Code:    ctx.Response().Header.StatusCode(),
		}))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.NewResponse(resp, "success login", fiber.StatusOK))
}
