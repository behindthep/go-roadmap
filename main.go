// с пакета main начинается выполнение программы. любая исполняемая программа должна быть в пакете main
package main

// подключение встроеенного пакета (модуля. набор функций) для форматированного вывода в консоль
import (
	"fmt"
	// "go-roadmap/week1basics"
)

// func main() - точка входа в программу. go run main.go, main() выполняется первой
func main() {
	var (
		innerText   = "LL"
		temperature float64
	)
	var i int
	var i2 = 1243
	i3 := -387

	var name, age = "Tom", 27

	fmt.Println("Hello, Juliet...", i, i2, i3, innerText, temperature, name, age)

	// week1
	// miniCalculator()
	// isEvenOrOdd()
	// multiplicationTable()

	// week2
}

// 2.1 Структура User с полями Name, Age, Email. В main создай пользователя и выведи его данные.
func () {
	// Как выглядит вывод если распечатать структуру целиком через:
	fmt.Printf("%+v, u")
}

// 2.2 Метда IsAdult() для User, возвращает true, если age >= 18.
// приёмник
func () {

}

// 2.3 Площадь прямоугольника. Структура Rectangle с Area(). Проверь результат через тест или печать.
func () {

}

// 2.4 Композиция. Создат Address и свтрой в User
// Выведи строку: Имя: ..., Город: ..., Улица: ...
func () {

}
