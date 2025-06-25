package domain

type MemberStatus int

const (
	MemberStatusPending MemberStatus = iota
	MemberStatusActive
	MemberStatusDeactivated
)

func NewMemberStatus(status string) MemberStatus {
	switch status {
	case "PENDING":
		return MemberStatusPending
	case "ACTIVE":
		return MemberStatusActive
	case "DEACTIVATED":
		return MemberStatusDeactivated
	default:
		return MemberStatusPending
	}

}

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
