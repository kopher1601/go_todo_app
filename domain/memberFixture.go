package domain

import (
	"strings"
	"testing"
)

type mockPasswordEncoder struct{}

func NewMockPasswordEncoder() *mockPasswordEncoder {
	return &mockPasswordEncoder{}
}

func (m *mockPasswordEncoder) Encode(password string) (string, error) {
	return strings.ToUpper(password), nil
}

func (m *mockPasswordEncoder) Matches(password, encodedPassword string) bool {
	return strings.ToUpper(password) == encodedPassword
}

func CreateMockMemgerRegisterRequest() *MemberRegisterRequest {
	return &MemberRegisterRequest{
		Email:    "kopher@goplearn.app",
		Nickname: "Kopher",
		Password: "secret",
	}
}

func CreateTestMember(t *testing.T) *Member {
	t.Helper()
	member, err := RegisterMember(CreateMockMemgerRegisterRequest(), NewMockPasswordEncoder())
	if err != nil {
		t.Fatal(err)
	}
	return member

}
