🔌 WebSocket: двунаправленная связь в реальном времени

❓ Что такое WebSocket?

WebSocket - это протокол, обеспечивающий полнодуплексную связь между клиентом и сервером поверх TCP-соединения.

💡 Основные концепции:

1. Установка соединения:
   - Handshake через HTTP
   - Upgrade до WebSocket
   - Постоянное соединение

2. Особенности протокола:
   - Двунаправленная передача данных
   - Минимальные накладные расходы
   - Поддержка сообщений и фреймов

⚠️ Когда использовать?

- Чаты и мессенджеры
- Онлайн-игры
- Системы мониторинга
- Торговые платформы

❗Важные моменты:

- Нагрузка на сервер при большом количестве соединений
- Необходимость обработки разрывов соединения
- Сложность масштабирования

---

markdown
### 1. «Чем WebSocket отличается от HTTP?»

#### { Короткий ответ }

HTTP — запрос-ответ: клиент спрашивает, сервер отвечает, соединение можно закрыть. WebSocket — persistent connection: после handshake соединение остаётся открытым, данные идут в обе стороны одновременно (full-duplex). Сервер может отправить данные клиенту без запроса.

#### { Развёрнутый ответ }

HTTP работает по модели запрос-ответ. Клиент отправляет запрос, сервер отвечает. Для real-time приходится использовать polling — клиент спрашивает каждые N секунд. Это нагружает сервер бессмысленными запросами.

WebSocket начинается как HTTP-запрос (Upgrade handshake), но потом протокол переключается. Соединение остаётся открытым. Данные идут в обе стороны — клиент может отправить сообщение серверу, а сервер может push-нуть данные клиенту в любой момент.

Важные отличия:
- HTTP — полудуплекс (запрос-ответ), WebSocket — полный дуплекс (одновременно)
- HTTP — новое соединение на каждый запрос (или keep-alive), WebSocket — одно соединение
- HTTP — overhead заголовков на каждый запрос (~800 байт), WebSocket — фрейм 2-6 байт
- WebSocket использует порты 80/443, проходит через прокси и фаерволы

#### { Пример кода }

```http
// HTTP polling — клиент дёргает каждые 3 секунды
// GET /api/messages?chat_id=1 -> 200 OK (даже если нет новых)
// GET /api/messages?chat_id=1 -> 200 OK
// GET /api/messages?chat_id=1 -> 200 OK (новое сообщение!)

// WebSocket — сервер сам присылает когда есть данные
// Client <-> Server: persistent connection
// Server -> Client: {type: "new_message", ...} (мгновенно)

// Handshake:
GET /ws HTTP/1.1
Upgrade: websocket
Connection: Upgrade

HTTP/1.1 101 Switching Protocols
Upgrade: websocket
```

#### { Типичная ошибка }

Говорить что WebSocket — это "быстрый HTTP" или "HTTP без задержки". Это другой протокол поверх TCP, с другой моделью коммуникации.

#### { Интервьюер хочет услышать }

Понимание модели коммуникации (full-duplex vs request-response), знание про handshake через HTTP Upgrade, и конкретный use-case — почему для мессенджера WebSocket, а не polling.

### 2. «Как обработать 10К одновременных WebSocket-соединений?»

#### { Короткий ответ }

В Go каждое WebSocket-соединение обслуживается двумя горутинами (read + write). 10К соединений = 20К горутин. Горутина весит ~2-8 КБ, итого ~40-160 МБ RAM. Go runtime использует epoll/kqueue под капотом, что эффективно для I/O-bound задач. 10К — не проблема для одного сервера.

#### { Развёрнутый ответ }

Go отлично подходит для большого количества WebSocket-соединений благодаря модели goroutine-per-connection.

Почему это работает:
1. Горутины легковесные — стартовый стек 2 КБ (растёт по необходимости до ~8 КБ)
2. Go runtime мультиплексирует горутины на OS-треды через M:N scheduling
3. Под капотом — epoll (Linux) / kqueue (macOS) для I/O multiplexing
4. Netpoller встроен в runtime — блокирующий Read() в горутине не блокирует OS-тред

Математика:
- 10К соединений x 2 горутины = 20К горутин
- 20К x 8 КБ = 160 МБ RAM (worst case)
- Плюс буферы: send channel на каждого клиента

