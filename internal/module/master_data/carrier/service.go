package carrier

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

func (s *Service) ListCarriers(ctx *fiber.Ctx, params ListCarrierRequest) ([]ListCarrierResponse, int64, error) {
	s.logger.Info("List Carrier")
	result := []ListCarrierResponse{}

	var carriers []models.Carrier
	var count int64

	queryResult := s.db.Gorm.WithContext(ctx.UserContext()).Model(&models.Carrier{})

	if params.IsActive != nil {
		queryResult = queryResult.Where("is_active = ?", params.IsActive)
	}

	if search := params.Search; search != "" {
		like := "%" + search + "%"
		queryResult = queryResult.Where("code ILIKE ? OR name ILIKE ?", like, like)
	}

	if err := queryResult.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := queryResult.
		Limit(params.PerPage).
		Offset((params.Page - 1) * params.PerPage).
		Order("updated_at desc").
		Find(&carriers).Error; err != nil {
		return nil, 0, err
	}

	result = make([]ListCarrierResponse, len(carriers))
	for i, w := range carriers {
		result[i] = ListCarrierResponse{
			ID:        w.ID,
			Code:      w.Code,
			Name:      w.Name,
			IsActive:  w.IsActive,
			CreatedAt: w.CreatedAt,
			UpdatedAt: w.UpdatedAt,
		}
	}

	return result, count, nil
}

func (s *Service) CreateCarrier(ctx *fiber.Ctx, body CreateCarrierRequest) error {
	s.logger.Info("Create Carrier")
	return s.db.Gorm.WithContext(ctx.UserContext()).Create(&models.Carrier{
		Code:      body.Code,
		Name:      body.Name,
		IsActive:  body.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error
}

func (s *Service) UpdateCarrier(ctx *fiber.Ctx, id uint, body CreateCarrierRequest) error {
	s.logger.Info("Update Carrier")
	if err := s.db.Gorm.WithContext(ctx.UserContext()).Model(&models.Carrier{}).Where("id = ?", id).Updates(&models.Carrier{
		Code:      body.Code,
		Name:      body.Name,
		IsActive:  body.IsActive,
		UpdatedAt: time.Now(),
	}).Error; err != nil {
		s.logger.Error(err.Error())
		return err
	}

	return nil
}

func (s *Service) DeleteCarrier(ctx *fiber.Ctx, id uint) error {
	s.logger.Info("Delete Carrier")
	if err := s.db.Gorm.WithContext(ctx.UserContext()).Delete(&models.Carrier{}, id).Error; err != nil {
		s.logger.Error(err.Error())
		return err
	}

	return nil
}
