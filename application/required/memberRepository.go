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
	FindByID(ctx context.Context, memberId string) (*domain.Member, error)
	Update(ctx context.Context, member *domain.Member) (*domain.Member, error)
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
		SetEmail(member.Email.Address).
		SetNickname(member.Nickname).
		SetPasswordHash(member.PasswordHash).
		SetStatus(member.Status.String()).
		Save(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	member.SetID(savedMember.ID)
	return member, nil
}

func (m *memberRepository) FindByID(ctx context.Context, memberId string) (*domain.Member, error) {
	panic("unimplemented")
}

func (m *memberRepository) Update(ctx context.Context, member *domain.Member) (*domain.Member, error) {
	updatedMember, err := m.client.Member.UpdateOne(
		&ent.Member{
			Email:        member.Email.Address,
			Nickname:     member.Nickname,
			PasswordHash: member.PasswordHash,
			Status:       member.Status.String(),
		},
	).Save(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	email, err := domain.NewEmail(updatedMember.Email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return domain.NewMember(
		updatedMember.ID,
		email,
		updatedMember.Nickname,
		updatedMember.PasswordHash,
		domain.NewMemberStatus(updatedMember.Status),
	), nil
}
