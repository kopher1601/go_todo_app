package provided_test

import (
	"context"
	"goplearn/application"
	"goplearn/application/required"
	"goplearn/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

func TestMemberFinder(t *testing.T) {
	ctrl := gomock.NewController(t)

	memberRepository := required.NewMockMemberRepository(ctrl)
	emailSender := required.NewMockEmailSender(ctrl)
	passwordEncoder := domain.NewMockPasswordEncoder(ctrl)

	memberRepository.EXPECT().
		FindByID(gomock.Any(), gomock.Any()).
		Return(&domain.Member{
			ID:       1,
			Email:    domain.Email{Address: "kopher@goplearn.app"},
			Nickname: "Kopher",
		}, nil)

	memberRegister := application.NewMemberRegister(
		nil,
		memberRepository,
		emailSender,
		passwordEncoder,
	)

	member, err := memberRegister.Find(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, member)
	assert.Equal(t, member.ID, 1)
	assert.Equal(t, member.Email.Address, "kopher@goplearn.app")
	assert.Equal(t, member.Nickname, "Kopher")
}
