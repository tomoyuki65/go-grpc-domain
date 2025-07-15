package registry

import (
	usecaseChat "go-grpc-domain/internal/application/usecase/chat"
	"go-grpc-domain/internal/infrastructure/logger"
	serverGrpcChat "go-grpc-domain/internal/presentation/server/grpc/chat"
)

// ハンドラーをまとめるコントローラー構造体
type GrpcController struct {
	Chat *serverGrpcChat.ChatServer
}

func NewGrpcController() *GrpcController {
	// ロガー設定
	logger := logger.NewSlogLogger()

	// chatドメインのハンドラー設定
	chatUsecase := usecaseChat.NewChatUsecase(logger)
	chatServer := serverGrpcChat.NewChatServer(chatUsecase)

	return &GrpcController{
		Chat: chatServer,
	}
}
