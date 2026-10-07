# Golang Roadmap

```
go mod init github.com
```

В Go принято разделять код на приватный (internal) и точку входа (cmd).

```
mkdir -p cmd/app internal/config internal/handler internal/service internal/repository
```

• cmd/app/ — только main.go. Точка входа в приложение.
• internal/ — код внутри защищен компилятором Go. Его невозможно случайно импортировать в чужие внешние проекты.
	• config/ — инициализация конфигов и файла .env.
	• handler/ — слой доставки (API, HTTP-контроллеры, обработчики запросов).
	• service/ — ядро проекта (бизнес-логика).
	• repository/ — слой работы с БД или внешними хранилищами.

---

## История Go
Google понадобился язык для решения проблемы: существовавшие на тот момент технологии перестали справляться с масштабами инфраструктуры компании.
Go создавался для огромных распределенных сетевых систем и высоконагруженных сервисов (Highload). 
Эффективно использовать мощность многоядерных процессоров, работая в облаках под миллионными запросами в секунду и оставаться простым в поддержке.

Бэкенд Google держался на: C++, Java и Python. Взять эффективность C++, скорость разработки Python и безопасность Java.
1. Медленная компиляция в C++
У Google были миллиарды строк кода на C++. Из-за сложной системы заголовочных файлов (#include) сборка огромных сервисов занимала часы.
В Go: Синтаксис и система импорта пакетов спроектированы, так что гигантские программы компилируются за секунды.
2. Сложная и тяжелая многопоточность
Процессоры перестали расти в тактовой частоте и стали многоядерными. Чтобы эффективно использовать железо, программы должны выполнять тысячи задач параллельно.
• В C++ параллельность требовала ювелирной работы с памятью, ошибки к падению серверов.
• В Java системные потоки (threads) были слишком «тяжелыми» — поток забирал 1 МБ памяти. Запустить одновременно 50 000 потоков на сервере было непозволительно.
• В Python существовал GIL (Global Interpreter Lock), не позволял выполнять Python-код параллельно на нескольких ядрах в рамках одного процесса.
В Go Goroutines — "легковесные потоки". Горутина занимает пару килобайт. Сервер на Go может запустить миллионы потоков одновременно.
3. Избыточная сложность языков (Проблема "Умников")
C++ и Java с годами обросли колоссальным количеством фич, абстракций, классов, шаблонов и метапрограммирования. Над проектом работают тысячи программистов, каждый пишет в своем стиле, код в кашу из-за избытка абстракций. Код нечитаемым, онбординг занимал месяцы.
• Из Go намеренно вырезали лишнее. Нет классического наследования классов, нет исключений (try-catch), нет неявных преобразований типов. Код однотипен и просто для чтения.
4. Проблемы управления памятью
• В C++ ручное управление памятью к утечкам и уязвимостям безопасности.
• В Java автоматический сборщик мусора (Garbage Collector) был тяжелым и периодически «замораживал» работу сервера на секунды, чтобы почистить память (эффект Stop-the-world), что недопустимо для поисковика или карт Google.
• В Go встроили эффективный сборщик мусора, работает параллельно и делает паузы в миллисекунды незаметными.

## Какую фундаментальную проблему индустрии решает Go?
Кризис сложности и масштабируемости программного обеспечения в эпоху многоядерных процессоров и облаков.
До появления Go индустрия застряла в дилемме:
• Писать код быстро , но программы получались медленными и неэффективными (Python, Ruby, PHP).
• Писать эффективный код, но разработка долгой, дорогой, а программы компилировались часами и содержали скрытые утечки памяти (C++).

Go дал индустрии компромисс: скорость выполнения и экономию ресурсов как в C++, но с простотой написания и чтения кода как в Python.

### Какие задачи бизнеса и инфраструктуры решает программист на Go?
Go — язык серверной логики и распределенных систем. На стыке потребностей бизнеса и стабильности серверов:
1. Задачи бизнеса (Деньги и продукт)
• Экономия на облачной инфраструктуре: Серверы на Go потребляют в разы меньше оперативной памяти и процессора, чем на Java или Python. Для бизнеса это уменьшение счетов за аренду серверов в Yandex Cloud или собственных дата-центров на 30–50%.
• Разработка API и бизнес-логики: Создание «мозгов» для мобильных приложений и сайтов. Код, который регистрирует пользователей, обрабатывает корзину покупок, рассчитывает скидки и интегрирует платежные системы (Мир, СБП).
• Стыковка систем (Интеграции): Бизнес редко состоит из одной программы. Разработчик пишет сервисы-шлюзы, связывают внутреннюю БД банка с внешними сервисами, государственными API или логистическими компаниями.
2. Задачи инфраструктуры (Надежность и Highload)
• Микросервисная архитектура: Перевод гигантских и неповоротливых старых программ (монолитов) на рельсы маленьких, независимых и быстрых сервисов, общаются по HTTP, gRPC или через очереди сообщений.
• Обработка сотен тысяч запросов (Highload): Настройка многопоточности, чтобы приложение не падало в моменты пиковых нагрузок.
• Оптимизация работы с БД: Написание быстрых SQL-запросов и построение кэширования (через Redis), чтобы пользователь получал ответ от сайта за миллисекунды, а БД не «задыхалась».

## Пет-проекты
Как production-код, не абстрактные учебные задачи, а инфраструктурные проблемы бизнеса: масштабируемость, отказоустойчивость, работу под высокой нагрузкой и микросервисное взаимодействие.
Система из нескольких сервисов, которая решает понятную прикладную бизнес-задачу.

1) Микросервисная система бронирования билетов (Электронная очередь / Кинотеатры / Стриминг)
Имитирует работу маркетплейсов в пиковые моменты нагрузок.
• Суть бизнес-логики: Пользователи могут просматривать доступные места, бронировать их на 10 минут и оплачивать. Место не должно быть куплено дважды (защита от Double Spending).
• Архитектурный стек под капотом:
	• Сервис 1 (Каталог): Только выдача расписания и мест. Кэшировать данные в Redis, чтобы не перегружать основную БД.
	• Сервис 2 (Бронирование): Принимает запросы на покупку. Чтобы БД не легла от одновременного наплыва тысяч людей, этот сервис не пишет данные сразу в БД, а кидает их в очередь сообщений.
	• Сервис 3 (Обработчик/Worker): Постепенно забирает сообщения из очереди и пачками (batch) обновляет статусы в PostgreSQL, используя строгие ACID-транзакции.