Ограничения:
- File descriptors — увеличить ulimit (по умолчанию 1024)
- Один Hub с map и select — может стать bottleneck при 50К+
- Решение: шардирование Hub по chat_id или user_id % N

#### { Пример кода }

```go
// Каждый клиент = 2 горутины
func (c *Client) ReadPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		c.hub.broadcast <- message
	}
}

func (c *Client) WritePump() {
	for msg := range c.send {
		err := c.conn.WriteMessage(
			websocket.TextMessage,
			msg,
		)
		if err != nil {
			break
		}
	}
}

// 10К соединений:
// ulimit -n 65535
// 20К горутин ~ 160 МБ RAM
```

#### { Типичная ошибка }

Говорить "нужен thread pool" или "один поток на соединение". В Go горутины — не потоки, и goroutine-per-connection — это стандартный и правильный подход.

#### { Интервьюер хочет услышать }

Конкретные цифры (размер горутины, итого RAM), знание про epoll/netpoller, и понимание где bottleneck (Hub, file descriptors). Бонус — упомянуть шардирование Hub.

### 3. «Что такое Hub-паттерн и зачем он нужен?»

#### { Короткий ответ }

Hub — центральный диспетчер, который управляет WebSocket-соединениями. Хранит map подключённых клиентов, обрабатывает регистрацию/отключение/отправку сообщений. Все операции с map идут в одной горутине через select — channel-based synchronization, никаких мьютексов.

#### { Развёрнутый ответ }

Hub решает две задачи:
1. Управление соединениями — кто подключён, кто отключился
2. Маршрутизация сообщений — кому отправить конкретное сообщение

Почему именно каналы, а не мьютексы:
- Map в Go не thread-safe — нужна синхронизация
- Можно sync.RWMutex — но получаем lock contention при высокой нагрузке
- Channel + select + единственная горутина — все мутации map происходят последовательно
- Каналы в Go thread-safe по определению — можно писать из любой горутины
- Этот подход называется "share memory by communicating" — идиоматический Go

Три канала Hub:
- register — новый клиент подключился
- unregister — клиент отключился
- broadcast — сообщение для доставки

Метод Run() работает в отдельной горутине, бесконечный select по трём каналам. Это единственное место, где мутируется map.

#### { Пример кода }

```go
type Hub struct {
	clients    map[int64]*Client // user_id -> client
	register   chan *Client
	unregister chan *Client
	broadcast  chan *Message
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.userID] = client

		case client := <-h.unregister:
			if _, ok := h.clients[client.userID]; ok {
				delete(h.clients, client.userID)
				close(client.send)
			}

		case msg := <-h.broadcast:
			for _, uid := range msg.RecipientIDs {
				if client, ok := h.clients[uid]; ok {
					select {
					case client.send <- msg.Data:
					default:
						// буфер полон — клиент тормозит
						close(client.send)
						delete(h.clients, uid)
					}
				}
			}
		}
	}
}
```

#### { Типичная ошибка }

Использовать sync.Mutex для map и получить deadlock при высокой нагрузке. Или broadcast всем клиентам вместо адресной доставки по user_id — при 10К соединений это O(N) на каждое сообщение.

#### { Интервьюер хочет услышать }

Понимание почему каналы, а не мьютексы (идиоматический Go). Знание про default в select (защита от slow consumer). И объяснение почему единственная горутина владеет map — нет race condition по определению.

### 4. «Как реализовать graceful shutdown для WebSocket?»

#### { Короткий ответ }

При получении SIGTERM/SIGINT: 1) перестать принимать новые соединения, 2) отправить каждому клиенту Close-фрейм, 3) дождаться завершения всех горутин (WaitGroup или таймаут), 4) закрыть HTTP-сервер и соединения с БД.

#### { Развёрнутый ответ }

Graceful shutdown для WebSocket сложнее, чем для обычного HTTP, потому что у нас долгоживущие соединения.

Порядок:
1. Перехватить сигнал (SIGTERM при деплое, SIGINT при Ctrl+C)
2. Вызвать http.Server.Shutdown(ctx) — перестаёт принимать новые подключения
3. Для каждого активного WebSocket-клиента:
   - Отправить Close-фрейм (websocket.CloseMessage)
   - Клиент получает Close, ReadPump завершается
   - WritePump завершается когда закроется канал send
