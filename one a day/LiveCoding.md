### 01. Запросы по списку URL

Ozon и Data Driven Lab · давали в 27 собесах

```go
func main() {
	urls := []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://domain.tld",
		"https://ya.ru",
	}
	// для каждого адреса напечатай «адрес - ok» или «адрес - not ok»
}
```

Потом интервьюер раскручивает задачу. Запросы в горутинах, результаты в порядке списка, не больше k запросов одновременно, таймаут

Как ограничить число одновременных запросов и не оставить висеть горутины?
Разбор и решение с проверкой кода в уроке «Горутины, лимит, таймаут»

### 02. 3 или 6 секунд

Ozon, VK и М.Видео · давали в 24 собесах

```go
func main() {
	start := time.Now()
	_, _ = <-worker(), <-worker()
	fmt.Println(int(time.Since(start).Seconds()))
}

func worker() chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()
	return ch
}
```

Сколько секунд будет работать программа и как ускорить ее вдвое?
Разбор и решение с проверкой кода в уроке «Каналы»

### 03. Слить N каналов в один

Ozon, Ostrovok и Wildberries · давали в 24 собесах

```go
// merge сливает каналы в один и закрывает его,
// когда все входы прочитаны
func merge(cs ...<-chan int) <-chan int {
	// допиши
}

func main() {
	// gen пишет числа в канал и закрывает его
	a, b, c := gen(1, 2), gen(3), gen(4, 5)
	for v := range merge(a, b, c) {
		fmt.Println(v)
	}
}
```

Как закрыть выходной канал, когда все входы прочитаны?
Разбор и решение с проверкой кода в уроке «Балансировщик, кэш, merge»

### 04. x, y и z после append

Wildberries, Флант и Kaspersky · давали в 22 собесах

```go
x := []int{}
x = append(x, 0)
x = append(x, 1)
x = append(x, 2)
y := append(x, 3)
z := append(x, 4)
fmt.Println(y, z)
```

Что напечатает программа?
Разбор и решение с проверкой кода в уроке «Слайсы»

### 05. Горутины в цикле печатают i

Ozon, Магнит и FrogStudio · давали в 21 собесе

```go
for i := 0; i < 5; i++ {
	go func() {
		fmt.Println(i)
	}()
}
time.Sleep(time.Second)
```

Что напечатает программа и почему ответ зависит от версии Go?
Разбор и решение с проверкой кода в уроке «Гонки»

### 06. Балансировщик

Яндекс и ITeco · давали в 18 собесах

```go
type Backend interface {
	Invoke(ctx context.Context, req Request) (Response, error)
}

type Balancer struct {
	// допиши
}

func NewBalancer(addrs []string) *Balancer

// Invoke отдает запрос самому свободному экземпляру
func (b *Balancer) Invoke(ctx context.Context, req Request) (Response, error)
```

Экземпляр после нескольких ошибок подряд выключается на время

Где хранить число запросов в работе, чтобы балансировщик выдержал вызовы из многих горутин?
Разбор и решение с проверкой кода в уроке «Балансировщик, кэш, merge»

### 07. In-memory кэш

Ozon, ВсеИнструменты.ру и Астра · давали в 18 собесах

```go
type Cache interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

type cache struct {
	// допиши
}

func NewCache() Cache
```

Как сделать кэш безопасным для многих горутин и добавить TTL?
Разбор и решение с проверкой кода в уроке «Балансировщик, кэш, merge»

### 08. Покупки до бана

Ozon · давали в 17 собесах

```sql
user     (id, firstname, lastname, birth)
purchase (sku, price, user_id, date)
ban_list (user_id, date_from)
```

Выведи уникальные пары «пользователь и товар» по покупкам, сделанным до бана. Если пользователя не банили, берутся все его покупки

Где здесь ловушки с JOIN, дублями и однофамильцами?
Разбор и решение с проверкой кода в уроке «SQL»

### 09. addNum и nums[0:2]

Ozon · давали в 16 собесах

```go
func main() {
	nums := []int{1, 2, 3}
	addNum(nums[0:2])
	fmt.Println(nums) // ?
	addNums(nums[0:2])
	fmt.Println(nums) // ?
}

func addNum(nums []int) {
	nums = append(nums, 4)
}

func addNums(nums []int) {
	nums = append(nums, 5, 6)
}
```

Что напечатают оба Println?
Разбор и решение с проверкой кода в уроке «Слайсы»

### 10. uniqRandn

Ozon · давали в 16 собесах

```go
// uniqRandn возвращает n уникальных случайных чисел
func uniqRandn(n int) []int {
	// допиши
}

func main() {
	fmt.Println(uniqRandn(10))
}
```

