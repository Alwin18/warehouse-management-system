package user

import (
	"time"

	"github.com/Alwin18/golang-modular-template/internal/shared/db"
	"github.com/Alwin18/golang-modular-template/internal/shared/db/models"
	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	logger logger.Logger
	db     *db.DB
}

func NewService(l logger.Logger, d *db.DB) *Service {
	return &Service{l, d}
}

func (s *Service) ListUsers(ctx *fiber.Ctx, params ListUserRequest) ([]ListUserResponse, int64, error) {
	var users []models.User
	var count int64

	query := s.db.Gorm.WithContext(ctx.UserContext()).Model(&models.User{})

	if err := query.Count(&count).Error; err != nil {
		s.logger.Error("failed to count users", err)
		return nil, 0, err
	}

	if err := query.Offset((params.Page - 1) * params.PerPage).
		Limit(params.PerPage).
		Order("updated_at desc").
		Find(&users).Error; err != nil {
		s.logger.Error("failed to find users", err)
		return nil, 0, err
	}

	result := make([]ListUserResponse, len(users))
	for i, v := range users {
		result[i] = ListUserResponse{
			ID:        v.ID,
			Username:  v.Username,
			Email:     v.Email,
			FullName:  v.FullName,
			Phone:     v.Phone,
			IsActive:  v.IsActive,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return result, count, nil
}

func (s *Service) CreateUser(ctx *fiber.Ctx, body CreateUserRequest) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("failed to generate password", err)
		return err
	}

	if err := s.db.Gorm.WithContext(ctx.UserContext()).Create(&models.User{
		Username:  body.Username,
		Email:     body.Email,
		FullName:  body.FullName,
		Password:  string(hashPassword),
		Phone:     body.Phone,
		IsActive:  body.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error; err != nil {
		s.logger.Error("failed to create user", err)
		return err
	}
	return nil
}

func (s *Service) UpdateUser(ctx *fiber.Ctx, id uint, body CreateUserRequest) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("failed to generate password", err)
		return err
	}

	if err := s.db.Gorm.WithContext(ctx.UserContext()).Model(&models.User{}).Where("id = ?", id).Updates(&models.User{
		Username:  body.Username,
		Email:     body.Email,
		FullName:  body.FullName,
		Phone:     body.Phone,
		Password:  string(hashPassword),
		IsActive:  body.IsActive,
		UpdatedAt: time.Now(),
	}).Error; err != nil {
		s.logger.Error("failed to update user", err)
		return err
	}
	return nil
}

func (s *Service) DeleteUser(ctx *fiber.Ctx, id uint) error {
	if err := s.db.Gorm.WithContext(ctx.UserContext()).Model(&models.User{}).Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		s.logger.Error("failed to delete user", err)
		return err
	}
	return nil
}
