package provided_test

import (
	"context"
	"goplearn/application"
	"goplearn/application/required"
	"goplearn/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type memberRepositoryStub struct {
}

func (m *memberRepositoryStub) Save(ctx context.Context, member *domain.Member) (*domain.Member, error) {
	member.SetID(1)
	return member, nil
}

func NewMemberRepositoryStub() required.MemberRepository {
	return &memberRepositoryStub{}
}

type emailSenderStub struct {
}

func NewEmailSenderStub() required.EmailSender {
	return &emailSenderStub{}
}

func (e *emailSenderStub) Send(ctx context.Context, to *domain.Email, subject string, body string) error {
	return nil
}

type emailSenderMock struct {
	Tos []domain.Email
}

func NewEmailSenderMock() *emailSenderMock {
	return &emailSenderMock{
		Tos: []domain.Email{},
	}
}

func (e *emailSenderMock) Send(ctx context.Context, to *domain.Email, subject string, body string) error {
	e.Tos = append(e.Tos, *to)
	return nil
}

func (e *emailSenderMock) GetTos() []domain.Email {
	return e.Tos
}

func TestMemberRegister_Stub(t *testing.T) {
	memberRegister := application.NewMemberRegister(
		NewMemberRepositoryStub(),
		NewEmailSenderStub(),
		domain.NewMockPasswordEncoder(),
	)

	member, err := memberRegister.Register(context.Background(), domain.CreateMockMemgerRegisterRequest())
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, member)
	assert.Equal(t, member.ID(), 1)
	assert.Equal(t, member.Email().Address, domain.CreateMockMemgerRegisterRequest().Email)
	assert.Equal(t, member.Nickname(), domain.CreateMockMemgerRegisterRequest().Nickname)
	assert.Equal(t, member.PasswordHash(), "SECRET")
	assert.Equal(t, member.Status(), domain.MemberStatusPending)
}

func TestMemberRegister_Mock(t *testing.T) {
	emailSenderMock := &emailSenderMock{
		Tos: []domain.Email{},
	}
	memberRegister := application.NewMemberRegister(
		NewMemberRepositoryStub(),
		emailSenderMock,
		domain.NewMockPasswordEncoder(),
	)

	member, err := memberRegister.Register(context.Background(), domain.CreateMockMemgerRegisterRequest())
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, member)
	assert.Equal(t, member.ID(), 1)
	assert.Equal(t, len(emailSenderMock.GetTos()), 1)
	assert.Equal(t, emailSenderMock.GetTos()[0].Address, member.Email().Address)
	assert.Equal(t, member.Email().Address, domain.CreateMockMemgerRegisterRequest().Email)
	assert.Equal(t, member.Nickname(), domain.CreateMockMemgerRegisterRequest().Nickname)
	assert.Equal(t, member.PasswordHash(), "SECRET")
	assert.Equal(t, member.Status(), domain.MemberStatusPending)
}
