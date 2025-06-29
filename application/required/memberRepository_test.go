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
	if err != nil {
		t.Fatal(err)
	}

	savedMember, err := repo.Save(ctx, tx, member)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, savedMember)
	assert.NotNil(t, savedMember.ID)
	assert.Equal(t, member.Email.Address, "kopher@goplearn.app")
}

func TestMemberRepository_DuplicateEmailFail(t *testing.T) {
	client := setupTestDB(t)
	defer client.Close()

	ctx := context.Background()
	repo := NewMemberRepository(client)

	member := domain.CreateTestMember(t)
	tx, err := client.Tx(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	_, err = repo.Save(ctx, tx, member)
	assert.NoError(t, err)

	member2 := domain.CreateTestMember(t)
	_, err = repo.Save(ctx, tx, member2)
	assert.Error(t, err)
	assert.True(t, ent.IsConstraintError(err))
}
