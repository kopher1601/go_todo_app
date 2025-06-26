package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	email1, err := NewEmail("test1@test.com")
	assert.NoError(t, err)

	email2, err := NewEmail("test1@test.com")
	assert.NoError(t, err)

	t.Run("Email equality", func(t *testing.T) {
		assert.Equal(t, email1.Address, email2.Address)
		assert.Equal(t, email1, email2)
	})
}
