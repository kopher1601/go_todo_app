package required

import (
	"context"
	"goplearn/domain"
)

// Email を送信する
type EmailSender interface {
	Send(ctx context.Context, email domain.Email, subject string, body string) error
}
