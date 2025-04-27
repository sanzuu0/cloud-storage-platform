package query

// TODO: Структура запроса на логин пользователя
//  - Email
//  - Password
//  - Для передачи данных в Service.Login

type LoginQuery struct {
	Email    string
	Password string
}
