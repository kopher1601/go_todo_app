package provided

import (
	"context"
	"goplearn/domain"
)

// 会員の登録と関連する機能を提供する
type MemberRegister interface {
	Register(ctx context.Context, registerRequest *domain.MemberRegisterRequest) (*domain.Member, error)
}
