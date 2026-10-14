package main

// 6. / packages /
import (
	"fmt"
	"go-roadmap/mathutils"
	"slices"
)

// 3. / func /
func Calculate(a, b float64, action string) (float64, error) {
	result := 0.0
	switch action {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Zero division")
		}
		result = a / b
	default:
		return 0, fmt.Errorf("Given undefined operation")
	}

	return result, nil // nil отсутствие ошибки
}

func RectangleStats(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // в больших функциях ухудшает читаемость
}

// ...int variadic параметры. передавать в функцию переменное числом аргументов. Внутри функции срез, по которому можно пройтись циклом. equivalent to []int. We can iterate over it with range, etc.
func MinMax(nums ...int) (int, int) {
	fmt.Println(nums, " ")
	min, max := nums[0], nums[0]
	var i, length int = 1, len(nums)

	for i < length {
		if nums[i] <= min {
			min = nums[i]
		}
		if nums[i] >= max {
			max = nums[i]
		}
		i++
	}

	return min, max
}

// 4. / array / slice / map(hash table) /
// Срез передается в функцию по значению, но копируется только заголовок (указатель, len, cap). Указатель копируется - все еще смотрит на тот же базовый массив.
func ModifySlice(s []int) {
	s[0] = 777         // изменение элемента внутри функции, изменит и снаружи. базовая память общая
	s = append(s, 888) // Если внутри функции сделать append, который превысит текущий cap, срез внутри функции переключится на новый массив, а старый срез снаружи функции останется смотреть на старый массив и ничего не узнает о новых элементах.
	// внутри функции произошел "escape" (выделение нового массива из-за нехватки cap), поэтому снаружи 888 мы не увидели.
	fmt.Println("Slice inside func:", s)
}

// Если ключа в карте нет, Go возвращает zero value. В карте лежит реальный 0 или ключа просто не существует?
func thereIsStudent(name string) {
	grades := map[string]int{"Alex": 5, "Jane": 0}
	val, ok := grades[name] // Проверка существования ключа.
	// значение можно опустить blank identifier _, ok
	if ok {
		fmt.Println("У", name, "оценка", val)
	} else {
		fmt.Println("Студента", name, "нет в базе")
	}
}

func DeleteKey(m map[string]int, key string) {
	delete(m, key)
}

// 5. / struct / method / pointer / composition /
func Double(x *int) {
	*x *= 2 // значение по адресу x умножать на 2. записывать результат обратно по этому же адресу
}

// Структура — пользовательский тип данных, объединяющий связанные поля
type Employee struct {
	// Инкапсуляция.
	Name     string  // имя с заглавной, публично доступно из других пакетов
	position string  // с маленькой — приватно (внутри пакета)
	Salary   float64 // Касается и имени структуры
}

func (e Employee) GetDetails() { //  тут эффективнее ресивер значение - просто читаем
	fmt.Println(e.Name, e.Salary, e.position)
}

// Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
func (e *Employee) UpdateSalary(bonus float64) { // * В типах данных: переменная хранит адрес, а не само значение (указатель). e *Employee — тип «указатель на Employee».
	e.Salary += bonus // * Перед переменной (Разыменование): «Сходи по этому адресу в память и возьми то, что там лежит». В Go рантайм делает это автоматически e.Salary += bonus, e — указатель. Полная запись (*e).Salary, но Go бережет пальцы.
}

// В Go роль конструкторов функции New[ИмяСтруктуры]. It’s idiomatic to encapsulate new struct creation in constructor functions
func NewEmployee(name, position string, salary float64) (*Employee, error) {
	if salary <= 0 {
		return nil, fmt.Errorf("Invalid salary format")
	}
	emp := Employee{name, position, salary}
	return &emp, nil // Go is a garbage collected language; you can safely return a pointer to a local variable - it will only be cleaned up by the garbage collector when there are no active references to it.
	// в Go безопасно возвращать адрес локальной переменной (return &emp, nil) — абсолютно верный. В языках вроде C++ это привело бы к катастрофе ( dangling pointer), но компилятор Go проводит escape-анализ и понимает, что переменную нужно перенести из стека в кучу (heap), чтобы GC не удалил её раньше времени.
}

type Address struct {
	City, Street string
}
type Company struct {
	CompanyName, City string // Внешнее поле «затеняет» Shadowing внутреннее.
	Address                  // Вместо наследования одна структура включать другую, без указания имени (Anonymous fields). Встраивание (Embedding). нет имени, только тип.
}

func NewCompany(name, city string, adr Address) (*Company, error) {
	company := Company{name, city, adr}

	return &company, nil
}

func (a Address) FullAddress() string {
	return "c. " + a.City + ", st. " + a.Street
}

//

