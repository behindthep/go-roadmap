package structmethodpointer

import "fmt"

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

func StructMethodPointer() {

	fmt.Println("Структуры (struct) и Методы-------------------------------------------------")
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
}
