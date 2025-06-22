package required

import "goplearn/domain"

// Email を送信する
type EmailSender interface {
	Send(email *domain.Email, subject string, body string) error
}
