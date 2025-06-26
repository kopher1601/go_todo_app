package domain

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestMember_Status(t *testing.T) {
	t.Run("activate", func(t *testing.T) {
		member := CreateTestMember(t)
		_ = member.Activate()

		assert.Equal(t, member.Status, MemberStatusActive)
	})

	t.Run("activate_fail", func(t *testing.T) {
		member := CreateTestMember(t)

		_ = member.Activate()
		err := member.Activate()

		assert.ErrorIs(t, err, ErrIllegalState)
	})

	t.Run("deactivate", func(t *testing.T) {
		member := CreateTestMember(t)

		_ = member.Activate()
		err := member.Deactivate()

		assert.NoError(t, err)
		assert.Equal(t, member.Status, MemberStatusDeactivated)
	})

	t.Run("deactivate_fail", func(t *testing.T) {
		member := CreateTestMember(t)

		err := member.Deactivate()

		assert.ErrorIs(t, err, ErrIllegalState)
	})

	t.Run("deactivate_fail_twice", func(t *testing.T) {
		member := CreateTestMember(t)

		_ = member.Deactivate()
		err := member.Deactivate()

		assert.ErrorIs(t, err, ErrIllegalState)
	})

	t.Run("verify_password", func(t *testing.T) {
		member := CreateTestMember(t)
		member.Activate()

		ctrl := gomock.NewController(t)
		mpe := NewMockPasswordEncoder(ctrl)
		mpe.EXPECT().Matches(gomock.Any(), gomock.Any()).Return(true)

		result := member.VerifyPassword("SECRET", mpe)
		assert.True(t, result)

		mpe.EXPECT().Matches(gomock.Any(), gomock.Any()).Return(false)
		result = member.VerifyPassword("HELLO", mpe)
		assert.False(t, result)
	})

	t.Run("change_nickname", func(t *testing.T) {
		member := CreateTestMember(t)

		member.ChangeNickname("Koma")

		assert.Equal(t, member.Nickname, "Koma")
	})

	t.Run("change_password", func(t *testing.T) {
		member := CreateTestMember(t)

		ctrl := gomock.NewController(t)
		mpe := NewMockPasswordEncoder(ctrl)
		mpe.EXPECT().Encode(gomock.Any()).DoAndReturn(
			func(password string) (string, error) {
				return strings.ToUpper(password), nil
			})

		member.ChangePassword("verysecret", mpe)

		mpe.EXPECT().Matches(gomock.Any(), gomock.Any()).Return(true)
		assert.True(t, member.VerifyPassword("VERYSECRET", mpe))
	})

	t.Run("is_active", func(t *testing.T) {
		member := CreateTestMember(t)

		member.Activate()
		assert.True(t, member.IsActive())

		member.Deactivate()
		assert.False(t, member.IsActive())
	})

}

func TestMemberVO(t *testing.T) {

	t.Run("invalid_email_1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mpe := NewMockPasswordEncoder(ctrl)
		mpe.EXPECT().Encode(gomock.Any()).DoAndReturn(
			func(password string) (string, error) {
				return strings.ToUpper(password), nil
			})

		member, err := RegisterMember(&MemberRegisterRequest{
			Email:    "test",
			Nickname: "Kopher",
			Password: "secret",
		}, mpe)

		assert.ErrorIs(t, err, ErrInvalidEmail)
		assert.Nil(t, member)
	})

	t.Run("invalid_email_2", func(t *testing.T) {
		member := CreateTestMember(t)

		assert.ErrorIs(t, member.ChangeEmail("k-opher"), ErrInvalidEmail)

		assert.NoError(t, member.ChangeEmail("kopher@goplearn.app"))
		assert.Equal(t, member.Email.Address, "kopher@goplearn.app")
	})
}
