//go:generate mockgen -source=memberRegister.go -destination=mockMemberRegister.go -package provided
package provided

import (
	"context"
	"goplearn/domain"
)

// MemberRegister 会員の登録と関連する機能を提供する
type MemberRegister interface {
	Register(ctx context.Context, registerRequest *domain.MemberRegisterRequest) (*domain.Member, error)
	Activate(ctx context.Context, memberId string) error
}
