package auth

import (
	"errors"
	"os"
	"time"

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
	logger logger.Logger
	db     *db.DB
}

func NewService(l logger.Logger, d *db.DB) Service {
	return &service{l, d}
}

func (s *service) Login(ctx *fiber.Ctx, body LoginRequest) (LoginResponse, error) {
	s.logger.Info("login user")
	var result LoginResponse
	var user models.User

	if err := s.db.Gorm.WithContext(ctx.UserContext()).Where("username = ?", body.Username).
		Preload("Roles").
		Where("is_active = ?", true).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.Error("akun tidak ditemukan")
			ctx.Status(fiber.StatusUnauthorized)
			return result, errors.New("akun tidak ditemukan")
		}

		s.logger.Error(err.Error())
		ctx.Status(fiber.StatusInternalServerError)
		return result, errors.New("terjadi kesalahan sistem")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		s.logger.Error(err.Error())
		ctx.Status(fiber.StatusUnauthorized)
		return result, errors.New("password yang anda masukan salah")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"username":  user.Username,
		"full_name": user.FullName,
		"email":     user.Email,
		"exp":       time.Now().Add(time.Hour * 3).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return LoginResponse{}, err
	}

	role := []UserRole{}
	for _, userRole := range user.Roles {
		role = append(role, UserRole{
			ID:   userRole.ID,
			Name: userRole.Name,
			Code: userRole.Code,
		})
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
			Roles:     role,
		},
		Token: tokenString,
	}, nil
}
