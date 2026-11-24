🔒 CORS: защита и доступ между источниками
CORS (Cross-Origin Resource Sharing) — механизм, позволяет или запрещает веб-страницам делать запросы к другим доменам.

1. Заголовки CORS:
     Access-Control-Allow-Origin: https://example.com
   Access-Control-Allow-Methods: GET, POST, OPTIONS
   Access-Control-Allow-Headers: Content-Type
   Access-Control-Max-Age: 86400

2. Типы запросов:
   - Простые (GET, POST, HEAD)
   - Предварительные (OPTIONS)
   - Сложные (с дополнительными заголовками)

⚠️ Когда настраивать?
- API с разных доменов
- Микрофронтенды
- CDN-ресурсы
- Сторонние сервисы

- Настраивается на сервере
- Влияет на безопасность
- Может блокировать легитимные запросы
