package chat

import (
	"go-grpc-domain/internal/application/usecase/logger"
	pb "go-grpc-domain/pb/chat"

	"google.golang.org/grpc"
)

type ChatUsecase interface {
	Bidirectional(stream grpc.BidiStreamingServer[pb.TextInput, pb.TextOutput]) error
}

type chatUsecase struct {
	logger logger.Logger
}

func NewChatUsecase(logger logger.Logger) ChatUsecase {
	return &chatUsecase{
		logger: logger,
	}
}