Как вернуть n уникальных случайных чисел и какая у решения сложность?
Разбор и решение с проверкой кода в уроке «Алгоритмы»

### 11. defer x.S()

Ozon · давали в 14 собесах

```go
type X struct{ V int }

func (x X) S() { fmt.Println(x.V) }

func main() {
	x := X{123}
	defer x.S()
	x.V = 456
}
```

Что напечатает программа?
Разбор и решение с проверкой кода в уроке «defer и указатели»

### 12. Порядок обхода map

Ozon и Иннотех · давали в 13 собесах

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}

for k, v := range m {
	fmt.Println(k, v)
}
fmt.Println(m)
```

В каком порядке напечатаются строки цикла и последняя строка?
Разбор и решение с проверкой кода в уроке «Мапы и интерфейсы»

### 13. Аналитика тормозит заказы

Ozon · давали в 13 собесах

```go
func (s *OrderService) CreateOrder(ctx context.Context, o Order) error {
	if err := s.orders.Save(ctx, o); err != nil {
		return err
	}
	// аналитика тормозит, и вместе с ней тормозят заказы
	return s.analytics.TrackOrder(ctx, o)
}
```

Что поменяешь в схеме, чтобы заказы не зависели от аналитики и события не терялись?
Разбор и решение с проверкой кода в уроке «Задачи на сервис»

### 14. Неудовлетворенность покупателей

Авито · давали в 12 собесах

```go
// goods это цены товаров, buyerNeeds это пожелания покупателей.
// Каждый берет товар с ближайшей ценой, неудовлетворенность
// покупателя это разница цен. Верни сумму по всем покупателям
func totalDissatisfaction(goods []int, buyerNeeds []int) int
```

Какая сложность у перебора и как решить быстрее?
Разбор и решение с проверкой кода в уроке «Алгоритмы»

### 15. msgChan и stopChan

Kaspersky, Ximilabs и Солар · давали в 10 собесах

```go
msgChan := make(chan string, 1)
close(msgChan)
stopChan := make(chan struct{})
close(stopChan)

select {
case msgChan <- "msg":
	fmt.Println("msg sent")
case <-stopChan:
	fmt.Println("stop signal received")
}
```

Что произойдет при запуске?
Разбор и решение с проверкой кода в уроке «Каналы»

### ТОП-10 ЗАДАЧ НА ЛАЙВКОДИНГ

Ниже реальные задачи с собеседований в крупнейших компаниях. Для каждой задачи приведено условие и код, который дают кандидату. Попробуйте решить их самостоятельно за 15-20 минут.

#### { Задача 1 } Что выведет код со слайсами и append

Самая популярная задача на собеседованиях. Вам дают фрагмент кода и просят предсказать вывод, а затем объяснить почему. Ключевые моменты: слайс это заголовок (ptr, len, cap); append может как изменить существующий underlying array, так и создать новый; передача слайса в функцию копирует заголовок, но не данные.

```go
func appendToSlice(s []int) {
	s = append(s, 4)
}

func addToMap(m map[string]int) {
	m["b"] = 2
}

func main() {
	s := []int{1, 2, 3}
	appendToSlice(s)
	fmt.Println(s) // Что выведет?

	m := map[string]int{"a": 1}
	addToMap(m)
	fmt.Println(m) // Что выведет?
}
```

#### { Задача 2 } Параллельный Worker Pool с каналами

Дан входной канал данных и выходной канал результатов. Нужно реализовать параллельную обработку в N воркерах. Каждый воркер читает из in, обрабатывает данные и пишет в out. Закрыть out после завершения всех воркеров. Усложнение: добавить поддержку context для отмены.

```go
func processInParallel(in <-chan int, out chan<- int, workers int) {
	// Необходимо написать параллельную обработку
	// в кол-ве горутин равном workers.
	// Получаем данные из канала in,
	// над каждым элементом выполняем processData,
	// передаём результат в out.
	panic("not implemented")
}

func processData(data int) int {
	time.Sleep(100 * time.Millisecond)
	return data * 2
}
```

#### { Задача 3 } Merge N каналов в один (fan in)

Написать функцию merge, которая принимает произвольное количество входных каналов и объединяет их в один выходной. Выходной канал должен закрыться, когда все входные каналы закрыты. Не допускать утечек горутин.

```go
func merge(chans ...<-chan int) <-chan int {
	out := make(chan int)
	// Реализуйте объединение каналов.
	// Закройте out когда ВСЕ входные каналы
	// будут закрыты.
	return out
}
```

#### { Задача 4 } Code Review: кэширующий прокси (МойОфис)

В проекте используется key value БД (latency ~100мс). Нужно написать кэширующий прокси. Требования: потокобезопасность, пользователи кэша не должны знать про БД за ним, данные в памяти. Ниже заготовка. Найдите проблемы и улучшите код.

```go
type KVStore interface {
	Get(key string) (string, error) // часто
	Keys() ([]string, error)        // редко
}

