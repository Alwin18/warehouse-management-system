package app

import (
	"github.com/Alwin18/golang-modular-template/config"
	"github.com/Alwin18/golang-modular-template/internal/module/auth"
	"github.com/Alwin18/golang-modular-template/internal/module/user"
	"github.com/Alwin18/golang-modular-template/internal/module/warehouse"
	"github.com/Alwin18/golang-modular-template/internal/shared/db"
	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
	"github.com/Alwin18/golang-modular-template/internal/shared/redis"
	"github.com/Alwin18/golang-modular-template/internal/shared/validation"
	"github.com/go-playground/validator/v10"
)

type Container struct {
	DB        *db.DB
	Redis     *redis.Client
	Logger    logger.Logger
	Validator *validator.Validate

	UserService      user.Service
	AuthService      auth.Service
	WarehouseService *warehouse.Service
}

func NewContainer(cfg *config.Config) (*Container, error) {
	log := logger.New()

	database, err := db.NewPostgres(cfg, log)
	if err != nil {
		return nil, err
	}

	validator := validation.NewValidator()

	// Run auto migration
	// if err := db.AutoMigrate(database.Gorm, log); err != nil {
	// 	return nil, err
	// }

	redisClient := redis.New()

	userService := user.NewService(log)
	authService := auth.NewService(log, database, cfg.JWTSecret, cfg.JWTExpired)
	warehouseService := warehouse.NewService(log, database)

	return &Container{
		DB:        database,
		Redis:     redisClient,
		Logger:    log,
		Validator: validator,

		// Module Services
		UserService:      userService,
		AuthService:      authService,
		WarehouseService: warehouseService,
	}, nil
}

// Cleanup closes all resources (database, redis, etc.)
func (c *Container) Cleanup() {
	c.Logger.Info("Starting cleanup...")

	// Close database connection
	if err := c.DB.Close(c.Logger); err != nil {
		c.Logger.Error("Error closing database:", err)
	}

	// Close redis connection
	if err := c.Redis.Close(); err != nil {
		c.Logger.Error("Error closing redis:", err)
	} else {
		c.Logger.Info("Redis connection closed successfully")
	}

	c.Logger.Info("Cleanup completed")
}
