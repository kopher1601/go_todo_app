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

func CreateMember(createRequest *MemberCreateRequest, passwordEncoder PasswordEncoder) (*Member, error) {
	passwordHash, err := passwordEncoder.Encode(createRequest.Password)
	if err != nil {
		return nil, err
	}

	return &Member{
		email:        createRequest.Email,
		nickname:     createRequest.Nickname,
		passwordHash: passwordHash,
		status:       MemberStatusPending,
	}, nil
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

func (m *Member) IsActive() bool {
	return m.status == MemberStatusActive
}
