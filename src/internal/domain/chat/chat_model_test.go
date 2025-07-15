package chat

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewChat(t *testing.T) {
	t.Run("ChatModel作成", func(t *testing.T) {
		text := "Hello"

		// ChatModel作成
		chat := NewChat(text)

		// 検証
		assert.Equal(t, text, chat.InputText)
	})
}

func TestChat_TextToUpper(t *testing.T) {
	t.Run("大文字に変換", func(t *testing.T) {
		text := "Hello"

		// ChatModel作成
		chat := NewChat(text)

		// 検証
		assert.Equal(t, strings.ToUpper(text), chat.TextToUpper())
	})
}

func TestChat_TextToLower(t *testing.T) {
	t.Run("小文字に変換", func(t *testing.T) {
		text := "Hello"

		// ChatModel作成
		chat := NewChat(text)

		// 検証
		assert.Equal(t, strings.ToLower(text), chat.TextToLower())
	})
}

func TestChat_TextAddTimeNow(t *testing.T) {
	text := "Hello"

	// ChatModel作成
	chat := NewChat(text)

	// 検証
	jst, _ := time.LoadLocation("Asia/Tokyo")
	assert.Equal(t, fmt.Sprintf("[%s] %s", time.Now().In(jst).Format("2006-01-02 15:04:05"), text), chat.TextAddTimeNow())
}