2) Распределенная система финтех-процессинга (Антифрод и Кошелек)
В финтехе важна точность данных, скорость работы и отслеживание подозрительных транзакций в реальном времени.
• Суть бизнес-логики: API принимает запросы на перевод денег между внутренними счетами пользователей. Перед транзакцией система мгновенно проверить, не является ли операция мошеннической.
• Архитектурный стек под капотом:
	• Взаимодействие сервисов по протоколу gRPC (Protobuf), а не по HTTP — так общаются микросервисы в проде для экономии трафика и скорости.
	• Реализация Graceful Shutdown в Go (правильное завершение работы): если вы экстренно обновляете или выключаете сервер, он должен не обрывать транзакции пользователей, а дождаться их завершения.
	• Хранение денег в БД с использованием SELECT FOR UPDATE (блокировки строк), чтобы два параллельных запроса не списали с баланса больше денег, чем там есть.

3) Инфраструктурный прокси-сервер / Ограничитель трафика (Rate Limiter)
Яндекс или Лаборатории Касперского обожают кандидатов, умеющих писать низкоуровневые инфраструктурные инструменты.
• Суть логики: умная «прослойка» (Middleware) перед основным сайтом, которая защищает его от перегрузки или DDOS-атак.
• Архитектурный стек под капотом:
	• Конкурентность на полную мощность: создание пула воркеров (Worker Pool), активное использование горутин, каналов и контекстов (context.Context) для управления таймаутами сетевых запросов.
	• Реализация алгоритма ограничения запросов (Token Bucket или Leaky Bucket) без использования сторонних библиотек. Написать его руками с помощью стандартного пакета sync (sync.Mutex или sync.Map), чтобы доказать глубокое знание многопоточности.

