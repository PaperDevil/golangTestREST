# Тест REST API на GoLang
Тестовый проект на ЯП GoLang реализующий REST API функциональность.
Разработано в соответствии с видеорядом -> https://www.youtube.com/watch?v=LxJLuW5aUDQ&t

# Зависимости
 - github.com/BurntSushi/toml v0.3.1
 - github.com/gorilla/mux v1.7.4
 - github.com/sirupsen/logrus v1.6.0
 - github.com/stretchr/testify v1.6.1
 
# Запуск
```
>> make build
>> ./apiserver
```

# Тестирование
Исходники лежат в ``./internal/app/apiserver``
```
>> make test
or
>> go test -v -race -timeout 30s ./...
```

# Лицензии
Этот проект лицензирован в соответствии с условиями лицензии MIT.