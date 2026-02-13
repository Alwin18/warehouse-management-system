package warehouse

import (
	"errors"
	"time"

	"github.com/Alwin18/golang-modular-template/internal/shared/constants"
	"github.com/Alwin18/golang-modular-template/internal/shared/db"
	"github.com/Alwin18/golang-modular-template/internal/shared/db/models"
	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Service struct {
	logger logger.Logger
	db     *db.DB
}

func NewService(l logger.Logger, d *db.DB) *Service {
	return &Service{l, d}
}

func (s *Service) ListWarehouse(ctx *fiber.Ctx, params ListWarehouseRequest) ([]ListWarehouseResponse, int64, error) {
	s.logger.Info("List Warehouse")
	result := []ListWarehouseResponse{}

	var warehouses []models.Warehouse
	var count int64

	queryResult := s.db.Gorm.WithContext(ctx.UserContext()).Model(&models.Warehouse{}).Preload("Zones").Preload("Locations")

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
		Find(&warehouses).Error; err != nil {
		return nil, 0, err
	}

	result = make([]ListWarehouseResponse, len(warehouses))
	for i, w := range warehouses {
		result[i] = ListWarehouseResponse{
			ID:            w.ID,
			Code:          w.Code,
			Name:          w.Name,
			Address:       w.Address,
			City:          w.City,
			Country:       w.Country,
			TimeZone:      w.TimeZone,
			IsActive:      w.IsActive,
			TotalZone:     int64(len(w.Zones)),
			TotalLocation: int64(len(w.Locations)),
		}
	}

	return result, count, nil
}

func (s *Service) CreateWarehouse(ctx *fiber.Ctx, body CreateWarehouseRequest) error {
	s.logger.Info("Create Warehouse")
	return s.db.Gorm.WithContext(ctx.UserContext()).Create(&models.Warehouse{
		Code:      body.Code,
		Name:      body.Name,
		Address:   body.Address,
		City:      body.City,
		Country:   body.Country,
		TimeZone:  body.TimeZone,
		IsActive:  body.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error
}

func (s *Service) UpdateWarehouse(ctx *fiber.Ctx, id uint, body CreateWarehouseRequest) error {
	s.logger.Info("Update Warehouse")
	if err := s.db.Gorm.WithContext(ctx.UserContext()).Model(&models.Warehouse{}).Where("id = ?", id).Updates(&models.Warehouse{
		Code:      body.Code,
		Name:      body.Name,
		Address:   body.Address,
		City:      body.City,
		Country:   body.Country,
		TimeZone:  body.TimeZone,
		IsActive:  body.IsActive,
		UpdatedAt: time.Now(),
	}).Error; err != nil {
		s.logger.Error(err.Error())
		return err
	}

	return nil
}

func (s *Service) DeleteWarehouse(ctx *fiber.Ctx, id uint) error {
	s.logger.Info("Delete Warehouse")
	if err := s.db.Gorm.WithContext(ctx.UserContext()).Delete(&models.Warehouse{}, id).Error; err != nil {
		s.logger.Error(err.Error())
		return err
	}

	return nil
}

func (s *Service) CreateWarehouseZone(ctx *fiber.Ctx, body CreateWarehouseZoneRequest) error {
	s.logger.Info("Create Warehouse Zone")

	var wz models.WarehouseZone
	err := s.db.Gorm.WithContext(ctx.UserContext()).Where("code = ?", body.Code).First(&wz).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.Info("Create Warehouse Zone")
		} else {
			s.logger.Error(err.Error())
			return constants.ErrInternalServer
		}
	}

	if wz.ID != 0 {
		s.logger.Error(constants.ErrDataAlreadyExists.Error())
		return constants.ErrDataAlreadyExists
	}

	if err := s.db.Gorm.WithContext(ctx.UserContext()).Create(&models.WarehouseZone{
		Code:        body.Code,
		Name:        body.Name,
		WarehouseID: body.WarehouseID,
		IsActive:    body.IsActive,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}).Error; err != nil {
		s.logger.Error(err.Error())
		return err
	}

	return nil
}

func (s *Service) CreateWarehouseLocation(ctx *fiber.Ctx, body CreateWarehouseLocationRequest) error {
	s.logger.Info("Create Warehouse Location")

	var wl models.Location
	err := s.db.Gorm.WithContext(ctx.UserContext()).Where("code = ?", body.Code).First(&wl).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.Info("Create Warehouse Location")
		} else {
			s.logger.Error(err.Error())
			return constants.ErrInternalServer
		}
	}

	if wl.ID != 0 {
		s.logger.Error(constants.ErrDataAlreadyExists.Error())
		return constants.ErrDataAlreadyExists
	}

	if err := s.db.Gorm.WithContext(ctx.UserContext()).Create(&models.Location{
		WarehouseID:  body.WarehouseID,
		ZoneID:       body.ZoneID,
		Code:         body.Code,
		IsActive:     body.IsActive,
		Description:  body.Description,
		LocationType: body.LocationType,
		MaxVolume:    body.MaxVolume,
		MaxWeight:    body.MaxWeight,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}).Error; err != nil {
		s.logger.Error(err.Error())
		return err
	}

	return nil
}
