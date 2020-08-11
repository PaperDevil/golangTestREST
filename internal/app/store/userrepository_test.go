package store_test

import (
	"github.com/PaperDevil/goREST/internal/app/model"
	"github.com/PaperDevil/goREST/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Проверка корректности создания пользователя в БД
func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

// Проверка корректности поиска пользователя по email
func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	// Для ещё не существующего User
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	// Для созданного User
	s.User().Create(model.TestUser(t))
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