На собеседовании смотреть не на бизнес-идею, а на выполнение инженерных стандартов:
1. Логирование и Метрики (Observability): код не должен молча писать ошибки в консоль. Интегрируйте структурированное логирование (библиотеку uber-go/zap) и собирайте метрики для Prometheus.
2. Тестирование: Покройте ключевую бизнес-логику юнит-тестами (unit-tests) и интеграционными тестами с библиотекой testcontainers-go (для автоматического поднятия тестовой БД).
3. Упаковка в Docker: Напишите многоэтапный (multi-stage) Dockerfile для сборки минимального по размеру образа и docker-compose.yaml для локального запуска всей системы (сервис + БД + очереди) одной командой.
4. БД: Нет хаоса. Все изменения в структурах таблиц PostgreSQL управляться через инструмент миграций (golang-migrate)

## Ресурсы
- Go.dev/doc. Особое внимание статьям Concurrency is not parallelism (про горутины) и Share memory by communicating.
- blog.golang.org — статьи от команды разработки. Архивы про модули (Go Modules), обработку ошибок и тестирование. Учит мыслить в парадигме языка.
- coddy.tech/landing/ru/go.
  
## Заметки
- Блог/README - фиксация прогресса даёт +30% к видимости

## Чтение чужого кода
Самый быстрый способ вырасти.
 - github trending (go). раз в неделю. проекты со 100-500 звёздами.
 - grafana, prometheus, caddy
 - sourcegraph - как в реал проектах реализована конкретная задача (golang context timeout example in production)
 - документации стандартной библиотеки

## Собеседования / стажировки
- подаваться на 15-20 стажировок/вакансий - первые 2-3 отказа не считаются. trainee-программы.

---

Go — это высокоуровневый язык программирования, который создавался для быстрого написания микросервисов. Его придумали в Google в качестве замены С++ для тех проектов, где важна скорость разработки и компиляции кода. Go — довольно простой язык, в нём мало ключевых слов и неявных элементов. У него есть два козыря: эффективное использование вычислительных ресурсов при минимальных затратах и написание кода через примитивы многопоточности, встроенные в язык.
Язык не зря пользуется спросом среди маркетплейсов, доставки, рекламных сетей и других подобных проектов, ведь все эти области объединяет высокая нагрузка на сервисы.


Выделение памяти в Go управляется компилятором в зависимости от контекста. Go автоматически определяет, где требуется разместить переменную — на стеке или в куче, что избавляет от многих ошибок и уязвимостей.
Управление памятью осуществляется планировщиком, в результате чего нельзя выйти за пределы массива.
В Go отсутствует арифметика указателей. Указатель только может указывать на некоторый объект, нельзя создать указатель на произвольный объект памяти.
есть встроенный сборщик мусора, благодаря которому разработчики могут забыть про контроль и очистку памяти.
В библиотеке подсмотреть варианты обработки многопоточных сценариев.

какую бы функцию ни написал разработчик, её можно запустить в фоновом режиме, и она будет работать. В то же время планировщик Go сам распределит нагрузку по ядрам, чтобы каждое из них было эффективно нагружено. асинхронной


## Многопоточность
Что такое «поток» в контексте операционной системы? 
Поток выполнения (native/kernel thread) — часть процесса, в которой инструкции могут выполняться независимо и иметь доступ к общим ресурсам. За управление потоками отвечает планировщик ОС. 
Многопоточность — свойство железа и софта, при котором несколько потоков могут выполняться параллельно, не мешая друг другу. Если разработчики ПО сумели эффективно распараллелить отдельные части программы, можно рассчитывать на увеличение производительности, кратное количеству доступных ядер процессора.
Но для реализации работы нескольких потоков требуется как аппаратная поддержка, так и программная. Под аппаратной подразумевается наличие выделенных ядер в процессоре под каждый поток ОС (по два с hyper-threading), под программной — поддержка многопоточности ОС и конструкциями языка. необходимость в разработке многопоточного софта возникла только в нулевых, когда рынок начали захватывать многоядерные процессоры.
в наши дни производители процессоров стали наращивать количество вычислительных ядер вместо рекордных показателей тактовой частоты одного ядра. Эффективно использовать такие процессоры могут только многопоточные вычисления.

