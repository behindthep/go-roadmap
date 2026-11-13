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
}

// 1.1
func miniCalculator() {
	var firstNum, secondNum float64

	fmt.Println("Write down two numbers:")
	// type not a pointer: int
	// В Go все аргументы передаются в функции по значению (копируются).
	// Если передать в fmt.Scan переменную, функция получит её локальную копию. запишет считанные данные в копию, но исходная переменная останется без изменений.
	// Передача указателя (&переменная) даёт функции доступ к ячейке памяти переменной.
	_, err := fmt.Scan(&firstNum, &secondNum)

	if err != nil {
		fmt.Println("Write only numbers:", err)
	} else {
		fmt.Println(firstNum + secondNum)
	}
}

// 1.2
func isEvenOrOdd() {
	var number int

	fmt.Println("Write down the number and find out it's even or odd:")
	_, err := fmt.Scan(&number)

	// if - ветвление логики программы
	if err != nil {
		fmt.Println("Write only number type:", err)
	} else {
		if isEven(number) {
			fmt.Println(number, " is Even.")
		} else {
			fmt.Println(number, " is Odd.")
		}
	}
}

func isEven(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}

// 1.3
// prints multiplication table from 1 to 10
func multiplicationTable() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			// formated print
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}
}
