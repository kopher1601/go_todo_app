// Code generated by MockGen. DO NOT EDIT.
// Source: memberRepository.go
//
// Generated by this command:
//
//	mockgen -source=memberRepository.go -destination=mockMemberRepository.go -package required
//

// Package required is a generated GoMock package.
package required

import (
	context "context"
	domain "goplearn/domain"
	ent "goplearn/ent"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMemberRepository is a mock of MemberRepository interface.
type MockMemberRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMemberRepositoryMockRecorder
	isgomock struct{}
}

// MockMemberRepositoryMockRecorder is the mock recorder for MockMemberRepository.
type MockMemberRepositoryMockRecorder struct {
	mock *MockMemberRepository
}

// NewMockMemberRepository creates a new mock instance.
func NewMockMemberRepository(ctrl *gomock.Controller) *MockMemberRepository {
	mock := &MockMemberRepository{ctrl: ctrl}
	mock.recorder = &MockMemberRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMemberRepository) EXPECT() *MockMemberRepositoryMockRecorder {
	return m.recorder
}

// FindByEmail mocks base method.
func (m *MockMemberRepository) FindByEmail(ctx context.Context, email domain.Email) (*domain.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", ctx, email)
	ret0, _ := ret[0].(*domain.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockMemberRepositoryMockRecorder) FindByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockMemberRepository)(nil).FindByEmail), ctx, email)
}

// FindByID mocks base method.
func (m *MockMemberRepository) FindByID(ctx context.Context, memberId int) (*domain.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, memberId)
	ret0, _ := ret[0].(*domain.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockMemberRepositoryMockRecorder) FindByID(ctx, memberId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockMemberRepository)(nil).FindByID), ctx, memberId)
}

// Save mocks base method.
func (m *MockMemberRepository) Save(ctx context.Context, tx *ent.Tx, member *domain.Member) (*domain.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, tx, member)
	ret0, _ := ret[0].(*domain.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockMemberRepositoryMockRecorder) Save(ctx, tx, member any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockMemberRepository)(nil).Save), ctx, tx, member)
}

// Update mocks base method.
func (m *MockMemberRepository) Update(ctx context.Context, member *domain.Member) (*domain.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, member)
	ret0, _ := ret[0].(*domain.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMemberRepositoryMockRecorder) Update(ctx, member any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMemberRepository)(nil).Update), ctx, member)
}
