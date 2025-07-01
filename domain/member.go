package domain

import (
	"errors"
)

var ErrIllegalState = errors.New("illegal state")

type Member struct {
	ID           int
	Email        Email
	Nickname     string
	PasswordHash string
	Status       MemberStatus
	Detail       *MemberDetail // ID 를 갖는게 좋을까?
}

func RegisterMember(registerRequest *MemberRegisterRequest, passwordEncoder PasswordEncoder) (*Member, error) {
	passwordHash, err := passwordEncoder.Encode(registerRequest.Password)
	if err != nil {
		return nil, err
	}

	email, err := NewEmail(registerRequest.Email)
	if err != nil {
		return nil, err
	}

	return &Member{
		Email:        email,
		Nickname:     registerRequest.Nickname,
		PasswordHash: passwordHash,
		Status:       MemberStatusPending,
	}, nil
}

func NewMember(id int, email Email, nickname string, passwordHash string, status MemberStatus) *Member {
	return &Member{
		ID:           id,
		Email:        email,
		Nickname:     nickname,
		PasswordHash: passwordHash,
		Status:       status,
	}
}

func (m *Member) Activate() error {
	if m.Status != MemberStatusPending {
		return ErrIllegalState
	}

	m.Status = MemberStatusActive
	return nil
}

func (m *Member) Deactivate() error {
	if m.Status != MemberStatusActive {
		return ErrIllegalState
	}

	m.Status = MemberStatusDeactivated
	return nil
}

func (m *Member) VerifyPassword(password string, passwordEncoder PasswordEncoder) bool {
	return passwordEncoder.Matches(password, m.PasswordHash)
}

func (m *Member) ChangeNickname(nickname string) {
	m.Nickname = nickname
}

func (m *Member) ChangeEmail(address string) error {
	email, err := NewEmail(address)
	if err != nil {
		return err
	}

	m.Email = email

	return nil
}

func (m *Member) ChangePassword(password string, passwordEncoder PasswordEncoder) error {
	passwordHash, err := passwordEncoder.Encode(password)
	if err != nil {
		return err
	}

	m.PasswordHash = passwordHash
	return nil
}

func (m *Member) IsActive() bool {
	return m.Status == MemberStatusActive
}

func (m *Member) SetID(id int) {
	m.ID = id
}