4. Дождаться завершения всех горутин (sync.WaitGroup)
5. Закрыть соединения с БД (pool.Close())

Важно:
- Таймаут на shutdown (30 сек) — если клиент не отвечает, принудительно закрыть
- Порядок закрытия: сначала WebSocket, потом HTTP, потом БД
- При деплое (rolling update) — новый pod поднимается ДО того как старый начнёт shutdown

#### { Пример кода }

```go
func main() {
	hub := NewHub()
	go hub.Run()

	srv := &http.Server{
		Addr: ":8080",
	}

	// Запуск HTTP сервера
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Ждём сигнал завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down...")

	// 1. Закрываем все WebSocket
	hub.Shutdown() // Close-фрейм каждому

	// 2. Shutdown HTTP с таймаутом
	ctx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second,
	)
	defer cancel()

	srv.Shutdown(ctx)

	// 3. Закрываем БД
	db.Close()

	log.Println("Server stopped")
}
```

#### { Типичная ошибка }

Забыть про WebSocket-соединения и сделать только srv.Shutdown(). HTTP-сервер закроется, но горутины ReadPump/WritePump продолжат работать с мёртвыми соединениями. Утечка горутин.

#### { Интервьюер хочет услышать }

Понимание порядка закрытия (WebSocket -> HTTP -> DB). Знание про Close-фрейм в протоколе WebSocket. И упоминание таймаута — нельзя ждать вечно.

### 5. «Как масштабировать WebSocket горизонтально?»

#### { Короткий ответ }

Один сервер — один Hub. При нескольких серверах пользователи на разных серверах не видят друг друга. Решение — Redis Pub/Sub (или NATS) как шина между серверами. Сообщение публикуется в Redis, все серверы получают его и доставляют своим клиентам.

#### { Развёрнутый ответ }

Проблема: пользователь A подключён к серверу 1, пользователь B — к серверу 2. A отправляет сообщение B. Hub на сервере 1 не знает про клиента на сервере 2.

Решение — внешняя шина сообщений:

1. Redis Pub/Sub — самый популярный вариант
   - Каждый сервер подписывается на канал (например, "chat:messages")
   - При отправке сообщения: сохраняем в БД, публикуем в Redis
   - Все серверы получают сообщение, каждый проверяет свой Hub
   - Плюсы: простой, быстрый, знакомый всем
   - Минусы: at-most-once delivery (если сервер упал — потеряет)

2. NATS / NATS JetStream — более надёжный вариант
   - JetStream даёт at-least-once и persistence
   - Лучше для критичных сообщений

3. Kafka — для highload (миллионы сообщений)
   - Партиционирование по chat_id
   - Replay — можно перечитать историю
   - Overkill для большинства задач

Дополнительно:
- Sticky sessions (балансировщик направляет клиента на тот же сервер)
Используйте код с осторожностью.
• Consistent hashing по user_id для предсказуемого распределения
• Service discovery для автоматического масштабирования

{ Пример кода }

go
// Публикация — при отправке сообщения
func (h *Hub) BroadcastMessage(msg *Message) {
	// 1. Сохраняем в БД
	h.messageService.Save(msg)

	// 2. Отправляем локальным клиентам
	h.deliverLocal(msg)

	// 3. Публикуем в Redis для других серверов
	data, _ := json.Marshal(msg)
	h.redis.Publish(ctx, "chat:messages", data)
}

// Подписка — при старте сервера
func (h *Hub) SubscribeRedis() {
	sub := h.redis.Subscribe(ctx, "chat:messages")

	for msg := range sub.Channel() {
		var m Message
		json.Unmarshal([]byte(msg.Payload), &m)

		// Доставляем только своим клиентам
		h.deliverLocal(&m)
	}
}
Используйте код с осторожностью.

{ Типичная ошибка }

Предлагать "общую базу данных" как решение синхронизации — это вернёт нас к polling. Или предлагать broadcast по HTTP между серверами — не масштабируется (N серверов = N^2 соединений).

{ Интервьюер хочет услышать }

Знание конкретного инструмента (Redis Pub/Sub) с объяснением trade-offs. Упоминание альтернатив (NATS, Kafka) и когда какой выбрать. И понимание что БД — source of truth, Redis — только доставка.
