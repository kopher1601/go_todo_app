package domain

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockPasswordEncoder struct{}

func (m *mockPasswordEncoder) Encode(password string) (string, error) {
	return strings.ToUpper(password), nil
}

func (m *mockPasswordEncoder) Matches(password, encodedPassword string) bool {
	return strings.ToUpper(password) == encodedPassword
}

func createTestMember(t *testing.T) *Member {
	t.Helper()
	member, err := CreateMember(&MemberCreateRequest{
		Email:    "kopher@goplearn.app",
		Nickname: "Kopher",
		Password: "secret",
	}, &mockPasswordEncoder{})
	if err != nil {
		t.Fatal(err)
	}
	return member
}

func TestMember_Status(t *testing.T) {
	t.Run("activate", func(t *testing.T) {
		member := createTestMember(t)
		member.Activate()

		assert.Equal(t, member.status, MemberStatusActive)
	})

	t.Run("activate_fail", func(t *testing.T) {
		member := createTestMember(t)
		member.Activate()
		err := member.Activate()

		if ok := assert.Error(t, err); ok {
			assert.Equal(t, err, ErrIllegalState)
		}
	})

	t.Run("deactivate", func(t *testing.T) {
		member := createTestMember(t)
		member.Activate()
		member.Deactivate()

		assert.Equal(t, member.status, MemberStatusDeactivated)
	})

	t.Run("deactivate_fail", func(t *testing.T) {
		member := createTestMember(t)
		err := member.Deactivate()

		if ok := assert.Error(t, err); ok {
			assert.Equal(t, err, ErrIllegalState)
		}
	})

	t.Run("deactivate_fail_twice", func(t *testing.T) {
		member := createTestMember(t)
		member.Deactivate()
		err := member.Deactivate()

		if ok := assert.Error(t, err); ok {
			assert.Equal(t, err, ErrIllegalState)
		}
	})

	t.Run("verify_password", func(t *testing.T) {
		member := createTestMember(t)
		member.Activate()
		result := member.VerifyPassword("secret", &mockPasswordEncoder{})
		assert.True(t, result)

		result = member.VerifyPassword("hello", &mockPasswordEncoder{})
		assert.False(t, result)
	})

	t.Run("change_nickname", func(t *testing.T) {
		member := createTestMember(t)

		assert.Equal(t, member.nickname, "Kopher")
		member.ChangeNickname("Koma")

		assert.Equal(t, member.nickname, "Koma")
	})

	t.Run("change_password", func(t *testing.T) {
		member := createTestMember(t)

		member.ChangePassword("verysecret", &mockPasswordEncoder{})

		assert.True(t, member.VerifyPassword("verysecret", &mockPasswordEncoder{}))
	})

	t.Run("is_active", func(t *testing.T) {
		member := createTestMember(t)

		member.Activate()
		assert.True(t, member.IsActive())

		member.Deactivate()
		assert.False(t, member.IsActive())
	})

}

func TestMemberVO(t *testing.T) {

	t.Run("invalid_email_1", func(t *testing.T) {
		member, err := CreateMember(&MemberCreateRequest{
			Email:    "test",
			Nickname: "Kopher",
			Password: "secret",
		}, &mockPasswordEncoder{})

		assert.ErrorIs(t, err, ErrInvalidEmail)
		assert.Nil(t, member)
	})

	t.Run("invalid_email_2", func(t *testing.T) {
		member := createTestMember(t)

		assert.ErrorIs(t, member.ChangeEmail("k-opher"), ErrInvalidEmail)

		assert.NoError(t, member.ChangeEmail("kopher@goplearn.app"))
		assert.Equal(t, member.email.address, "kopher@goplearn.app")
	})
}
