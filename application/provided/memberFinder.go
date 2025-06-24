package provided

import (
	"context"
	"goplearn/domain"
)

type MemberFinder interface {
	Find(ctx context.Context, memberID int) (*domain.Member, error)
}
