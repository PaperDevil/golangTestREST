package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

// Фикстура, которую мы вызываем перед тестами
func TestMain(m *testing.M) {
	// Устанавливаем тестовое значение БД-урла
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=testing_rest sslmode=disable user=postgres password=admin"
	}

	os.Exit(m.Run())
}
