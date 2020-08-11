package model

import "testing"

// Фикстура для тестов юзера
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "examplepassword",
	}
}
