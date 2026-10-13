package conditionalsloop

import (
	"fmt"
	"time"
)

func ConditionalsLoop() {

	fmt.Println("1.2 Условия (if/else, switch) и циклы (for)")

	if num := 17; num%2 == 0 {
		fmt.Println(num, " is even")
	} else {
		fmt.Println(num, "is odd")
	}
	// fmt.Println(num, "u here?") undefined: num. суть области видимости (scope) инициализатора

	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	// Роль while
	countdown := 5
	for countdown != 0 { // > 0
		fmt.Println(countdown)
		countdown--
	}
	fmt.Println("Let's go!")

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
}
