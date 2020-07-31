package model

// Файл служит для описания моделей БД в виде языковых структур

// Таблица users в БД
type User struct {
	ID                int
	Email             string
	EncryptedPassword string
}
