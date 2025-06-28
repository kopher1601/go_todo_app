package application

import (
	"context"
	"fmt"
	"goplearn/application/provided"
	"goplearn/application/required"
	"goplearn/domain"

	"github.com/go-playground/validator/v10"
)

type memberRegister struct {
	memberRepository required.MemberRepository
	emailSender      required.EmailSender
	passwordEncoder  domain.PasswordEncoder
}

func NewMemberRegister(
	memberRepository required.MemberRepository,
	emailSender required.EmailSender,
	passwordEncoder domain.PasswordEncoder,
) provided.MemberRegister {
	return &memberRegister{
		memberRepository: memberRepository,
		emailSender:      emailSender,
		passwordEncoder:  passwordEncoder,
	}
}

func validateRequest(registerRequest *domain.MemberRegisterRequest) error {
	if err := registerRequest.Validate(); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			fmt.Println(err.Field(), err.Tag())
		}
		return err
	}
	return nil
}

func (m *memberRegister) sendWelcomeEmail(ctx context.Context, member *domain.Member) error {
	return m.emailSender.Send(ctx, member.Email, "登録を完了してください", "下記のリンクをクリックして登録を完了してください")
}

func (m *memberRegister) checkDuplicateEmail(ctx context.Context, registerRequest *domain.MemberRegisterRequest) error {
	email, err := domain.NewEmail(registerRequest.Email)
	if err != nil {
		return err
	}
	if foundEmail, err := m.memberRepository.FindByEmail(ctx, email); err == nil && foundEmail != nil {
		return domain.ErrDuplicateEmail
	}
	return nil
}

func (m *memberRegister) Register(ctx context.Context, registerRequest *domain.MemberRegisterRequest) (*domain.Member, error) {
	err := validateRequest(registerRequest)
	if err != nil {
		return nil, err
	}

	err = m.checkDuplicateEmail(ctx, registerRequest)
	if err != nil {
		return nil, err
	}

	member, err := domain.RegisterMember(registerRequest, m.passwordEncoder)
	if err != nil {
		return nil, err
	}

	member, err = m.memberRepository.Save(ctx, member)
	if err != nil {
		return nil, err
	}

	err = m.sendWelcomeEmail(ctx, member)
	if err != nil {
		return nil, err
	}

	return member, nil
}

func (m *memberRegister) Activate(ctx context.Context, memberId string) error {
	member, err := m.memberRepository.FindByID(ctx, memberId)
	if err != nil {
		return err
	}

	err = member.Activate()
	if err != nil {
		return err
	}

	_, err = m.memberRepository.Update(ctx, member)

	return nil
}
