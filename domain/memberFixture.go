package domain

import (
	"strings"
	"testing"
)

type MockPasswordEncoder struct{}

func (m *MockPasswordEncoder) Encode(password string) (string, error) {
	return strings.ToUpper(password), nil
}

func (m *MockPasswordEncoder) Matches(password, encodedPassword string) bool {
	return strings.ToUpper(password) == encodedPassword
}

func CreateTestMember(t *testing.T) *Member {
	t.Helper()
	member, err := RegisterMember(&MemberRegisterRequest{
		Email:    "kopher@goplearn.app",
		Nickname: "Kopher",
		Password: "secret",
	}, &MockPasswordEncoder{})
	if err != nil {
		t.Fatal(err)
	}
	return member

}
