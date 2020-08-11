package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// Файл служит для описания моделей БД в виде языковых структур

// Таблица users в БД
type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

// Валидация входных значений юзера
func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
	)
}

// Функция вызываемая перед созданием пользователя в БД
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 { // Проверяем, что у Юзера есть пароль
		enc, err := encryptString(u.Password) // Шифруем пароль
		if err != nil {
			// Выкидываем ошибки
			return err
		}

		u.EncryptedPassword = enc // Устанавливаем зашифрованный пароль
	}
	return nil
}

// Шифруем пароль
func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
