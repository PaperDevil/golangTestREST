package store

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
)

// Структура описывающая хранилище БД
type Store struct {
	config         *Config
	db             *sql.DB         // Инстанс БД драйвера
	userRepository *UserRepository // Абстрактное представление хранилища юзеров
}

// Фабрика инстансов Store
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Открыть соединение
func (s *Store) Open() error {
	// Открываем соединение с SQLite-файлом по пути из конфига
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		// Опри ошибке соеднинения
		return err
	}
	if err := db.Ping(); err != nil {
		// Если запросы к БД не проходят
		return err
	}
	s.db = db // Делаем инстанс открытой БД основным для APIServer
	return nil
}

// Закрыть соединение
func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		// В случае, если у нас уже есть реп
		return s.userRepository
	}
	// Создаём новый абстрактный репозиторий
	s.userRepository = &UserRepository{
		s,
	}
	return s.userRepository
}
