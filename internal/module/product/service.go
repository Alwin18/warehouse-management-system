package product

import (
	"time"

	"github.com/Alwin18/golang-modular-template/internal/shared/db"
	"github.com/Alwin18/golang-modular-template/internal/shared/db/models"
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

func (s *Service) ListProduct(ctx *fiber.Ctx, params ListProductRequest) ([]ListProductResponse, int64, error) {
	result := []ListProductResponse{}
	var count int64

	queryResult := s.db.Gorm.WithContext(ctx.UserContext()).Model(&models.Product{})

	if err := queryResult.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := queryResult.
		Limit(params.PerPage).
		Offset((params.Page - 1) * params.PerPage).
		Order("updated_at desc").
		Find(&result).Error; err != nil {
		return nil, 0, err
	}

	return result, count, nil
}

func (s *Service) CreateProduct(ctx *fiber.Ctx, req CreateProductRequest) error {

	if err := s.db.Gorm.WithContext(ctx.UserContext()).Create(&models.Product{
		SKU:            req.SKU,
		Name:           req.Name,
		Barcode:        req.Barcode,
		Description:    req.Description,
		Weight:         req.Weight,
		Volume:         req.Volume,
		IsBatchManaged: req.IsBatchManaged,
		IsSerialized:   req.IsSerialized,
		IsActive:       req.IsActive,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}).Error; err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateProduct(ctx *fiber.Ctx, id uint, req CreateProductRequest) error {
	if err := s.db.Gorm.WithContext(ctx.UserContext()).Model(&models.Product{}).Where("id = ?", id).Updates(&models.Product{
		SKU:            req.SKU,
		Name:           req.Name,
		Barcode:        req.Barcode,
		Description:    req.Description,
		Weight:         req.Weight,
		Volume:         req.Volume,
		IsBatchManaged: req.IsBatchManaged,
		IsSerialized:   req.IsSerialized,
		IsActive:       req.IsActive,
		UpdatedAt:      time.Now(),
	}).Error; err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteProduct(ctx *fiber.Ctx, id uint) error {
	if err := s.db.Gorm.WithContext(ctx.UserContext()).Delete(&models.Product{}, id).Error; err != nil {
		return err
	}

	return nil
}
