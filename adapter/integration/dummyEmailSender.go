package integration

import (
	"context"
	"goplearn/domain"
	"log"
)

type dummyEmailSender struct {
}

func (d *dummyEmailSender) Send(ctx context.Context, email domain.Email, subject string, body string) error {
	log.Println("DummyEmailSender send email : ", email)
	return nil
}
