package main

import (
	"fmt"
	"time"
)

func main() {
	age := 3
	name := "Jonhy"
	isStudent := false
	fmt.Printf("Hi, my name is %s, i'm %d years old. Am I student? %t\n", name, age, isStudent)

	const DaysInWeek = 7
	// DaysInWeek = 8
	// 1. cannot assign to DaysInWeek (нельзя присвоить значение для DaysInWeek). Они создаются один раз во время компиляции, и их значение «вжигается» в бинарный файл. Изменить его во время работы программы физически невозможно.
	// 2. (neither addressable nor a map index expression) (значение не является ни адресуемым, ни выражением индекса карты). Компилятор Go пытается подсказать, почему это невозможно с точки зрения архитектуры памяти.
	// • Переменные имеют свой адрес в оперативной памяти (RAM), куда можно пойти и записать новые байты.
	// • Константы в Go не имеют адреса в памяти (к ним нельзя применить оператор взятия адреса &). Они существуют только как абстрактные значения на этапе сборки кода. Раз адреса в памяти нет — процессору буквально некуда записывать цифру 8.

	// Строгая статическая типизация: тип переменной известен при компиляции и не может меняться на лету.
	a := 10
	b := 5.5
	result := float64(a) + b // Приведение типов явное, автоматического приведения типов в Go нет.
	fmt.Printf("Sum of a (%d) and b (%0.1f) = (%0.1f)\n", a, b, result)

	// Целые числа: int (зависит от платформы, 32 или 64 бита), int8, int16, int32, int64. Беззнаковые аналоги: uint, uint8...
	// Числа с плавающей точкой: float32, float64.
	var (
		defInt  int
		defF64  float64
		defStr  string
		defBool bool
	)
	fmt.Println("(Default) Zero values:\nint =", defInt, "; float64 =", defF64, "; string =", defStr, "; bool =", defBool)

	fmt.Println()
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println()

	if num := 17; num%2 == 0 {
		fmt.Println(num, " is even")
	} else {
		fmt.Println(num, "is odd")
	}
	// fmt.Println(num, "u here?") undefined: num. суть области видимости (scope) инициализатора
	fmt.Println()

	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println()

	// Роль while
	countdown := 5
	for countdown != 0 { // > 0
		fmt.Println(countdown)
		countdown--
	}
	fmt.Println("Let's go!")
	fmt.Println()

	// switch заменяет длинные цепочки else if
	dayNumber := time.Now().Weekday()
	switch dayNumber {
	case time.Monday:
		fmt.Println("Monday")
	case time.Tuesday:
		fmt.Println("Tuesday")
	case time.Wednesday:
		fmt.Println("Wednesday")
	case time.Thursday:
		fmt.Println("Thursday")
	case time.Friday:
		fmt.Println("Friday")
	case time.Saturday, time.Sunday:
		fmt.Println("Weekends")
	default:
		fmt.Println("There is no such day of the week.")
	}
	fmt.Println()

	fizzBuzz := func(n int) { // анонимная функция
		for i := 1; i <= n; i++ {
			if i%3 == 0 && i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else if i%3 == 0 {
				fmt.Println("Fizz")
			} else if i%5 == 0 {
				fmt.Println("Buzz")
			} else {
				fmt.Println(i)
			}
		}
	}
	fizzBuzz(20)

	for {
		fmt.Println("Бесконечный цикл")
		break // return вышел бы из всей функции main
	}

	fmt.Println()
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println()

	Calculate := func(a, b float64, action string) (float64, error) { // запрещено создавать именованные вложенные функции
		// operators := []string{"+", "-", "*", "/"}
		// if !slices.Contains(operators, action) {
		// 	return 0, fmt.Errorf("передана неизвестная строка операции")
		// }
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
				return 0, fmt.Errorf("Zero division") // fmt.Errorf("текст") создаёт ошибку
			}
			result = a / b
		default:
			return 0, fmt.Errorf("Given undefined operation")
		}

		return result, nil // nil отсутствие ошибки
	}
	fmt.Println(Calculate(3, 4, "+"))
	fmt.Println(Calculate(3, 4, "u"))
	fmt.Println(Calculate(3, 4, "-"))
	fmt.Println(Calculate(3, 4, "*"))
	fmt.Println(Calculate(3, 4, "/"))
	fmt.Println(Calculate(3, 0, "/"))

	// • Задача 2. Сокращение аргументов и именованный возврат. Напиши функцию RectangleStats(width, height float64) (area, perimeter float64). сокращенную запись параметров одинакового типа. Используй именованные возвращаемые значения внутри тела функции и сделай пустой return. Вызови её в main и распечатай оба результата.

	// • Задача 3. Функция внутри функции. Создай в main переменную-функцию isEven, которая принимает int и возвращает bool (true, если число четное). Используй её внутри цикла от 1 до 10, чтобы выводить в консоль только четные числа.

	// • Задача 4. Множественный возврат в действии. Напиши функцию MinMax(nums ...int) (int, int). ...int (variadic параметры) передавать в функцию любое количество чисел через запятую. Внутри функции они работают как срез, по которому можно пройтись обычным циклом. Функция найти и вернуть минимальное и максимальное число из переданных. Вызови её как MinMax(10, -5, 30, 0, 7).

	fmt.Println()
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println()

}
