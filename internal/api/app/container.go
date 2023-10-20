package app

import (
	"database/sql"
	"frame/internal/config"
	"frame/internal/database"
	"frame/internal/lib/logger"
	"frame/internal/repository"
	userRepository "frame/internal/repository/user"
	userProfileRepository "frame/internal/repository/user_profile"
	"frame/internal/service"
	userService "frame/internal/service/user"
	"log/slog"
)

type container struct {
	config                *config.Config
	logger                *slog.Logger
	db                    *sql.DB
	transactionManager    database.TransactionManager
	userRepository        repository.UserRepository
	userProfileRepository repository.UserProfileRepository
	userService           service.UserService
}

func newContainer() *container {
	return &container{}
}

func (c *container) Config() *config.Config {
	if c.config == nil {
		c.config = config.Parse()
	}

	return c.config
}

func (c *container) Logger() *slog.Logger {
	if c.logger == nil {
		c.logger = logger.NewLogger(c.Config())
	}

	return c.logger
}

func (c *container) DB() *sql.DB {
	if c.db == nil {
		c.db = database.NewPostgres(c.Config())
		c.Logger().Info("database connection pull established")
	}

	return c.db
}

func (c *container) TransactionManager() database.TransactionManager {
	if c.transactionManager == nil {
		c.transactionManager = database.NewTransactionManager(c.DB())
	}

	return c.transactionManager
}

func (c *container) UserRepository() repository.UserRepository {
	if c.userRepository == nil {
		c.userRepository = userRepository.NewRepository(c.DB())
	}

	return c.userRepository
}

func (c *container) UserProfileRepository() repository.UserProfileRepository {
	if c.userProfileRepository == nil {
		c.userProfileRepository = userProfileRepository.NewRepository(c.DB())
	}

	return c.userProfileRepository
}

func (c *container) UserService() service.UserService {
	if c.userService == nil {
		c.userService = userService.NewService(
			c.UserRepository(),
			c.UserProfileRepository(),
			c.TransactionManager(),
		)
	}

	return c.userService
}
