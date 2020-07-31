package apiserver

import (
	"github.com/PaperDevil/goREST/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// Описание полей обьекта сервера
type APIServer struct {
	config *Config        // Наша структура Config
	logger *logrus.Logger // Логер
	router *mux.Router    // Роутер
	store  *store.Store   // Хранилище БД
}

// Инициализируем сервер
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Запускам сервер
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		// Пытаемся сконфигурировать наш логер
		return err
	}
	s.configgerRouter() // Устанавливаем маршруты
	if err := s.configureStore(); err != nil {
		// Конфигурирование БД
		return err
	}
	// Выводим лог, о начале работы сервера
	s.logger.Info("Starting API Server")
	// Запускаем прослушку по адресу из конфига
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// Конфигурирования логера logrus
func (s *APIServer) configureLogger() error {
	// Проверяем, что в конфиге подходящее имя уровня
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		// Ошибку выкидываем
		return err
	}
	// Если ОК, то устанавливаем уровень
	s.logger.SetLevel(level)
	return nil
}

// Конфигурация роутера gorilla/mux
func (s *APIServer) configgerRouter() {
	// По типу s.route.HandleFunc("/route", func())
	s.router.HandleFunc("/hello", s.handleHello())
}

// Конфигурация хранилища БД
func (s *APIServer) configureStore() error {
	// var st *store.Store = store.New(s.config.Store)
	st := store.New(s.config.Store) // Создаём экземпляр БД-коннектора
	if err := st.Open(); err != nil {
		// В случае ошибки соединения с БД
		return err
	}
	s.store = st // Устанавливаем st как хранилище для сервера
	return nil
}

// Роут /hello
func (s *APIServer) handleHello() http.HandlerFunc {
	// Возвращаемый тип http.HandleFunc т.к. можно сделать какую-то обработку в этой функции
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
