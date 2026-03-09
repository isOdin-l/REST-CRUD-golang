# REST API сервер

[![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat-square&logo=go&logoColor=white)](https://golang.org/)
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat-square&logo=docker&logoColor=white)](https://www.docker.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-18.1-336791?style=flat-square&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![JWT](https://img.shields.io/badge/JWT-black?style=flat-square&logo=JSON%20web%20tokens)](https://www.jwt.io/introduction)
[![Echo](https://img.shields.io/badge/Echo-v5-00ADD8?style=flat-square&logo=go&logoColor=white)](https://echo.labstack.com/)
[![Clean Architecture](https://img.shields.io/badge/Clean_Architecture-Implemented-success?style=flat-square)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

Целью данного проекта было научиться разработке REST API сервера на Go, следуя подходу чистой архитектуры в проекте.

### Основные моменты, которые удалось реализовать:
- **Безопасность** - JWT аутентификация, хэширование паролей, чувствительные данные подаются через .env файлы
- **Валидация** - данные проверяются на уровне handler
- **Работа с БД** - использовалась PostgreSQL с библиотекой squirrel для создания sql запросов. Также имеется поддержка транзакций на уровне service
- **CRUD** - реализованы основные CRUD операции сервиса
- **Чистая архитектура** - Строгое разделение задач на уровни domain, application и infrastructure
- **Контейнеризация** - Docker с multi-stage builds и health checks
- **Структура репозитория** - использовался стандарт [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- **Dependency Injection** - Interface-based дизайн для удобства в тестировании
- **Обработка ошибок** - Кастомные ошибки

## Стэк

| Component | Technology |
|-----------|------------|
| **Language** | Go 1.26 |
| **Framework** | [Echo](https://github.com/labstack/echo) | 
| **Database (PostgreSQL)** | [jackc/pgx](https://github.com/jackc/pgx) |
| **Query Builder** | [Masterminds/squirrel](https://github.com/Masterminds/squirrel) |
| **Authentication** | [golang-jwt/jwt](https://github.com/golang-jwt/jwt) |
| **Validation** | [go-playground/validator](https://github.com/go-playground/validator) |
| **Containerization** | Docker |
| **Configuration** | [caarlos0/env](https://github.com/caarlos0/env)  |


## Структура проекта

```

cmd/rest-api-server/            # Запуск приложения
configs/                        # Структура конфига
internal/
    server/                     # Роутинг и запуск сервера                     
    entities/                   # Модели для бизнес-логики
    api/                        # mappers для преобразования API<->entit
    repository/                 # Слой репозитория
    service/                    # Слой бизнес-логики
    handler/                    # Хэндлер на echo
    middleware/                 # Кастомные middleware (jwt)
    database/                   # Работа с БД - postgres, sql builder, названия таблиц
    errors/                     # Кастомные ошибки
pkg/api/                        # API модели 
migrations/                     # схема БД
```

## API Эндпоинты

#### Start point: `api/v1/`

### Authentication
- `POST /auth/sign-up` - Создание пользователя
- `POST /auth/sign-in` - Вход в аккаунт

### List
- `POST /list/` - Создание list
- `GET /list/:list_id` - Получение информации о list по id
- `PATCH /list/:list_id` - Обновление определённых полей list по id
- `DELETE /list/:list_id` - Удаление list по id

### Item
- `POST /list/:list_id/item/` - Создание item
- `GET /list/:list_id/item/:item_id` - Получение информации о item
- `PATCH /list/:list_id/item/:item_id` - Обновление определённых полей item
- `DELETE /list/:list_id/item/:item_id` - Удаление item


## Запуск сервера
* Клонирование репозитория
```bash
git clone https://github.com/isOdin-l/REST-CRUD-golang
cd REST-CRUD-golang
```

* Сборка и запуск в docker контейнерах

```bash
docker-compose up --build
```

## Special thanks
Огромная благодарность [@TobbyMax](https://github.com/TobbyMax) за ценные советы, review и помощь в реализации проекта.