//go:generate mockgen -source=memberFinder.go -destination=mockMemberFinder.go -package provided
package provided

import (
	"context"
	"goplearn/domain"
)

// MemberFinder 会員の検索機能を提供する
type MemberFinder interface {
	Find(ctx context.Context, memberID int) (*domain.Member, error)
}
