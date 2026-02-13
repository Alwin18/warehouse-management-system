package purchaseorder

import (
	"github.com/Alwin18/golang-modular-template/internal/shared/errors"
	"github.com/Alwin18/golang-modular-template/internal/shared/formatting"
	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
	"github.com/Alwin18/golang-modular-template/internal/shared/response"
	"github.com/Alwin18/golang-modular-template/internal/shared/validation"
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

func (h *Handler) ListPurchaseOrder(ctx *fiber.Ctx) error {
	var params ListPurchaseOrderRequest
	if err := ctx.QueryParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.NewErrorResponse(response.ResponseError{
			Message: err.Error(),
			Code:    fiber.StatusBadRequest,
		}))
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.PerPage == 0 {
		params.PerPage = 10
	}

	if err := h.validate.Struct(params); err != nil {
		validationErrors := validation.FormatValidationErrors(err, params)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.NewErrorResponse(response.ResponseError{
			Message: "body request tidak sesuai",
			Code:    fiber.StatusBadRequest,
			Errors:  validationErrors,
		}))
	}

	resp, count, err := h.service.ListPurchaseOrder(ctx, params)
	if err != nil {
		return errors.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(response.NewResponseWithPagination(resp, "success", fiber.StatusOK, response.Meta{
		Page:      params.Page,
		PerPage:   params.PerPage,
		TotalData: count,
		Totalpage: int64(response.TotalPage(count, params.PerPage)),
	}))
}

func (h *Handler) CreatePurchaseOrder(ctx *fiber.Ctx) error {
	var body CreatePurchaseOrderRequest
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.NewErrorResponse(response.ResponseError{
			Message: err.Error(),
			Code:    fiber.StatusBadRequest,
		}))
	}

	if err := h.validate.Struct(body); err != nil {
		validationErrors := validation.FormatValidationErrors(err, body)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.NewErrorResponse(response.ResponseError{
			Message: "body request tidak sesuai",
			Code:    fiber.StatusBadRequest,
			Errors:  validationErrors,
		}))
	}

	err := h.service.CreatePurchaseOrder(ctx, body)
	if err != nil {
		return errors.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.NewProcessResponse("success", fiber.StatusCreated))
}

func (h *Handler) UpdatePurchaseOrder(ctx *fiber.Ctx) error {
	id := formatting.ConvertStringIntoUint(ctx.Params("id"))
	if id == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.NewErrorResponse(response.ResponseError{
			Message: "Invalid id",
			Code:    fiber.StatusBadRequest,
		}))
	}

	var body CreatePurchaseOrderRequest
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.NewErrorResponse(response.ResponseError{
			Message: err.Error(),
			Code:    fiber.StatusBadRequest,
		}))
	}

	if err := h.validate.Struct(body); err != nil {
		validationErrors := validation.FormatValidationErrors(err, body)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.NewErrorResponse(response.ResponseError{
			Message: "body request tidak sesuai",
			Code:    fiber.StatusBadRequest,
			Errors:  validationErrors,
		}))
	}

	err := h.service.UpdatePurchaseOrder(ctx, id, body)
	if err != nil {
		return errors.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(response.NewProcessResponse("success", fiber.StatusOK))
}

func (h *Handler) DeletePurchaseOrder(ctx *fiber.Ctx) error {
	id := formatting.ConvertStringIntoUint(ctx.Params("id"))
	if id == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.NewErrorResponse(response.ResponseError{
			Message: "Invalid id",
			Code:    fiber.StatusBadRequest,
		}))
	}

	err := h.service.DeletePurchaseOrder(ctx, id)
	if err != nil {
		return errors.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(response.NewProcessResponse("success", fiber.StatusOK))
}
