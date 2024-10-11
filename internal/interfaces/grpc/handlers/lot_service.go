// internal/interfaces/grpc/handlers/lot_service.go

package handlers

import (
	"context"
	"time"

	"indiv/internal/application/usecases"
	"indiv/internal/domain/entities"
	lotpb "indiv/proto/v1/lot"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LotServiceServer struct {
	useCase *usecases.LotUseCase
	lotpb.UnimplementedLotServiceServer
}

func NewLotServiceServer(useCase *usecases.LotUseCase) *LotServiceServer {
	return &LotServiceServer{useCase: useCase}
}

// CreateLot создает новый лот
func (s *LotServiceServer) CreateLot(ctx context.Context, req *lotpb.CreateLotRequest) (*lotpb.CreateLotResponse, error) {
	auctionStart, err := time.Parse(time.RFC3339, req.AuctionStart)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат времени начала аукциона")
	}
	auctionEnd, err := time.Parse(time.RFC3339, req.AuctionEnd)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат времени окончания аукциона")
	}

	newLot := &entities.Lot{
		SellerID:        req.SellerId,
		Title:           req.Title,
		Description:     req.Description,
		StartingPrice:   req.StartingPrice,
		MinBidIncrement: req.MinBidIncrement,
		AuctionStart:    auctionStart,
		AuctionEnd:      auctionEnd,
	}

	err = s.useCase.CreateLot(ctx, newLot)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка создания лота: %v", err)
	}

	return &lotpb.CreateLotResponse{
		Lot: &lotpb.Lot{
			Id:              newLot.ID,
			SellerId:        newLot.SellerID,
			Title:           newLot.Title,
			Description:     newLot.Description,
			StartingPrice:   newLot.StartingPrice,
			MinBidIncrement: newLot.MinBidIncrement,
			AuctionStart:    newLot.AuctionStart.Format(time.RFC3339),
			AuctionEnd:      newLot.AuctionEnd.Format(time.RFC3339),
		},
	}, nil
}

// GetLotByID получает лот по ID
func (s *LotServiceServer) GetLotByID(ctx context.Context, req *lotpb.GetLotByIDRequest) (*lotpb.GetLotByIDResponse, error) {
	lot, err := s.useCase.GetLotByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка получения лота: %v", err)
	}
	if lot == nil {
		return nil, status.Errorf(codes.NotFound, "Лот не найден")
	}

	return &lotpb.GetLotByIDResponse{
		Lot: &lotpb.Lot{
			Id:              lot.ID,
			SellerId:        lot.SellerID,
			Title:           lot.Title,
			Description:     lot.Description,
			StartingPrice:   lot.StartingPrice,
			MinBidIncrement: lot.MinBidIncrement,
			AuctionStart:    lot.AuctionStart.Format(time.RFC3339),
			AuctionEnd:      lot.AuctionEnd.Format(time.RFC3339),
		},
	}, nil
}

// ListLots возвращает список лотов
func (s *LotServiceServer) ListLots(ctx context.Context, req *lotpb.ListLotsRequest) (*lotpb.ListLotsResponse, error) {
	lots, err := s.useCase.ListLots(ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка получения списка лотов: %v", err)
	}

	var lotList []*lotpb.Lot
	for _, lot := range lots {
		lotList = append(lotList, &lotpb.Lot{
			Id:              lot.ID,
			SellerId:        lot.SellerID,
			Title:           lot.Title,
			Description:     lot.Description,
			StartingPrice:   lot.StartingPrice,
			MinBidIncrement: lot.MinBidIncrement,
			AuctionStart:    lot.AuctionStart.Format(time.RFC3339),
			AuctionEnd:      lot.AuctionEnd.Format(time.RFC3339),
		})
	}

	return &lotpb.ListLotsResponse{
		Lots: lotList,
	}, nil
}
