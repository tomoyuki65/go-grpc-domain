package router

import (
	ic "go-grpc-domain/internal/presentation/interceptor"
	"go-grpc-domain/internal/registry"
	pbChat "go-grpc-domain/pb/chat"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetupGrpcServer(c *registry.GrpcController, i *ic.Interceptor) *grpc.Server {
	// gRPCサーバーの作成
	s := grpc.NewServer(
		// インターセプターの適用
		grpc.ChainUnaryInterceptor(
			i.RequestUnary,
			i.AuthUnary,
		),
		grpc.ChainStreamInterceptor(
			i.RequestStream,
			i.AuthStream,
		),
	)

	// サービス設定
	pbChat.RegisterChatServiceServer(s, c.Chat)

	// リフレクション設定
	reflection.Register(s)

	return s
}
