// cmd/main.go
package main

import (
	"indiv/internal/application/usecases"
	"indiv/internal/infrastructure/repositories"
	"indiv/internal/interfaces/grpc"
	"indiv/internal/interfaces/migration"
	"indiv/internal/interfaces/rest"
	"indiv/pkg/config"
	"indiv/pkg/database"
	"indiv/pkg/logger"
	"log"
)

func main() {
	// Инициализация конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Инициализация логгера
	logger := logger.New(cfg.LogLevel)

	// Инициализация базы данных
	db, err := database.NewPostgresConnection(cfg.Database)
	if err != nil {
		logger.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Миграции
	if err := migration.Migrate(db); err != nil {
		logger.Fatalf("Ошибка миграции базы данных: %v", err)
	}

	// Инициализация репозиториев
	userRepo := repositories.NewUserRepository(db)
	// Аналогично для других репозиториев

	// Инициализация use cases
	userUseCase := usecases.NewUserUseCase(userRepo)
	// Аналогично для других use cases

	// Запуск gRPC сервера
	go func() {
		if err := grpc.RunServer(cfg.GRPCPort, userUseCase, logger); err != nil {
			logger.Fatalf("Ошибка запуска gRPC сервера: %v", err)
		}
	}()

	// Запуск REST Gateway
	if err := rest.RunRESTGateway(cfg.GRPCPort, cfg.RESTPort); err != nil {
		logger.Fatalf("Ошибка запуска REST Gateway: %v", err)
	}
}
