package domain

type MemberStatus int

const (
	MemberStatusPending MemberStatus = iota
	MemberStatusActive
	MemberStatusDeactivated
)

func (m MemberStatus) String() string {
	switch m {
	case MemberStatusPending:
		return "PENDING"
	case MemberStatusActive:
		return "ACTIVE"
	case MemberStatusDeactivated:
		return "DEACTIVATED"
	}
	return "UNKNOWN"
}