type AddressStore struct {
	store KVStore
	cache map[string]string
	mu    sync.RWMutex
}

func (a *AddressStore) Get(key string) (string, error) {
	a.mu.RLock()
	ip, found := a.cache[key]
	a.mu.RUnlock()
	if found {
		return ip, nil
	}
	ip, err := a.store.Get(key)
	if err != nil {
		return "", err
	}
	a.mu.Lock()
	a.cache[key] = ip
	a.mu.Unlock()
	return ip, nil
}
```

#### { Задача 5 } Code Review: корзина с ошибками (tutu)

Дан код корзины интернет магазина. Нужно провести Code Review и найти как можно больше проблем. Затрагиваются темы: SOLID, конкурентность, безопасность, архитектура.

```go
var items []Product
var vatRate = 10.0
var removalChannel chan bool

func init() {
	items = make([]Product, 100) // 100 пустых!
	removalChannel = make(chan bool, 1)
}

func AddItemToShoppingCart(item Product) {
	items = append(items, item)
}

func RemoveItemFromShoppingCart(name string) {
	go func() {
		for i, item := range items {
			if item.NameOfProduct == name {
				items = append(items[:i], items[i+1:]...)
				removalChannel <- true
				return
			}
		}
	}()
	removalChannel <- false // deadlock!
}
```

#### { Задача 6 } Слайсы и shared underlying array (Ozon)

Вариация той же задачи, но с акцентом на общий underlying array. Когда cap слайса позволяет append не аллоцирует новый массив, а перезаписывает данные, которые «видны» через другой слайс.

```go
func main() {
	nums := []int{1, 2, 3}

	addNum(nums[0:2])
	fmt.Println(nums) // ?

	addNums(nums[0:2])
	fmt.Println(nums) // ?
}

func addNum(nums []int) {
	nums = append(nums, 4)
}

func addNums(nums []int) {
	nums = append(nums, 5, 6)
}
```

#### { Задача 7 } Горутины в цикле: замыкание на переменную

Классическая задача на понимание замыканий в Go. Что выведет этот код и почему? Как исправить? Примечание: в Go 1.22+ поведение изменилось переменная цикла теперь per iteration. Но на собесах по прежнему спрашивают.

```go
func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}
```

#### { Задача 8 } defer и значение аргумента (Ozon, WB)

Задача на понимание того, когда вычисляются аргументы defer. Два варианта: прямой вызов метода и вызов через замыкание. Что выведет каждый вариант?

```go
type X struct{ V int }

func (x X) S() { fmt.Println(x.V) }

// Вариант 1
func main() {
	x := X{123}
	defer x.S()
	x.V = 456
}

// Вариант 2
func main() {
	x := X{123}
	defer func() { x.S() }()
	x.V = 456
}
```

#### { Задача 9 } SQL: покупки до бана + агрегация

Даны три таблицы: users (id, name), purchases (id, user_id, product, amount, created_at) и ban_list (user_id, banned_at). 1) Вывести уникальные пары (пользователь, товар) для покупок, совершённых ДО бана. 2) Вывести пользователей с суммой покупок > 5000.

```sql
-- Запрос 1: покупки до бана
SELECT DISTINCT u.name, p.product
FROM users u
JOIN purchases p ON p.user_id = u.id
JOIN ban_list b ON b.user_id = u.id
WHERE p.created_at < b.banned_at;

-- Запрос 2: сумма покупок > 5000
SELECT u.name, SUM(p.amount) AS total
Используйте код с осторожностью.
FROM users u
JOIN purchases p ON p.user_id = u.id
GROUP BY u.name
HAVING SUM(p.amount) > 5000
ORDER BY total DESC;

#### { Задача 10 } In Memory Cache с TTL и потокобезопасностью

Реализовать потокобезопасный In Memory кэш с поддержкой TTL. Методы: Get(key) и Set(key, value, ttl). Элементы должны автоматически удаляться по истечении TTL. Усложнение: шардирование для уменьшения contention на мьютексе.

```go
type Cache struct {
	mu    sync.RWMutex
	items map[string]cacheItem
}

type cacheItem struct {
	value     interface{}
	expiresAt time.Time
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	// реализуйте
}

func (c *Cache) Get(key string) (interface{}, bool) {
	// реализуйте
}
