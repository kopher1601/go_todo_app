package domain

import (
	"errors"
)

var ErrIllegalState = errors.New("illegal state")

type Member struct {
	id           int
	email        Email
	nickname     string
	passwordHash string
	status       MemberStatus
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
		email:        email,
		nickname:     registerRequest.Nickname,
		passwordHash: passwordHash,
		status:       MemberStatusPending,
	}, nil
}

func NewMember(id int, email Email, nickname string, passwordHash string, status MemberStatus) *Member {
	return &Member{
		id:           id,
		email:        email,
		nickname:     nickname,
		passwordHash: passwordHash,
		status:       status,
	}
}

func (m *Member) Activate() error {
	if m.status != MemberStatusPending {
		return ErrIllegalState
	}

	m.status = MemberStatusActive
	return nil
}

func (m *Member) Deactivate() error {
	if m.status != MemberStatusActive {
		return ErrIllegalState
	}

	m.status = MemberStatusDeactivated
	return nil
}

func (m *Member) VerifyPassword(password string, passwordEncoder PasswordEncoder) bool {
	return passwordEncoder.Matches(password, m.passwordHash)
}

func (m *Member) ChangeNickname(nickname string) {
	m.nickname = nickname
}

func (m *Member) ChangeEmail(address string) error {
	email, err := NewEmail(address)
	if err != nil {
		return err
	}

	m.email = email

	return nil
}

func (m *Member) ChangePassword(password string, passwordEncoder PasswordEncoder) error {
	passwordHash, err := passwordEncoder.Encode(password)
	if err != nil {
		return err
	}

	m.passwordHash = passwordHash
	return nil
}

func (m *Member) IsActive() bool {
	return m.status == MemberStatusActive
}

func (m *Member) SetID(id int) {
	m.id = id
}

// id 필드 getter
func (m *Member) ID() int {
	return m.id
}

// email 필드 getter
func (m *Member) Email() *Email {
	return &m.email
}

// nickname 필드 getter
func (m *Member) Nickname() string {
	return m.nickname
}

// passwordHash 필드 getter
func (m *Member) PasswordHash() string {
	return m.passwordHash
}

// status 필드 getter
func (m *Member) Status() MemberStatus {
	return m.status
}
