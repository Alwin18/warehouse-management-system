package supplier

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

func (s *Service) ListSupplier(ctx *fiber.Ctx, params ListSupplierRequest) ([]ListSupplierResponse, int64, error) {
	var supplier []models.Supplier
	var count int64

	query := s.db.Gorm.WithContext(ctx.UserContext()).Model(&models.Supplier{})

	if err := query.Count(&count).Error; err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}

	if err := query.Limit(params.PerPage).
		Offset((params.Page - 1) * params.PerPage).
		Order("updated_at desc").
		Find(&supplier).Error; err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}

	result := make([]ListSupplierResponse, len(supplier))

	for i, v := range supplier {
		result[i] = ListSupplierResponse{
			ID:        v.ID,
			Code:      v.Code,
			Name:      v.Name,
			Address:   v.Address,
			City:      v.City,
			Country:   v.Country,
			Phone:     v.Phone,
			Email:     v.Email,
			IsActive:  v.IsActive,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return result, count, nil
}

func (s *Service) CreateSupplier(ctx *fiber.Ctx, body CreateSupplierRequest) error {
	if err := s.db.Gorm.WithContext(ctx.UserContext()).Create(&models.Supplier{
		Code:      body.Code,
		Name:      body.Name,
		Address:   body.Address,
		City:      body.City,
		Country:   body.Country,
		Phone:     body.Phone,
		Email:     body.Email,
		IsActive:  body.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error; err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Service) UpdateSupplier(ctx *fiber.Ctx, id uint, body CreateSupplierRequest) error {
	if err := s.db.Gorm.WithContext(ctx.UserContext()).Where("id = ?", id).Updates(&models.Supplier{
		Code:      body.Code,
		Name:      body.Name,
		Address:   body.Address,
		City:      body.City,
		Country:   body.Country,
		Phone:     body.Phone,
		Email:     body.Email,
		IsActive:  body.IsActive,
		UpdatedAt: time.Now(),
	}).Error; err != nil {
		s.logger.Error(err.Error())
		return err
	}

	return nil
}

func (s *Service) DeleteSupplier(ctx *fiber.Ctx, id uint) error {
	if err := s.db.Gorm.WithContext(ctx.UserContext()).Where("id = ?", id).Delete(&models.Supplier{}).Error; err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}
