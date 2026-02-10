package constants

import "errors"

const (
	StatusSuccess = "success"
	StatusError   = "error"
	StatusFail    = "fail"
)

const (
	ErrCodeValidation   = "VALIDATION_ERROR"
	ErrCodeNotFound     = "NOT_FOUND"
	ErrCodeUnauthorized = "UNAUTHORIZED"
)

var (
	ErrDataNotFound      = errors.New("data tidak ditemukan")
	ErrInvalidPassword   = errors.New("password yang anda masukan salah")
	ErrInvalidRequest    = errors.New("request tidak valid")
	ErrDataAlreadyExists = errors.New("data sudah ada")
	ErrInternalServer    = errors.New("terjadi kesalahan sistem")
)
