package app

import (
	"database/sql"
	"frame/internal/config"
	"frame/internal/database"
	"frame/internal/lib/logger"
	"frame/internal/lib/session"
	"frame/internal/repository"
	userRepository "frame/internal/repository/user"
	userProfileRepository "frame/internal/repository/user_profile"
	"frame/internal/service"
	userService "frame/internal/service/user"
	"log"
	"log/slog"
)

type container struct {
	projectConfig         *config.ProjectConfig
	apiServerConfig       *config.APIServerConfig
	dbConfig              *config.DBConfig
	logger                *slog.Logger
	session               *session.Session
	db                    *sql.DB
	transactionManager    database.TransactionManager
	userRepository        repository.UserRepository
	userProfileRepository repository.UserProfileRepository
	userService           service.UserService
}

func newContainer() *container {
	return &container{}
}

func (c *container) ProjectConfig() *config.ProjectConfig {
	if c.projectConfig == nil {
		cfg, err := config.NewProjectConfig()
		if err != nil {
			log.Fatalf("failed to get project config: %s", err)
		}
		c.projectConfig = cfg
	}

	return c.projectConfig
}

func (c *container) APIServerConfig() *config.APIServerConfig {
	if c.apiServerConfig == nil {
		cfg, err := config.NewAPIServerConfig()
		if err != nil {
			log.Fatalf("failed to get API server config: %s", err)
		}
		c.apiServerConfig = cfg
	}

	return c.apiServerConfig
}

func (c *container) DBConfig() *config.DBConfig {
	if c.dbConfig == nil {
		cfg, err := config.NewDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err)
		}
		c.dbConfig = cfg
	}

	return c.dbConfig
}

func (c *container) Logger() *slog.Logger {
	if c.logger == nil {
		c.logger = logger.NewLogger(c.ProjectConfig())
	}

	return c.logger
}

func (c *container) Session() *session.Session {
	if c.session == nil {
		c.session = session.NewSession()
	}

	return c.session
}

func (c *container) DB() *sql.DB {
	if c.db == nil {
		c.db = database.NewPostgres(c.DBConfig())
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
