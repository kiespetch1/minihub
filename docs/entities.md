# Сущности MiniHub

## User (Пользователь)

Основная сущность для аутентификации.

**Поля:**
- `id` (UUID) — уникальный идентификатор пользователя
- `email` (string, unique) — email адрес для входа
- `password_hash` (string) — хеш пароля (bcrypt)
- `created_at` (timestamp) — дата создания аккаунта
- `updated_at` (timestamp) — дата последнего обновления
- `timezone` (string) — временная зона (IANA, например "Europe/Moscow")

**Связи:**
- 1:N с Bookmark
- 1:N с Collection
- 1:N с Tag

---

## Bookmark (Закладка)

Сохраненная ссылка с метаданными.

**Поля:**
- `id` (UUID) — уникальный идентификатор
- `user_id` (UUID, FK) — владелец закладки
- `url` (string) — сохраненный URL
- `title` (string) — название закладки
- `description` (string, nullable) — описание/заметка
- `collection_id` (UUID, FK, nullable) — коллекция (опционально)
- `tags` ([]string) — массив тегов (ID)
- `created_at` (timestamp) — дата создания
- `updated_at` (timestamp) — дата последнего обновления

**Связи:**
- N:1 с User
- N:1 с Collection (опционально)
- M:N с Tag

---

## Collection (Коллекция)

Группировка закладок по темам.

**Поля:**
- `id` (UUID) — уникальный идентификатор
- `user_id` (UUID, FK) — владелец коллекции
- `name` (string) — название коллекции
- `color` (string) — цвет в формате hex (например "#FF5733")
- `created_at` (timestamp) — дата создания
- `updated_at` (timestamp) — дата последнего обновления

**Связи:**
- N:1 с User
- 1:N с Bookmark

---

## Tag (Тег)

Метка для категоризации закладок.

**Поля:**
- `id` (UUID) — уникальный идентификатор
- `user_id` (UUID, FK) — владелец тега (теги приватные)
- `name` (string) — название тега
- `created_at` (timestamp) — дата создания

**Связи:**
- N:1 с User
- M:N с Bookmark

**Примечание:** Теги привязаны к конкретному пользователю. При удалении тега связи с закладками также удаляются.