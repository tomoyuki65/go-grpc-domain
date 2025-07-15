package chat

import (
	"errors"
	"fmt"
	"io"

	pb "go-grpc-domain/pb/chat"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *chatUsecase) Bidirectional(stream grpc.BidiStreamingServer[pb.TextInput, pb.TextOutput]) error {
	// コンテキストを取得
	ctx := stream.Context()

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}

		// バリデーションチェック
		if err := req.Validate(); err != nil {
			msg := fmt.Sprintf("バリデーションエラー：%s", err.Error())
			u.logger.Warn(ctx, msg)

			return status.Errorf(codes.InvalidArgument, "%s", err.Error())
		}

		msg := fmt.Sprintf("Hello, %s !!", req.GetText())
		if err := stream.Send(&pb.TextOutput{Text: msg}); err != nil {
			return err
		}
	}
}
