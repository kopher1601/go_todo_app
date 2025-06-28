package provided_test

import (
	"context"
	"goplearn/application"
	"goplearn/application/required"
	"goplearn/domain"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestMemberRegister_MockGen(t *testing.T) {
	ctrl := gomock.NewController(t)
	mrm := required.NewMockMemberRepository(ctrl)
	mes := required.NewMockEmailSender(ctrl)
	mpe := domain.NewMockPasswordEncoder(ctrl)

	mrm.EXPECT().
		FindByEmail(gomock.Any(), gomock.AssignableToTypeOf(domain.Email{})).
		Return(nil, nil)

	mrm.EXPECT().
		Save(gomock.Any(), gomock.AssignableToTypeOf(&domain.Member{})).
		DoAndReturn(
			func(ctx context.Context, member *domain.Member) (*domain.Member, error) {
				member.SetID(1)
				return member, nil
			})

	mes.EXPECT().
		Send(gomock.Any(), gomock.AssignableToTypeOf(domain.Email{}), gomock.Any(), gomock.Any()).
		Times(1).
		Return(nil)

	mpe.EXPECT().
		Encode(gomock.Any()).
		DoAndReturn(
			func(password string) (string, error) {
				return strings.ToUpper(password), nil
			})

	memberRegister := application.NewMemberRegister(
		mrm,
		mes,
		mpe,
	)

	member, err := memberRegister.Register(context.Background(), domain.CreateMockMemberRegisterRequest())
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, member)
	assert.Equal(t, member.ID, 1)
	assert.Equal(t, member.Email.Address, "kopher@goplearn.app")
	assert.Equal(t, member.Nickname, "Kopher")
	assert.Equal(t, member.PasswordHash, "SECRETPASSWORD")
	assert.Equal(t, member.Status, domain.MemberStatusPending)
}

func TestMemberRegister_DuplicateEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	mrm := required.NewMockMemberRepository(ctrl)
	mes := required.NewMockEmailSender(ctrl)
	mpe := domain.NewMockPasswordEncoder(ctrl)

	existingMember := &domain.Member{
		ID:           1,
		Email:        domain.Email{Address: "kopher@goplearn.app"},
		Nickname:     "ExistingKopher",
		PasswordHash: "EXISTINGPASSWORD",
		Status:       domain.MemberStatusActive,
	}

	mrm.EXPECT().
		FindByEmail(gomock.Any(), gomock.AssignableToTypeOf(domain.Email{})).
		Return(existingMember, nil)

	memberRegister := application.NewMemberRegister(
		mrm,
		mes,
		mpe,
	)

	member, err := memberRegister.Register(context.Background(), domain.CreateMockMemberRegisterRequest())

	assert.ErrorIs(t, err, domain.ErrDuplicateEmail)
	assert.Nil(t, member)
}
