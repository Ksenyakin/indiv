// internal/infrastructure/workers/auction_worker.go

package workers

import (
	"context"
	"go.uber.org/zap"
	"indiv/internal/application/usecases"
	"time"
)

type AuctionWorker struct {
	auctionUC *usecases.AuctionUseCase
	logger    *zap.SugaredLogger
}

func NewAuctionWorker(auctionUC *usecases.AuctionUseCase, logger *zap.SugaredLogger) *AuctionWorker {
	return &AuctionWorker{
		auctionUC: auctionUC,
		logger:    logger,
	}
}

func (w *AuctionWorker) Run() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			w.processAuctions()
		}
	}
}

func (w *AuctionWorker) processAuctions() {
	ctx := context.Background()
	auctions, err := w.auctionUC.GetAuctionsToClose(ctx)
	if err != nil {
		w.logger.Errorf("Ошибка получения аукционов для закрытия: %v", err)
		return
	}

	for _, auction := range auctions {
		_, err := w.auctionUC.CloseAuction(ctx, auction.ID)
		if err != nil {
			w.logger.Errorf("Ошибка закрытия аукциона %d: %v", auction.ID, err)
			continue
		}
		w.logger.Infof("Аукцион %d успешно закрыт", auction.ID)
		// Отправка уведомлений может быть реализована здесь
	}
}
