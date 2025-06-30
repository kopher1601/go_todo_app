package provided_test

import (
	"context"
	"goplearn/application"
	"goplearn/application/required"
	"goplearn/domain"
	"goplearn/ent"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
	_ "gorm.io/driver/sqlite"
)

func setupTestDB(t *testing.T) *ent.Client {
	// 테스트용 데이터베이스 설정
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)

	// 스키마 생성
	err = client.Schema.Create(context.Background())
	require.NoError(t, err)

	return client
}

func TestMemberRegister_MockGen(t *testing.T) {
	client := setupTestDB(t)
	defer client.Close()

	ctrl := gomock.NewController(t)
	mrm := required.NewMockMemberRepository(ctrl)
	mes := required.NewMockEmailSender(ctrl)
	mpe := domain.NewMockPasswordEncoder(ctrl)

	mrm.EXPECT().
		FindByEmail(gomock.Any(), gomock.AssignableToTypeOf(domain.Email{})).
		Return(nil, nil)

	mrm.EXPECT().
		Save(gomock.Any(), gomock.AssignableToTypeOf(&ent.Tx{}), gomock.AssignableToTypeOf(&domain.Member{})).
		DoAndReturn(
			func(ctx context.Context, tx *ent.Tx, member *domain.Member) (*domain.Member, error) {
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
				log.Println("password >>> ", password)
				return strings.ToUpper(password), nil
			})

	memberRegister := application.NewMemberRegister(
		client,
		mrm,
		mes,
		mpe,
	)

	member, err := memberRegister.Register(context.Background(), domain.CreateMockMemberRegisterRequest())
	if err != nil {
		t.Logf("Register error: %v", err)
		t.Fatal(err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, member)
	if member != nil {
		assert.Equal(t, member.ID, 1)
		assert.Equal(t, member.Email.Address, "kopher@goplearn.app")
		assert.Equal(t, member.Nickname, "Kopher")
		assert.Equal(t, member.PasswordHash, "SECRETPASSWORD")
		assert.Equal(t, member.Status, domain.MemberStatusPending)
	}
}

func TestMemberRegister_DuplicateEmail(t *testing.T) {
	client := setupTestDB(t)
	defer client.Close()

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
		client,
		mrm,
		mes,
		mpe,
	)

	member, err := memberRegister.Register(context.Background(), domain.CreateMockMemberRegisterRequest())

	assert.ErrorIs(t, err, domain.ErrDuplicateEmail)
	assert.Nil(t, member)
}
