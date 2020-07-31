package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Тестируем ручку /hello
func TestAPIServer_HandleHello(t *testing.T) {
	s := New(NewConfig())                                    // Создаём APIServer
	rec := httptest.NewRecorder()                            // Создаём тестер
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil) // Создаём поле http запроса для ручки
	// Вызываем метод роута /hello передавая ему тестер (как сервер) и тестовый http запрос
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, "Hello", rec.Body.String()) // Ассертим, что ответ равен "Hello"
}
