package domain

import (
	"strings"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

func CreateMockMemberRegisterRequest() *MemberRegisterRequest {
	return &MemberRegisterRequest{
		Email:    "kopher@goplearn.app",
		Nickname: "Kopher",
		Password: "secretpassword",
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

	member, err := RegisterMember(CreateMockMemberRegisterRequest(), mpe)
	if err != nil {
		t.Fatal(err)
	}
	return member

}
