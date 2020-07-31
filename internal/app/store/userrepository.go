package store

import "github.com/PaperDevil/goREST/internal/app/model"

// По сути абстрактая обёртка вокург БД, для реализации методов конкретной сущности
type UserRepository struct {
	store *Store
}

// Создать пользователя в БД
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow( // Запрос к БД
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil { // Scan() установит вернувшееся от RETURNING id значение в &u.ID
		// А в случае ошибок просто возвращаем их
		return nil, err
	}
	// Если всё проходит успешно, то возвращаем экз. пользователя
	return u, nil
}

// Поиск пользователя по полю email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}             // Создаём пустой шаблон, который заполним полученными с БД данными
	if err := r.store.db.QueryRow( // Снова выполняем запрос к БД
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		// В этот раз Scan() установит в пустой обьект u все значения из БД
		return nil, err
	}
	// Даже если в БД не найдётся юзер с таким email, мы получим пустой обьект без ошибок
	return u, nil
}
