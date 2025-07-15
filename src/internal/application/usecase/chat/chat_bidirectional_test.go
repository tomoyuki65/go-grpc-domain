package chat

import (
	"context"
	"fmt"
	"io"
	"testing"

	mockLogger "go-grpc-domain/internal/application/usecase/logger/mock_logger"
	pb "go-grpc-domain/pb/chat"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc"
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

func TestChatUsecase_Bidirectional(t *testing.T) {
	// モック
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ロガーのモック
	mockLogger := mockLogger.NewMockLogger(ctrl)

	t.Run("正常終了", func(t *testing.T) {
		// ユースケースのインスタンス化
		chatUsecase := NewChatUsecase(mockLogger)

		// パラメータの設定
		mockBidirectionalStream := &mockBidirectionalStream{
			ctx:       context.Background(),
			recv:      []*pb.TextInput{{Text: "Hello"}},
			recvError: nil,
			sendError: nil,
		}

		// テストの実行
		err := chatUsecase.Bidirectional(mockBidirectionalStream)

		// 検証
		assert.Equal(t, nil, err)
	})

	t.Run("受信エラー", func(t *testing.T) {
		// ユースケースのインスタンス化
		chatUsecase := NewChatUsecase(mockLogger)

		// パラメータの設定
		mockBidirectionalStream := &mockBidirectionalStream{
			ctx:       context.Background(),
			recv:      []*pb.TextInput{{Text: "Hello"}},
			recvError: fmt.Errorf("error"),
			sendError: nil,
		}

		// テストの実行
		err := chatUsecase.Bidirectional(mockBidirectionalStream)

		// 検証
		assert.NotEqual(t, nil, err)
	})

	t.Run("バリデーションエラー", func(t *testing.T) {
		// モック化
		mockLogger.EXPECT().Warn(gomock.Any(), gomock.Any()).Return()

		// ユースケースのインスタンス化
		chatUsecase := NewChatUsecase(mockLogger)

		// パラメータの設定
		mockBidirectionalStream := &mockBidirectionalStream{
			ctx:       context.Background(),
			recv:      []*pb.TextInput{{Text: ""}},
			recvError: nil,
			sendError: nil,
		}

		// テストの実行
		err := chatUsecase.Bidirectional(mockBidirectionalStream)

		// 検証
		assert.NotEqual(t, nil, err)
	})

	t.Run("送信エラー", func(t *testing.T) {
		// ユースケースのインスタンス化
		chatUsecase := NewChatUsecase(mockLogger)

		// パラメータの設定
		mockBidirectionalStream := &mockBidirectionalStream{
			ctx:       context.Background(),
			recv:      []*pb.TextInput{{Text: "Hello"}},
			recvError: nil,
			sendError: fmt.Errorf("error"),
		}

		// テストの実行
		err := chatUsecase.Bidirectional(mockBidirectionalStream)

		// 検証
		assert.NotEqual(t, nil, err)
	})
}
