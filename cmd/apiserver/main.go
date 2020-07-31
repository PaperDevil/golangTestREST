package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/PaperDevil/goREST/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

// init...
func init() {
	// Мета информация для скомпилированной программы
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config")
}

// main ...
func main() {
	flag.Parse() // Парсим переданные в командной строке флаги

	config := apiserver.NewConfig()               // Создаём конфиг
	_, err := toml.DecodeFile(configPath, config) // Раскидываем значение по структуре Config
	if err != nil {
		// Если в файле toml есть не совпадения типов, то ошибка
		log.Fatal(err)
	}
	s := apiserver.New(config) // Создаём экземпляр сервера
	if err := s.Start(); err != nil {
		// Если при запуске есть ошибка - выход
		log.Fatal(err)
	}
}
