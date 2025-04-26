package application

// TODO: Реализовать бизнес-логику аутентификации
//  - Метод RegisterUser(ctx, email, password) (создание пользователя, хеширование пароля, запись в БД, публикация события в Kafka)
//  - Метод Login(ctx, email, password) (проверка пароля, выдача токенов)
//  - Метод RefreshTokens(ctx, refreshToken) (обновление токенов)
//  - Использовать интерфейсы Repository, TokenManager, SessionStore

type Service struct {
	// зависимости сюда через конструктор
}
