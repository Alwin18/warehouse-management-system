package user

import "github.com/Alwin18/golang-modular-template/internal/shared/logger"

type Service interface {
	GetUsers() ([]User, error)
}

type service struct {
	logger logger.Logger
}

func NewService(l logger.Logger) Service {
	return &service{l}
}

func (s *service) GetUsers() ([]User, error) {
	s.logger.Info("fetch users")
	return nil, nil
}
