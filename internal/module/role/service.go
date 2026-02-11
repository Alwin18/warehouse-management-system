package role

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

func (s *Service) ListRole(ctx *fiber.Ctx, params ListRoleRequest) ([]ListRoleResponse, int64, error) {
	s.logger.Info("List Role")
	result := []ListRoleResponse{}

	var roles []models.Role
	var count int64

	queryResult := s.db.Gorm.WithContext(ctx.UserContext()).Model(&models.Role{})

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
		Find(&roles).Error; err != nil {
		return nil, 0, err
	}

	result = make([]ListRoleResponse, len(roles))
	for i, r := range roles {
		result[i] = ListRoleResponse{
			ID:          r.ID,
			Code:        r.Code,
			Name:        r.Name,
			Description: r.Description,
			CreatedAt:   r.CreatedAt,
			UpdatedAt:   r.UpdatedAt,
		}
	}

	return result, count, nil
}

func (s *Service) CreateRole(ctx *fiber.Ctx, body CreateRoleRequest) error {
	s.logger.Info("Create Role")
	return s.db.Gorm.WithContext(ctx.UserContext()).Create(&models.Role{
		Code:        body.Code,
		Name:        body.Name,
		Description: body.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}).Error
}

func (s *Service) UpdateRole(ctx *fiber.Ctx, id uint, body CreateRoleRequest) error {
	s.logger.Info("Update Role")
	return s.db.Gorm.WithContext(ctx.UserContext()).Model(&models.Role{}).Where("id = ?", id).Updates(&models.Role{
		Code:        body.Code,
		Name:        body.Name,
		Description: body.Description,
		UpdatedAt:   time.Now(),
	}).Error
}

func (s *Service) DeleteRole(ctx *fiber.Ctx, id uint) error {
	s.logger.Info("Delete Role")
	return s.db.Gorm.WithContext(ctx.UserContext()).Delete(&models.Role{}, id).Error
}
