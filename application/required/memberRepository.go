package required

import (
	"context"
	"goplearn/domain"
	"goplearn/ent"
	"log"
)

// 会員の情報を保存や検索できる
type MemberRepository interface {
	Save(ctx context.Context, member *domain.Member) (*domain.Member, error)
}

func NewMemberRepository(client *ent.Client) MemberRepository {
	return &memberRepository{
		client: client,
	}
}

type memberRepository struct {
	client *ent.Client
}

// Save implements MemberRepository.
func (m *memberRepository) Save(ctx context.Context, member *domain.Member) (*domain.Member, error) {
	// 会員の情報を保存する
	savedMember, err := m.client.Member.Create().
		SetEmail(member.Email().Address).
		SetNickname(member.Nickname()).
		SetPasswordHash(member.PasswordHash()).
		SetStatus(member.Status()).
		Save(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	member.SetID(savedMember.ID)
	return member, nil
}
