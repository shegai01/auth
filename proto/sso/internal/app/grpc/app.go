package grpc

import (
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	grpcServer *grpc.Server
	port       int
}

func New(log *slog.Logger, authService AuthService, port int) *App {

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(),
	))
	authgrpc.Register(gRPCServer, authService)
	return &App{
		log:        log,
		grpcServer: grpcServer,
		port:       port,
	}
}
