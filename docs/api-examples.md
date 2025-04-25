# 📡 Примеры API запросов

Ниже приведены примеры базовых REST API запросов к системе Cloud Storage Platform через API Gateway.

---

## 🔐 Аутентификация (Auth Service)

### Регистрация пользователя

**Запрос:**

```bash
POST http://localhost:8080/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "StrongPassword123",
  "username": "newuser"
}
```

**Ответ:**

```bash
{
  "access_token": "ACCESS_TOKEN",
  "refresh_token": "REFRESH_TOKEN"
}
```
---

### Логин пользователя

**Запрос:**

```bash
POST http://localhost:8080/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "StrongPassword123"
}
```

**Ответ:**

```bash
{
  "access_token": "ACCESS_TOKEN",
  "refresh_token": "REFRESH_TOKEN"
}
```
---

## 👤 Работа с профилем (User Service)

### Получение профиля пользователя

**Запрос:**

```bash
GET http://localhost:8080/user/profile
Authorization: Bearer ACCESS_TOKEN
```

**Ответ:**

```bash
{
  "id": "UUID",
  "email": "user@example.com",
  "username": "newuser",
  "created_at": "2025-04-25T10:00:00Z"
}
```

## Обновление профиля пользователя

**Запрос:**

```bash
PATCH http://localhost:8080/user/profile
Authorization: Bearer ACCESS_TOKEN
Content-Type: application/json

{
  "username": "updateduser",
  "email": "updateduser@example.com"
}
```

**Ответ:**

```bash
{
  "message": "Profile updated successfully."
}
```
---

## 🗂️ Работа с файлами (Storage Service)

### Загрузка файла

**Запрос:**

```bash
POST http://localhost:8080/storage/upload
Authorization: Bearer ACCESS_TOKEN
Content-Type: multipart/form-data

Form-Data:
file=@path/to/your/file.txt
```

**Ответ:**

```bash
{
  "file_id": "UUID",
  "file_url": "http://localhost:8080/storage/files/UUID"
}
```

### Скачивание файла

**Запрос:**

```bash
GET http://localhost:8080/storage/files/{file_id}
Authorization: Bearer ACCESS_TOKEN
```

**Ответ:**
```bash
<Файл будет загружен>
```

##Получение списка файлов пользователя

**Запрос:**

```bash
GET http://localhost:8080/storage/list
Authorization: Bearer ACCESS_TOKEN
```

**Ответ:**
```bash
[
  {
    "file_id": "UUID",
    "file_name": "file.txt",
    "uploaded_at": "2025-04-25T10:10:00Z",
    "access_level": "private"
  },
  {
    "file_id": "UUID",
    "file_name": "image.png",
    "uploaded_at": "2025-04-25T10:20:00Z",
    "access_level": "public"
  }
]
```