Один из первых нюансов, с которым сталкиваются разработчики при проектировании многопоточного ПО, — организация доступа к общим ресурсам, а конкретно — к памяти. Неверное разделение доступов между потоками может привести к порче данных. если два потока одновременно и параллельно пишут своё значение в одну переменную, какое из значений записано? необходимость мыслить несколькими потоками осложняет понимание процесса выполнения программы.
Средствами языка разработчикам хотелось избавиться от боли, возникнуть при переходе на многопоточное программирование.большинство существовавших на тот момент языков создавались без учёта многопоточности, её поддержку приделывали «сбоку».
В интерпретируемых яп, занимавших львиную долю рынка того времени, были доп сложности. организовать механизмы обращения к глобальному состоянию интерпретатора из нескольких потоков и обмен данными между этими потоками. необходим новый яп, который бы был изначально создан с расчётом на многопоточность и быстро и удобно писать приложения, использующие доступные ядра процессоров.

<img width="2800" height="1316" alt="image" src="https://github.com/user-attachments/assets/98d93a46-da0d-4880-884f-e2d76b3dbfaf" />

Для эффективного использования доступной вычислительной мощности и потоков ввода/вывода Go оперирует несколькими системными потоками, распределяя между ними ещё больше своих собственных легковесных потоков со стратегией m*n. на одном системном потоке могут исполняться несколько горутин. Если системный поток блокирован ожиданием ввода/вывода или перегружен, диспетчер Go может перенести горутину на свободный. Если захваченных системных потоков недостаточно, диспетчер Go может потребовать у системы новых.



---

## Синтаксис

