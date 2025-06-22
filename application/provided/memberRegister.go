package provided

import "goplearn/domain"

// 会員の登録と関連する機能を提供する
type MemberRegister interface {
	Register(registerRequest *domain.MemberRegisterRequest) (*domain.Member, error)
}
