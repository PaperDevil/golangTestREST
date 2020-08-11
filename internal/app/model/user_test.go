package model_test

import (
	"github.com/PaperDevil/goREST/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Проверка работоспособности валидатора
func TestUser_Validate(t *testing.T) {
	tesCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
	}
	for _, tc := range tesCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

// Тестирование функции BeforeCreate
func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)                  // Создаём тестового юзера
	assert.NoError(t, u.BeforeCreate())     // Проверяем на наличие ошибок
	assert.NotEmpty(t, u.EncryptedPassword) // Проверяем, что пароль шифруется
}
