# API для работы с сообщениями

Простой REST API сервер для управления сообщениями, построенный с использованием Go, Echo framework и PostgreSQL.

## Технологии

- Go
- Echo framework
- GORM
- PostgreSQL

## Требования

- Go 1.x
- PostgreSQL
- Настроенная база данных PostgreSQL

## Настройка базы данных

Убедитесь, что у вас установлен и запущен PostgreSQL сервер со следующими параметрами:
```
host: localhost
user: postgres
password: 123
dbname: postgres
port: 5432
```

## Запуск проекта

1. Клонируйте репозиторий
2. Установите зависимости:
```bash
go mod download
```
3. Запустите сервер:
```bash
go run cmd/main.go
```

Сервер будет доступен по адресу: `http://localhost:8080`

## API Endpoints

### GET /messages
Получить все сообщения

**Ответ:**
```json
[
    {
        "id": 1,
        "text": "Текст сообщения"
    }
]
```

### POST /messages
Создать новое сообщение

**Тело запроса:**
```json
{
    "text": "Текст сообщения"
}
```

**Ответ:**
```json
{
    "status": "Success",
    "message": "Message have been added succsesfuly"
}
```

### PATCH /messages/:id
Обновить существующее сообщение

**Тело запроса:**
```json
{
    "text": "Новый текст сообщения"
}
```

**Ответ:**
```json
{
    "status": "Success",
    "message": "Message was updated"
}
```

### DELETE /messages/:id
Удалить сообщение

**Ответ:**
```json
{
    "status": "Success",
    "message": "Message was deleted"
}
```

## Коды ответов

- 200: Успешное выполнение операции
- 400: Ошибка в запросе или при обработке данных 