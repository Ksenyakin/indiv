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

	"google.golang.org/grpc"
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
	lotRepo := repositories.NewLotRepository(db)
	bidRepo := repositories.NewBidRepository(db)
	auctionRepo := repositories.NewAuctionRepository(db)

	// Инициализация use cases
	userUseCase := usecases.NewUserUseCase(userRepo)
	lotUseCase := usecases.NewLotUseCase(lotRepo)
	bidUseCase := usecases.NewBidUseCase(bidRepo, auctionRepo, userRepo)
	auctionUseCase := usecases.NewAuctionUseCase(auctionRepo, bidRepo, userRepo)

	// Запуск воркера
	auctionWorker := workers.NewAuctionWorker(auctionUseCase, logger)
	go auctionWorker.Run()

	// Запуск gRPC сервера
	go func() {
		if err := grpc.RunServer(cfg.GRPCPort, userUseCase, lotUseCase, bidUseCase, auctionUseCase, logger); err != nil {
			logger.Fatalf("Ошибка запуска gRPC сервера: %v", err)
		}
	}()

	// Запуск REST Gateway
	if err := rest.RunRESTGateway(cfg.GRPCPort, cfg.RESTPort); err != nil {
		logger.Fatalf("Ошибка запуска REST Gateway: %v", err)
	}
}
