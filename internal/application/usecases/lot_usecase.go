// internal/application/usecases/lot_usecase.go

package usecases

import (
	"context"

	"indiv/internal/domain/entities"
	"indiv/internal/domain/repositories"
)

type LotUseCase struct {
	lotRepo repositories.LotRepository
}

func NewLotUseCase(lotRepo repositories.LotRepository) *LotUseCase {
	return &LotUseCase{lotRepo: lotRepo}
}

func (uc *LotUseCase) CreateLot(ctx context.Context, lot *entities.Lot) error {
	return uc.lotRepo.Create(ctx, lot)
}

func (uc *LotUseCase) GetLotByID(ctx context.Context, id int64) (*entities.Lot, error) {
	return uc.lotRepo.GetByID(ctx, id)
}

func (uc *LotUseCase) ListLots(ctx context.Context, page, pageSize int32) ([]*entities.Lot, error) {
	return uc.lotRepo.List(ctx, page, pageSize)
}
