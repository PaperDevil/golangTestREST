package apiserver

import "github.com/PaperDevil/goREST/internal/app/store"

// Описание полей конфига
type Config struct {
	// Приставка toml является флагом обработки для BurntSushi/toml
	BindAddr string        `toml:"bind_addr"`
	LogLevel string        `toml:"log_level"`
	Store    *store.Config // Конфигурация БД
}

// Создаём экземпляр конфига
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
