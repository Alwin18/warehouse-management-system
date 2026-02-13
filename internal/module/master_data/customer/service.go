package customer

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

// List
func (s *Service) ListCustomer(ctx *fiber.Ctx, params ListCustomerRequest) ([]ListCustomerResponse, int64, error) {
	var customers []models.Customer
	var count int64

	query := s.db.Gorm.WithContext(ctx.UserContext()).Model(&models.Customer{})

	if err := query.Count(&count).Error; err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}

	if err := query.Limit(params.PerPage).
		Offset((params.Page - 1) * params.PerPage).
		Order("updated_at desc").
		Find(&customers).Error; err != nil {
		s.logger.Error(err.Error())
		return nil, 0, err
	}

	result := make([]ListCustomerResponse, len(customers))
	for i, customer := range customers {
		result[i] = ListCustomerResponse{
			ID:        customer.ID,
			Code:      customer.Code,
			Name:      customer.Name,
			Address:   customer.Address,
			City:      customer.City,
			Country:   customer.Country,
			Phone:     customer.Phone,
			Email:     customer.Email,
			IsActive:  customer.IsActive,
			CreatedAt: customer.CreatedAt,
			UpdatedAt: customer.UpdatedAt,
		}
	}

	return result, count, nil
}

// Create
func (s *Service) CreateCustomer(ctx *fiber.Ctx, body CreateCustomerRequest) error {
	if err := s.db.Gorm.WithContext(ctx.UserContext()).Create(&models.Customer{
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

// Update
func (s *Service) UpdateCustomer(ctx *fiber.Ctx, id uint, body CreateCustomerRequest) error {
	if err := s.db.Gorm.WithContext(ctx.UserContext()).Where("id = ?", id).Updates(&models.Customer{
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

// Delete
func (s *Service) DeleteCustomer(ctx *fiber.Ctx, id uint) error {
	if err := s.db.Gorm.WithContext(ctx.UserContext()).Where("id = ?", id).Delete(&models.Customer{}).Error; err != nil {
		s.logger.Error(err.Error())
		return err
	}

	return nil
}
