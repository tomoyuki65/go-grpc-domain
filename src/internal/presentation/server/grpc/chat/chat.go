package chat

import (
	usecase "go-grpc-domain/internal/application/usecase/chat"
	pb "go-grpc-domain/pb/chat"

	"google.golang.org/grpc"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
	chatUsecase usecase.ChatUsecase
}

func NewChatServer(
	chatUsecase usecase.ChatUsecase,
) *ChatServer {
	return &ChatServer{
		chatUsecase: chatUsecase,
	}
}

func (s *ChatServer) Bidirectional(stream grpc.BidiStreamingServer[pb.TextInput, pb.TextOutput]) error {
	return s.chatUsecase.Bidirectional(stream)
}