func langbase() {
	// 1. / var / const / type /
	const DaysInWeek = 7 // Константы создаются один раз во время компиляции, их значение «вжигается» в бинарный файл. Изменить его во время работы нельзя. (neither addressable nor a map index expression) (значение не является ни адресуемым, ни выражением индекса карты). У констант нет адреса в RAM (нельзя взять адрес &), процессору некуда записывать значение. Абстрактные значения на этапе сборки кода.

	// Строгая статическая типизация: тип переменной известен при компиляции и не может меняться в рантайме.
	a := 10                  // Целые числа: int (зависит от платформы, 32 или 64 бита), int8, int16, int32, int64. Беззнаковые аналоги: uint, uint8...
	b := float64(a) + 5.5    // Числа с плавающей точкой: float32, float64.
	fmt.Printf("%0.1f\n", b) // Приведение типов явное, автоматического приведения нет.

	// 2. / loops /
	// 	for, количество итераций известно. Scope Переменная недоступна после завершения.
	// 	for condition (while), до выполнения условия, которое зависит от внешних факторов (чтение данных из сети или файла), количество итераций неизвестно. Переменная снаружи.

	// 3. / func /
	/*
		В Go функции «гражданами первого класса» (их можно передавать как аргументы, возвращать из других функций и присваивать переменным, как ты уже сделал с fizzBuzz).

		myFunc - приватна для текущего пакета. MyFunc - публична (экспортируема) для других пакетов.

		for _, num := range &nums {
		    total += num
		}
	*/
	numberSlice := []int{100, -15, 20, 0}
	fmt.Println(MinMax(numberSlice...)) // Распаковка

	// 4. / array / slice / map (hash table) /
	/*
		Массивы в Go используются редко. У них фиксированная длина, является частью их типа. [3]int и [5]int — два разных типа, их нельзя сравнивать или присваивать друг другу. Массивы передаются в функции по значению (копируются целиком), что неэффективно для больших объемов.

		Срез — динамическая «обертка» над массивом. typed only by the elements they contain (not the number of elements). Не хранит данные. Внутри структуры среза (Header) три элемента:
		1. Указатель (Pointer) на первый элемент базового массива.
		2. Длина (Length / len)
		3. Емкость (Capacity / cap) — сколько элементов вместит базовый массив, с точки старта среза.
	*/
	var mySlice []int // len=0, cap=0, указатель=nil.
	fmt.Println("nil-slice:", mySlice, "len =", len(mySlice), "cap =", cap(mySlice))
	for i := 1; i <= 10; i++ {
		mySlice = append(mySlice, i) // Аллокатор (подсистема управления оперативной памятью, отвечает за выделение и освобождение памяти для переменных и объектов во время работы программы). Go проверяет: хватает ли емкости (cap)?
		// Если хватает (len < cap), Go пишет элемент в базовый массив и увеличивает len на 1.
		// Если НЕ хватает (len == cap), Go совершает «тяжелую» операцию: выделяет в памяти новый базовый массив (x2 cap от прежнего), копирует данные, добавляет новый элемент, меняет указатель среза на новый массив и пересчитывает len и cap. Старый срез об этом не знает.
		fmt.Println(mySlice, "len =", len(mySlice), "cap =", cap(mySlice))
	}

	arr := [5]int{10, 20, 4: 50} // If you specify the index with :, the elements in between will be zeroed.
	s1 := arr[1:4]               // [:4] [2:]
	s1[0] = 999                  // указывает на срез s1, а не на базовый массив
	fmt.Println(s1)              // [999 0 0]
	fmt.Println(arr)             // [10 999 0 0 50] срез s1 — это лишь окно в массив arr. Изменение s1[0] затерло arr[1]
	cpy := make([]int, len(s1))
	copy(cpy, s1)
	fmt.Println("cpy:", cpy)
	if slices.Equal(s1, cpy) { // same for maps.Equal
		fmt.Println("s1 == cpy")
	}

	mainSlice := make([]int, 2) // len=2 (By default an array is zero-valued, which for ints means 0s), cap=2
	fmt.Println(mainSlice, len(mainSlice), cap(mainSlice))
	// append, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value
	mainSlice = append(mainSlice, 1, 2)
	fmt.Println(mainSlice)
	ModifySlice(mainSlice)
	fmt.Println(mainSlice) // 777 применилось, а 888 пропало

	optimizedSlice := make([]int, 0, 1000000) // срез с предварительно выделенной памятью
	fmt.Println("optimizedSlice:", len(optimizedSlice), cap(optimizedSlice))
	var notOptimizedSlice []int
	fmt.Println("notOptimizedSlice:", len(notOptimizedSlice), cap(notOptimizedSlice))
	for i := 1; i <= 1000000; i++ {
		optimizedSlice = append(optimizedSlice, i)
		notOptimizedSlice = append(notOptimizedSlice, i) // растёт скачками
		if i == 600 {
			// Оптимизация через make, чтобы компилятор Go не тратил ресурсы на постоянное выделение новой памяти и копирование элементов при append.
			fmt.Println("optimizedSlice:", len(optimizedSlice), cap(optimizedSlice))          // 600 1000000
			fmt.Println("notOptimizedSlice:", len(notOptimizedSlice), cap(notOptimizedSlice)) // 600 848 сколько лишних перевыделений памяти (848 -> 1000+ и т.д.) сделал бы Go для второго среза.
			break
		}
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // 2d:  [[0] [1 2] [2 3 4]]

	/*
		map — встроенный тип ассоциативного массива (ключ-значение).
	*/

	// В Go разница между объявлением map и её инициализацией:
	var m map[string]int // nil-map. len = 0.
	// Разделение на var m ... и затем m = make(...) используют редко, когда объявление переменной и её заполнение разнесены по разным блокам кода (например, в условиях).
	fmt.Println(m["key"]) // Из nil-map можно читать ключи. Ключа нет - Go вернет дефолтное значение (zero value) для типа.
	// m["key"] = 42  panic: assignment to entry in nil map. в nil-map запрещено записывать данные. Runtime Panic (программа упадет).

	// Перед записью в карту, её инициализировать:
	// Если размер неизвестен и она пустая: m := make(map[string]int) или m := map[string]int{} — идентичны.
	// makeMap := make(map[string]int, 100) Если карта создается пустой, но известна её будущая емкость (туда запишут 100 элементов): make с указанием размера. Это выделит память под бакеты заранее и сбережет ресурсы рантайма на переаллокацию.
	// literalMap := map[string]int{"Alex": 5} Если карта заполняется сразу: используй короткую форму и литерал
	m = make(map[string]int)
	// Под капотом map в Go — заголовок (структура hmap из пакета runtime), который содержит метаданные, и набор бакетов (buckets / ведер).
	// Добавляешь элемент, Go берет хеш от ключа, определяет нужный бакет и складывает пару туда.
	m["key"] = 42
	fmt.Println(m["key"])

	words := []string{"go", "java", "go", "python", "go", "java"}
	wordMap := map[string]int{}
	for i := 0; i < len(words); i++ {
		wordMap[words[i]]++
	}
	fmt.Println(wordMap)

	mapInFunc := map[string]int{"Alex": 5, "Jane": 0, "Angel": 12}
	fmt.Println(mapInFunc)
	DeleteKey(mapInFunc, "Jane")
	// В отличие от срезов, map — под капотом является указателем (pointer) на hmap. Передаешь map в функцию, копия заголовка не создается. Функция получает тот же указатель. Изменения карты внутри функции видны на карте снаружи, без небоходимости возврата значения через аналоги append.
	fmt.Println(mapInFunc)

	clear(mapInFunc)
	fmt.Println("cleared map:", mapInFunc) // remove all pairs map[]

	// 5. / struct / method / pointer / composition /

	/*
		Компилятор работает на этапе сборки.
		Runtime — среда выполнения Go, которая управляет памятью во время работы программы.


		В Go нет классов и объектно-ориентированного наследования. Cтруктуры (для хранения данных) и методы (для реализации поведения) заменой.

		Метод (Связка функций со структурами) — функция, у которой есть получатель (receiver) в скобках перед именем функции.
		Привязываются к типам через ресивер. Получатели двух типов:
		Получатель-значение (Value receiver):
		func (u User) Greet() { // При вызове метода рантайтм делает копию структуры User (безопасно для чтения). Изменения полей внутри метода НЕ изменят структуру снаружи.
			fmt.Println("Привет, я", u.Name)
		}

		Получатель-указатель (Pointer receiver):
		func (u *User) CelebrateBirthday() { // В метод передается адрес структуры. Изменения полей меняют оригинальный объект. Использовать указатель: если метод должен изменить состояние структуры / структура тяжелая (много полей) — чтобы рантайм не тратил ресурсы на копирование данных в памяти.
			u.Age++
		}

		Переменная — квартира, в которой живут данные (42). У квартиры два параметра:
		- Что внутри: значение 42.
		- Адрес: улица и номер дома, где квартира находится в оперативной памяти (0xc0000140e8).

		1. Оператор & — Узнаем адрес (Взятие адреса)
		& — риелтор. не показывает, что внутри квартиры, просто выдает её адрес.
		// &age — «Дай мне адрес переменной».
		age := 25
		fmt.Println(&age) // 0xc0000140e8 (адрес памяти)

		Тип переменной, которая хранит адрес, а не число (значение), записывается со звездочкой: *int («указатель на int»).

		2. Оператор * — Сходим по адресу (Разыменование)
		Если на руках есть адрес (*int), узнать или изменить то, что лежит внутри «квартиры», ставишь звездочку перед переменной-указателем.
		ptr := &age // ptr хранит адрес памяти. Тип ptr — *int
		fmt.Println(*ptr) // 25 рантайм сходил по адресу
		*ptr = 30         // зашли в квартиру и поменяли значение прямо там. «Сходи по адресу и возьми/измени значение внутри».
		fmt.Println(age)  // 30 (оригинальная переменная изменилась)

		3. Зачем это в функциях?
		По умолчанию аргументы в функции копируются. Передашь в функцию переменную и изменишь её, оригинал снаружи не изменится. Передашь адрес (&), функция сможет зайти в оригинальную «квартиру» и изменить данные там.

		Указатель смотрит «в никуда» - значение = nil. Зайти в nil-адрес к панике (nil pointer dereference).
	*/

	literalEmp := Employee{"Jonh", "Boss", 228.0}
	fmt.Println(literalEmp)
	var emptyEmp Employee // До инициализации рантайм заполняет поля их дефолтными значениями (zero value)
	fmt.Println(emptyEmp)

	literalEmp.GetDetails()

	literalEmp.UpdateSalary(12)
	fmt.Println(literalEmp.Salary)

	// & (Оператор взятия адреса): Берет существующую переменную и находит её «прописку» в оперативной памяти. &literalEmp: «Дай мне адрес в памяти, где лежит этот сотрудник».

	fmt.Println(NewEmployee("Anton", "Jawler", 555))
	fmt.Println(NewEmployee("Alex", "Idiot", -999))

	score := 100
	ptr := &score // адрес score
	*ptr = 200
	fmt.Println(score) // 200

	Double(&score)
	fmt.Println(score) // 400

	var strPtr *string            // указатель на строку через полную форму
	fmt.Println("strPtr", strPtr) // Нулевое значение указателя. nil дефолтный адрес (zero value) для указателей

	// *strPtr = "hello" // разыменование нулевого указателя
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x8 pc=0x4a5968] программа пытается обратиться к памяти недопустимым способом

	// Перед раскомментированием указателя в него нужно положить валидный адрес (new(string) или адрес сущ переменной).

	/*
		В Go нет наследования. «Композиция лучше наследования». Не строим иерархии, а собираем сложные объекты из простых, LEGO.
	*/

	company := Company{"Jarvis", "Moscow", Address{"Saint Petersburg", "Green 2"}}
	fmt.Println(company.City)
	fmt.Println(company.Address.City)

	// Сквозной доступ к полям и методам (Promoted fields)

	myCompany, _ := NewCompany("Stark", "Boston", Address{"York", "Sleepy 11"})
	fmt.Println(myCompany.FullAddress()) // метод «поднялся» наверх

	// 6. / packages /
	/*
		Пакеты помогают организовывать код в изолированные модули и управлять областью видимости.

		Компилятор ищет пакет main и функцию main() внутри: это точка входа, из этого кода собрать исполняемый бинарный файл. Код в других пакетах (package math) компилируется в библиотеки, которые нельзя запустить сами по себе.

		2. Область видимости в пакете
		В рамках пакета (даже если код разнесен по 10 файлам в одной папке) функции, структуры и переменные видны без импортов.

		• func Calculate() — экспортируемая. импортировать и вызвать в других пакетах  имя_пакета.Calculate().
		• func calculate() — неэкспортируемая. Доступна только внутри пакета.

		3. Система модулей (go mod)
		Управлять зависимостями и импортировать собственные локальные пакеты.

		Если импортировал пакет, но не используешь его, компилятор ошибку. Go за чистоту и не позволяет тащить мертвый код.


		module go-roadmap в go.mod?
		module go-roadmap — объявление имени модуля (корневого пути проекта).
		Компилятор Go использует это имя как базовый префикс для импорта внутренних пакетов проекта.
		пакет mathutils лежал в папке algebra, в main.go  импортировал его по полному пути относительно корня модуля:
		import "go-roadmap/algebra"

		go.mod превращает папку на диске в изолированный Go-модуль, позволяя рантайму управлять зависимостями и связывать файлы при сборке.
	*/
	fmt.Println(mathutils.Add(2, 3))
	// fmt.Println(mathutils.secretSubtract(4, 5)) // undefined: mathutils.secretSubtract  рантайм защищает приватные функции пакета от внешнего доступа
}

//

func main() {
	sayHello() // без импорта, если в том же пакете, даже если в разных папках

	langbase()
}
