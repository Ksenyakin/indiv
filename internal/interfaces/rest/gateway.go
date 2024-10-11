// internal/interfaces/rest/gateway.go
package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/yourusername/auction-system/api/proto"
	"google.golang.org/grpc"
	"net/http"
)

func RunRESTGateway(grpcPort, restPort string) error {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := proto.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcPort, opts); err != nil {
		return err
	}
	return http.ListenAndServe(restPort, mux)
}
