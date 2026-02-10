package auth

import (
	"context"

	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
)

type Service interface {
	Login(ctx context.Context, body LoginRequest) (LoginResponse, error)
}

type service struct {
	logger logger.Logger
}

func NewService(l logger.Logger) Service {
	return &service{l}
}

func (s *service) Login(ctx context.Context, body LoginRequest) (LoginResponse, error) {
	s.logger.Info("login user")
	return LoginResponse{
		User:  UserLogin{},
		Token: "AAAA",
	}, nil
}
