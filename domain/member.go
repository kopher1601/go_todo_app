package domain

import "errors"

var ErrIllegalState = errors.New("illegal state")

type MemberStatus string

const (
	MemberStatusPending     MemberStatus = "PENDING"
	MemberStatusActive      MemberStatus = "ACTIVE"
	MemberStatusDeactivated MemberStatus = "DEACTIVATED"
)

type Member struct {
	email        string
	nickname     string
	passwordHash string
	status       MemberStatus
}

func newMember(email, nickname, passwordHash string) *Member {
	return &Member{
		email:        email,
		nickname:     nickname,
		passwordHash: passwordHash,
		status:       MemberStatusPending,
	}
}

func CreateMember(email, nickname, password string, passwordEncoder PasswordEncoder) (*Member, error) {
	passwordHash, err := passwordEncoder.Encode(password)
	if err != nil {
		return nil, err
	}

	return newMember(email, nickname, passwordHash), nil
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

func (m *Member) ChangePassword(password string, passwordEncoder PasswordEncoder) error {
	passwordHash, err := passwordEncoder.Encode(password)
	if err != nil {
		return err
	}

	m.passwordHash = passwordHash
	return nil
}
