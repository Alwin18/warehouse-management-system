package auth

import (
	"errors"
	"time"

	"github.com/Alwin18/golang-modular-template/internal/shared/constants"
	"github.com/Alwin18/golang-modular-template/internal/shared/db"
	"github.com/Alwin18/golang-modular-template/internal/shared/db/models"
	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Login(ctx *fiber.Ctx, body LoginRequest) (LoginResponse, error)
}

type service struct {
	logger     logger.Logger
	db         *db.DB
	jwtSecret  string
	jwtExpired string
}

func NewService(l logger.Logger, d *db.DB, jwtSecret, jwtExpired string) Service {
	return &service{l, d, jwtSecret, jwtExpired}
}

func (s *service) Login(ctx *fiber.Ctx, body LoginRequest) (LoginResponse, error) {
	s.logger.Info("login user")
	var result LoginResponse
	var user models.User

	if err := s.db.Gorm.WithContext(ctx.UserContext()).
		Preload("Roles").
		Where("username = ? AND is_active = ?", body.Username, true).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.Error(constants.ErrDataNotFound.Error())
			return result, constants.ErrDataNotFound
		}

		s.logger.Error(err.Error())
		return result, constants.ErrInternalServer
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		s.logger.Error("invalid password attempt for username: " + body.Username)
		return result, constants.ErrInvalidPassword
	}

	expDuration, err := time.ParseDuration(s.jwtExpired)
	if err != nil {
		s.logger.Error("invalid JWT expiration format: " + err.Error())
		expDuration = time.Hour * 3
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(expDuration).Unix(),
		"iss":     "warehouse-management-system",
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		s.logger.Error(err.Error())
		return result, constants.ErrInternalServer
	}

	roles := make([]UserRole, len(user.Roles))
	for i, userRole := range user.Roles {
		roles[i] = UserRole{
			ID:   userRole.ID,
			Name: userRole.Name,
			Code: userRole.Code,
		}
	}

	return LoginResponse{
		User: UserLogin{
			ID:        user.ID,
			Username:  user.Username,
			FullName:  user.FullName,
			Email:     user.Email,
			Phone:     user.Phone,
			IsActive:  user.IsActive,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Roles:     roles,
		},
		Token: tokenString,
	}, nil
}
