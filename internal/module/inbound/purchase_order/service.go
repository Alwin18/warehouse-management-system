package purchaseorder

import (
	"math/rand"
	"time"

	"github.com/Alwin18/golang-modular-template/internal/shared/db"
	"github.com/Alwin18/golang-modular-template/internal/shared/db/models"
	"github.com/Alwin18/golang-modular-template/internal/shared/formatting"
	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
	"github.com/gofiber/fiber/v2"
)

type Service struct {
	logger logger.Logger
	db     *db.DB
}

func NewService(l logger.Logger, d *db.DB) *Service {
	return &Service{l, d}
}

func (s *Service) ListPurchaseOrder(ctx *fiber.Ctx, params ListPurchaseOrderRequest) ([]ListPurchaseOrderResponse, int64, error) {
	var users []models.PurchaseOrder
	var count int64

	query := s.db.Gorm.WithContext(ctx.UserContext()).
		Model(&models.PurchaseOrder{}).
		Preload("User")

	if err := query.Count(&count).Error; err != nil {
		s.logger.Error("failed to count purchase orders", err)
		return nil, 0, err
	}

	if err := query.Offset((params.Page - 1) * params.PerPage).
		Limit(params.PerPage).
		Order("updated_at desc").
		Find(&users).Error; err != nil {
		s.logger.Error("failed to find purchase orders", err)
		return nil, 0, err
	}

	result := make([]ListPurchaseOrderResponse, len(users))
	for i, v := range users {
		result[i] = ListPurchaseOrderResponse{
			ID:           v.ID,
			PONumber:     v.PONumber,
			SupplierID:   v.SupplierID,
			WarehouseID:  v.WarehouseID,
			Status:       v.Status,
			OrderDate:    v.OrderDate,
			ExpectedDate: v.ExpectedDate,
			Currency:     v.Currency,
			TotalAmount:  v.TotalAmount,
			CreatedBy:    v.User.FullName,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		}
	}

	return result, count, nil
}

func (s *Service) CreatePurchaseOrder(ctx *fiber.Ctx, body CreatePurchaseOrderRequest) error {
	rand := rand.Uint64()
	if err := s.db.Gorm.WithContext(ctx.UserContext()).Create(&models.PurchaseOrder{
		PONumber:     formatting.GeneratePONumber(int64(rand)),
		SupplierID:   body.SupplierID,
		WarehouseID:  body.WarehouseID,
		Status:       body.Status,
		OrderDate:    body.OrderDate,
		ExpectedDate: body.ExpectedDate,
		Currency:     body.Currency,
		TotalAmount:  body.TotalAmount,
		CreatedBy:    1,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}).Error; err != nil {
		s.logger.Error("failed to create purchase order", err)
		return err
	}
	return nil
}

func (s *Service) UpdatePurchaseOrder(ctx *fiber.Ctx, id uint, body CreatePurchaseOrderRequest) error {
	if err := s.db.Gorm.WithContext(ctx.UserContext()).Model(&models.PurchaseOrder{}).Where("id = ?", id).Updates(&models.PurchaseOrder{
		SupplierID:   body.SupplierID,
		WarehouseID:  body.WarehouseID,
		Status:       body.Status,
		OrderDate:    body.OrderDate,
		ExpectedDate: body.ExpectedDate,
		Currency:     body.Currency,
		TotalAmount:  body.TotalAmount,
		UpdatedAt:    time.Now(),
	}).Error; err != nil {
		s.logger.Error("failed to update purchase order", err)
		return err
	}
	return nil
}

func (s *Service) DeletePurchaseOrder(ctx *fiber.Ctx, id uint) error {
	if err := s.db.Gorm.WithContext(ctx.UserContext()).Model(&models.PurchaseOrder{}).Where("id = ?", id).Delete(&models.PurchaseOrder{}).Error; err != nil {
		s.logger.Error("failed to delete purchase order", err)
		return err
	}
	return nil
}
