//go:generate mockgen -source=memberRepository.go -destination=mockMemberRepository.go -package required
package required

import (
	"context"
	"goplearn/domain"
	"goplearn/ent"
	"goplearn/ent/member"
	"log"
)

// MemberRepository 会員の情報を保存や検索できる
type MemberRepository interface {
	Save(ctx context.Context, tx *ent.Tx, member *domain.Member) (*domain.Member, error)
	FindByID(ctx context.Context, memberId string) (*domain.Member, error)
	Update(ctx context.Context, member *domain.Member) (*domain.Member, error)
	FindByEmail(ctx context.Context, email domain.Email) (*domain.Member, error)
}

func NewMemberRepository(client *ent.Client) MemberRepository {
	return &memberRepository{
		client: client,
	}
}

type memberRepository struct {
	client *ent.Client
}

func (m *memberRepository) Save(ctx context.Context, tx *ent.Tx, member *domain.Member) (*domain.Member, error) {
	savedMember, err := tx.Member.Create().
		SetEmail(member.Email.Address).
		SetNickname(member.Nickname).
		SetPasswordHash(member.PasswordHash).
		SetStatus(member.Status.String()).
		Save(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	email, err := domain.NewEmail(savedMember.Email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	status := domain.NewMemberStatus(savedMember.Status)

	return &domain.Member{
		ID:           savedMember.ID,
		Email:        email,
		Nickname:     savedMember.Nickname,
		PasswordHash: savedMember.PasswordHash,
		Status:       status,
	}, nil
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

func (m *memberRepository) FindByEmail(ctx context.Context, email domain.Email) (*domain.Member, error) {
	member, err := m.client.Member.Query().
		Where(member.EmailEQ(email.Address)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	domainEmail, err := domain.NewEmail(member.Email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	status := domain.NewMemberStatus(member.Status)

	return &domain.Member{
		ID:           member.ID,
		Email:        domainEmail,
		Nickname:     member.Nickname,
		PasswordHash: member.PasswordHash,
		Status:       status,
	}, nil
}
