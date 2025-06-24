package domain

import (
	"strings"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

// type mockPasswordEncoder struct{}

// func NewMockPasswordEncoder() *mockPasswordEncoder {
// 	return &mockPasswordEncoder{}
// }

// func (m *mockPasswordEncoder) Encode(password string) (string, error) {
// 	return strings.ToUpper(password), nil
// }

// func (m *mockPasswordEncoder) Matches(password, encodedPassword string) bool {
// 	return strings.ToUpper(password) == encodedPassword
// }

func CreateMockMemgerRegisterRequest() *MemberRegisterRequest {
	return &MemberRegisterRequest{
		Email:    "kopher@goplearn.app",
		Nickname: "Kopher",
		Password: "secret",
	}
}

func CreateTestMember(t *testing.T) *Member {
	t.Helper()

	ctrl := gomock.NewController(t)
	mpe := NewMockPasswordEncoder(ctrl)
	mpe.EXPECT().Encode(gomock.Any()).DoAndReturn(
		func(password string) (string, error) {
			return strings.ToUpper(password), nil
		})

	member, err := RegisterMember(CreateMockMemgerRegisterRequest(), mpe)
	if err != nil {
		t.Fatal(err)
	}
	return member

}