```go
// A function can return any number of results.
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
var i, j int = 1, 2
func main() {
	var c, python, java = true, false, "no!"
	// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	k := 3
	fmt.Println(i, j, c, python, java)
}


// int, uint, and uintptr 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
int  int8  int16  int32(rune alias for int32. represents a Unicode code point)  int64
uint uint8(byte alias for uint8) uint16 uint32 uint64 uintptr
float32 float64
complex64 complex128

// When you need an integer value use int unless you have a specific reason to use a sized or unsigned integer type.

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
func main() {
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
}


// Variables declared without an explicit initial value are given their **zero value**:
0 for numeric typeS,
false for the boolean type, and
"" for strings.
func main() {
	var s string
	fmt.Printf("%q\n", s)
}

The expression T(v) converts the value v to the type T.
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
Or, put more simply:
i := 42
f := float64(i)
u := uint(f)
// Assignment between items of different type requires an explicit conversion.
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

var i int
j := i // int
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128

const Pi = 3.14 // Constants can be character, string, boolean, or numeric values.

// An int can store at maximum a 64-bit integer, and sometimes less.
const (
	Big = 1 << 100 // huge number by shifting a 1 bit left 100 places. binary number 1 followed by 100 zeroes
	Small = Big >> 99 // end up with 1<<1, or 2.
)
func needInt(x int) int { return x*10 + 1 }
func main() {
	fmt.Println(needInt(Big)) ./prog.go:22:22: cannot use Big (untyped int constant 1267650600228229401496703205376) as int value in argument to needInt (**overflows**)
}

func main() {
	sum := 1
	for i := 0; i < 10; i++ {
		sum += i
	}
	for ; sum < 1000; { // init and post statements are optional.
		sum += sum // 1024
	}
	for sum < 1000 { // C's while is spelled for in Go.
		sum += sum
	}
	for { // infinite loop
	}
}
	
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// Both calls to pow return their results before the call to fmt.Println in main begins.
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // short statement to execute before the condition
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
func main() {
	fmt.Println(
		pow(3, 2, 10), 
		pow(3, 3, 20), 
	)
}
// 27 >= 20 сначала вывод из функции
// 9 20 потом main


func Sqrt(x float64) float64 { // This general approach is called Newton's method
	// Try other initial guesses for z, like x, or x/2. How close are your function's results to the math.Sqrt in the standard library?
	z := 1.0 // To declare and initialize a floating point value, give it floating point syntax or use a conversion
	z := float64(1) // starting guess, no matter what the input
	// change the loop condition to stop once the value has stopped changing (or only changes by a very small amount). See if that's more or fewer than 10 iterations.
	for ;z <= 10; z++ { // Computers compute the square root of x using a loop
	// z² − x above is how far away z² is from where it needs to be (x), and the division by 2z is the derivative of z², to scale how much we adjust z by how quickly z² is changing
	z -= (z*z - x) / (2*z) // adjust z based on how close z² is to x, producing a better guess. Repeating this adjustment makes the guess better and better until we reach an answer that is as close to the actual square root as can be.
		fmt.Println(z)
	}
	return z;
}


// Go's switch cases need not be constants, and the values involved need not be integers.
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	switch i {
		case 0:
		case f():
	} // does not call f if i==0.)

	// Time in the Go playground always appears to start at 2009-11-10 23:00:00 UTC, a value whose significance is left as an exercise for the reader. ?
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// construct can be a clean way to write long if-then-else chains.
	t := time.Now()
	switch { // without a condition is the same as switch true.
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	defer fmt.Println("world") // defers the execution of a function until the surrounding function returns
	// deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	fmt.Println("hello")
}
// hello
// world

// Stacking defers. Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in **last-in-first-out order**. https://go.dev/blog/defer-panic-and-recover
func main() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
done 9 8 7


// A pointer holds the memory address of a value.
var p *int // type *T is a pointer to a T value. Its zero value is nil.
// Go has no pointer arithmetic.
func main() {
	i, j := 42, 2701
	p := &i         // point to i. & generates a pointer to its operand.
	// * denotes the pointer's underlying value.
	fmt.Println(*p) // 42. read i through the pointer p
	*p = 21         // set i through the pointer p
	// This is known as "dereferencing" or "indirecting".
	fmt.Println(i)  // 21

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // 73
}


type Vertex struct {
	X int
	Y int
}
// A struct literal denotes a newly allocated struct value by listing the values of its fields. The special prefix & returns a pointer to the struct value.

var (
	v1 = Vertex{1, 2}  // {1 2} type Vertex
	v2 = Vertex{X: 1}  // {1 0}
	v3 = Vertex{}	   // {0 0}
	p = &Vertex{1, 2} // &{1 2} type *Vertex
)
func main() {
	fmt.Println(Vertex{1, 2}) // {1 2}
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X) // 4

	p := &v // Struct fields can be accessed through a struct pointer.
	p.X = 1e9 // without the explicit dereference (*p).X
	fmt.Println(v) // {1000000000 2}
}


func main() {
	// An array's length is part of its type, so arrays cannot be resized.
	var a [2]string // type [n]T is an array of n values of type T
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) // Hello World
	fmt.Println(a) // [Hello World]

	primes := [6]int{2, 3, 5, 7, 11, 13}
	// slice is a dynamically-sized
	var s []int = primes[1:4] // [3 5 7]. type []T is a slice with elements of type T


	names := [4]string{
		"John", "Paul", "George", "Ringo",
	}
	// slice does not store any data, it just describes (view) a section of an underlying array
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // [John Paul] [Paul George]
	// slices are like references to arrays.
	b[0] = "XXX"
	fmt.Println(a, b) // [John XXX] [XXX George]
	fmt.Println(names) // [John XXX George Ringo]
}


func main() {
	// [3]bool{true, true, false} // an array literal
	// []bool{true, true, false} // slice literal is like an array literal without the length. creates the same array, then builds a slice that references it
	q := []int{2, 3, 5, 7, 11, 13}
	r := []bool{true, false, true, true, false, true}
	s := []struct {
		i int
		b bool
	}{
		{2, true}, {3, false}, {5, true}, {7, true}, {11, false}, {13, true},
	}
}


var a [10]int equivalent:
a[0:10] a[:10] a[0:] a[:]
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4] // [3 5 7]
	s = s[:2] // [3 5] уже на основе первого слайса а не базового массива
	s = s[1:] // [5]
}


// capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
func main() {
	s := []int{2, 3, 5, 7, 11, 13} // len=6 cap=6 [2 3 5 7 11 13]
	// Slice the slice to give it zero length.
	s = s[:0] // len=0 cap=6 []
	// Extend its length.
	s = s[:4] // len=4 cap=6 [2 3 5 7]
	// Drop its first two values.
	s = s[2:] // len=2 cap=4 [5 7]
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


The zero value of a slice is nil.
func main() {
	var s []int // The zero value of a slice is nil.
	// s := []int{} // not nil
	fmt.Println(s, len(s), cap(s)) // [] 0 0
	if s == nil {
		fmt.Println("nil!") // nil!
	}
}


func main() {
	// make this is how you create dynamically-sized arrays (Slices). make allocates a zeroed array and returns a slice that refers to that array:
	a := make([]int, 5)
	printSlice("a", a) // a len=5 cap=5 [0 0 0 0 0]
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice("b", b) // b len=0 cap=5 []
	// b = b[:cap(b)] // len(b)=5, cap(b)=5
	// b = b[1:]      // len(b)=4, cap(b)=4
	c := b[:2]
	printSlice("c", c) // c len=2 cap=5 [0 0]
	d := c[2:5]
	printSlice("d", d) // d len=3 cap=3 [0 0 0]
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}


func main() {
	board := [][]string{ // не просто слайс, а слайс внутри которого слайсы
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// board[1][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}


// If the backing array of s small to fit all given values a bigger array will be allocated. The returned slice will point to the newly allocated array. https://go.dev/blog/slices-intro
func main() {
	var s []int // len=0 cap=0 []
	// append works on nil slices.
	s = append(s, 0) // len=1 cap=1 [0]
	s = append(s, 1) // len=2 cap=2 [0 1]
	s = append(s, 2, 3, 4) // len=5 cap=6 [0 1 2 3 4]
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
func main() {
	for index, value := range pow { // range form of the for loop iterates over a slice or map 
		fmt.Printf("2**%d = %d\n", index, value)
		2**0 = 1
		2**1 = 2
		2**2 = 4
		2**3 = 8
		2**4 = 16
		2**5 = 32
		2**6 = 64
		2**7 = 128
	}
}


func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i битовый сдвиг
	}
	// for i := range pow // If you only want the index
	for _, value := range pow { // skip the index or value by assigning to _.
		fmt.Printf("%d\n", value)
		1
		2
		4
		8
		16
		32
		64
		128
		256
		512
	}
}


import "golang.org/x/tour/pic"
func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for i := 0; i < dy; i++ { // loop to allocate each []uint8 inside the [][]uint8
		s[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			s[i][j] = uint8((i+j)/2)
		}
	}	
	return s // slice of length dy, each element of which is a slice of dx 8-bit unsigned integers
}
func main() {
	pic.Show(Pic)
}


// zero value of a map is nil. A nil map has no keys, nor can keys be added.
type Vertex struct {
	Lat, Long float64
}
var m map[string]Vertex
func main() {
	m = make(map[string]Vertex) // make returns a map of the given type
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
}

// Map literals are like struct literals, but the keys are required.
var m = map[string]Vertex{
	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	// "Bell Labs": {40.68433, -74.39967},
	// "Google":    {37.42202, -122.08408},
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	}, // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
}


func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	m["Answer"] = 48
	delete(m, "Answer") // 0 - мэтчит тип
	v, ok := m["Answer"] 
	fmt.Println("The value:", v, "Present?", ok) // The value: 0 Present? false
}


strings.Fields helpful. https://pkg.go.dev/strings#Fields
import (
	"golang.org/x/tour/wc"
)
// Implement
func WordCount(s string) map[string]int {
	words := strings.Fields(s);
	leng := len(words);
	wcr := map[string]int
	for i: = 0; i <leng; i++ {
		v, ok := words[i]
		if ok {
			wcr[v]++
		}	
	}
	return map[string]int{"x": 1} // map of the counts of each “word” in the string s.
}
func main() {
	wc.Test(WordCount) // runs a test suite against the provided function and prints success or failure.
}


```
