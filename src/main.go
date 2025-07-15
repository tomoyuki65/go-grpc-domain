package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	"go-grpc-domain/internal/infrastructure/logger"
	ic "go-grpc-domain/internal/presentation/interceptor"
	"go-grpc-domain/internal/presentation/router"
	"go-grpc-domain/internal/registry"

	"github.com/joho/godotenv"
)

func main() {
	// .env ファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		slog.Error(fmt.Sprintf(".envファイルの読み込みに失敗しました。: %v", err))
	}

	// ENVの設定
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	// gRPC用のポート番号の設定
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}

	// Listenerの設定
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		slog.Error(fmt.Sprintf("Listenerの設定に失敗しました。: %v", err))
	}

	// サーバー設定
	c := registry.NewGrpcController()
	l := logger.NewSlogLogger()
	i := ic.NewInterceptor(l)
	s := router.SetupGrpcServer(c, i)

	// gRPCサーバーの起動（非同期）
	slog.Info(fmt.Sprintf("[ENV=%s] start gRPC-Server port: %s", env, grpcPort))
	if err := s.Serve(listener); err != nil {
		slog.Error(fmt.Sprintf("gRPC-Server の起動に失敗しました。: %v", err))
	}
}
