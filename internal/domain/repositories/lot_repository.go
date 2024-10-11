// internal/domain/repositories/lot_repository.go

package repositories

import (
	"context"
	"indiv/internal/domain/entities"
)

// LotRepository определяет методы для доступа к данным лотов.
type LotRepository interface {
	// Create добавляет новый лот в репозиторий.
	Create(ctx context.Context, lot *entities.Lot) error

	// GetByID получает лот по его ID.
	GetByID(ctx context.Context, id int64) (*entities.Lot, error)

	// List получает список лотов с поддержкой пагинации.
	List(ctx context.Context, page, pageSize int32) ([]*entities.Lot, error)
}
