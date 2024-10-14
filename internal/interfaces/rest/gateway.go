// internal/interfaces/rest/gateway.go

package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	auctionpb "indiv/proto/v1/auction"
	bidpb "indiv/proto/v1/bid"
	lotpb "indiv/proto/v1/lot"
	userpb "indiv/proto/v1/user"
)

func RunRESTGateway(grpcPort, restPort string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	// Обратите внимание на формат grpcPort (например, ":50051")
	endpoint := fmt.Sprintf("localhost%s", grpcPort)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Регистрация хендлеров для каждого сервиса
	err := userpb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return fmt.Errorf("ошибка регистрации сервиса пользователя: %v", err)
	}

	err = lotpb.RegisterLotServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return fmt.Errorf("ошибка регистрации сервиса лотов: %v", err)
	}

	err = bidpb.RegisterBidServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return fmt.Errorf("ошибка регистрации сервиса ставок: %v", err)
	}

	err = auctionpb.RegisterAuctionServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return fmt.Errorf("ошибка регистрации сервиса аукционов: %v", err)
	}

	fmt.Printf("REST Gateway запущен на %s\n", restPort)
	return http.ListenAndServe(restPort, mux)
}
