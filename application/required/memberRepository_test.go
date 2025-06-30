package required

import (
	"context"
	"goplearn/domain"
	"goplearn/ent"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
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

func TestMemberRepository_Save(t *testing.T) {
	client := setupTestDB(t)
	defer client.Close()

	ctx := context.Background()
	repo := NewMemberRepository(client)

	member := domain.CreateTestMember(t)

	tx, err := client.Tx(ctx)
	require.NoError(t, err)
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	savedMember, err := repo.Save(ctx, tx, member)
	require.NoError(t, err)
	assert.NotNil(t, savedMember)
	assert.NotNil(t, savedMember.ID)
	assert.Equal(t, member.Email.Address, "kopher@goplearn.app")

	// 트랜잭션 커밋
	err = tx.Commit()
	require.NoError(t, err)
}

func TestMemberRepository_DuplicateEmailFail(t *testing.T) {
	client := setupTestDB(t)
	defer client.Close()

	ctx := context.Background()
	repo := NewMemberRepository(client)

	member := domain.CreateTestMember(t)

	// 첫 번째 멤버 저장
	tx1, err := client.Tx(ctx)
	require.NoError(t, err)

	savedMember, err := repo.Save(ctx, tx1, member)
	require.NoError(t, err)
	assert.NotNil(t, savedMember)

	// 첫 번째 트랜잭션 커밋
	err = tx1.Commit()
	require.NoError(t, err)

	// 두 번째 멤버 저장 (같은 이메일)
	member2 := domain.CreateTestMember(t)
	tx2, err := client.Tx(ctx)
	require.NoError(t, err)
	defer tx2.Rollback()

	_, err = repo.Save(ctx, tx2, member2)
	assert.Error(t, err)
	assert.True(t, ent.IsConstraintError(err))
}
