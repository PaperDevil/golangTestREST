package apiserver

// Описание полей конфига
type Config struct {
	// Приставка toml является флагом обработки для BurntSushi/toml
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

// Создаём экземпляр конфига
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
