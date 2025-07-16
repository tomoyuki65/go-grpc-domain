package chat

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	mockChat "go-grpc-domain/internal/application/usecase/chat/mock_chat"
	chatModel "go-grpc-domain/internal/domain/chat"
	pb "go-grpc-domain/pb/chat"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// Bidirectionalのstream用のモック構造体
type mockBidirectionalStream struct {
	grpc.BidiStreamingServer[pb.TextInput, pb.TextOutput]
	ctx       context.Context
	recv      []*pb.TextInput
	sent      []*pb.TextOutput
	recvIndex int
	sendError error
	recvError error
}

func (m *mockBidirectionalStream) Recv() (*pb.TextInput, error) {
	if m.recvError != nil {
		return nil, m.recvError
	}
	if m.recvIndex >= len(m.recv) {
		return nil, io.EOF
	}
	req := m.recv[m.recvIndex]
	m.recvIndex++
	return req, nil
}

func (m *mockBidirectionalStream) Send(resp *pb.TextOutput) error {
	if m.sendError != nil {
		return m.sendError
	}
	m.sent = append(m.sent, resp)
	return nil
}

func (m *mockBidirectionalStream) Context() context.Context {
	return m.ctx
}

// 初期処理
func init() {
	// テスト用の環境変数ファイル「.env.testing」を読み込んで使用する。
	if err := godotenv.Load("../../../../../.env.testing"); err != nil {
		fmt.Println(".env.testingの読み込みに失敗しました。")
	}
}

func TestChatServer_Bidirectional(t *testing.T) {
	// モック
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockChatUsecase := mockChat.NewMockChatUsecase(ctrl)

	// gRPCクライアントの設定
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}
	target := fmt.Sprintf("dns:///localhost:%s", grpcPort)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewChatServiceClient(conn)

	t.Run("ChatServerが正常終了すること", func(t *testing.T) {
		// モック
		mockChatUsecase.EXPECT().Bidirectional(gomock.Any()).Return(nil)

		// ChatServerのインスタンス化
		chatServer := NewChatServer(mockChatUsecase)

		// パラメータの設定
		mockBidirectionalStream := &mockBidirectionalStream{
			ctx:       context.Background(),
			recv:      []*pb.TextInput{{Text: "Hello"}},
			recvError: nil,
			sendError: nil,
		}

		// テストの実行
		err := chatServer.Bidirectional(mockBidirectionalStream)

		// 検証
		assert.Equal(t, nil, err)
	})

	t.Run("レスポンス結果が想定通りであること", func(t *testing.T) {
		// メタデータにauthorizationを追加
		ctx := context.Background()
		md := metadata.New(map[string]string{"authorization": "Bearer token"})
		ctx = metadata.NewOutgoingContext(ctx, md)

		// テストの実行
		stream, err := client.Bidirectional(ctx)
		if err != nil {
			t.Fatalf("Failed to call client.Bidirectional: %v", err)
		}

		var sendEnd, recvEnd bool
		inputText := "Hello"
		for !(sendEnd && recvEnd) {
			// 送信処理
			if !sendEnd {
				if err := stream.Send(&pb.TextInput{Text: inputText}); err != nil {
					t.Fatalf("Failed to stream.Send: %v", err)
				}
				if err := stream.CloseSend(); err != nil {
					t.Fatalf("Failed to stream.CloseSend: %v", err)
				}
				sendEnd = true
			}

			// 受信処理
			if !recvEnd {
				res, err := stream.Recv()
				if err != nil {
					t.Fatalf("Failed to stream.Recv: %v", err)
				}

				// 検証
				chat := chatModel.NewChat(inputText)
				msg := fmt.Sprintf("ToUppder: %s, ToLower: %s, AddTimeNow: %s", chat.TextToUpper(), chat.TextToLower(), chat.TextAddTimeNow())
				assert.Equal(t, msg, res.Text)

				recvEnd = true
			}
		}
	})
}
