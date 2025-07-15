package chat

import (
	"errors"
	"fmt"
	"io"

	pb "go-grpc-domain/pb/chat"
	chatModel "go-grpc-domain/internal/domain/chat"

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

		// Chatドメインを利用してメッセージを設定
		chat := chatModel.NewChat(req.GetText())
		msg := fmt.Sprintf("ToUppder: %s, ToLower: %s, AddTimeNow: %s", chat.TextToUpper(), chat.TextToLower(), chat.TextAddTimeNow())

		if err := stream.Send(&pb.TextOutput{Text: msg}); err != nil {
			return err
		}
	}
}
