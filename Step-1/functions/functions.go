package functions

import "fmt"

func Functions() {

	fmt.Println("1.3 Функции------------------------------------------")
	/*
		В Go функции «гражданами первого класса» (их можно передавать как аргументы, возвращать из других функций и присваивать переменным, как ты уже сделал с fizzBuzz).

		myFunc - приватна для текущего пакета. MyFunc - публична (экспортируема) для других пакетов.

		for _, num := range &nums {
		    total += num
		}
	*/

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

	RectangleStats := func(width, height float64) (area, perimeter float64) {
		area = width * height
		perimeter = 2 * (width + height)
		return // в больших функциях ухудшает читаемость
	}
	fmt.Println(RectangleStats(4, 5))

	vals := func() (int, int) {
		return 3, 7
	}
	_, c := vals()
	fmt.Println(c) // 7

	isEven := func(n int) bool {
		return n%2 == 0
		// if n%2 == 0 {
		// 	return true
		// }
		// return false
	}
	for i := 1; i <= 10; i++ {
		if isEven(i) {
			fmt.Println(i)
		}
	}

	// ...int variadic параметры. передавать в функцию переменное числом аргументов. Внутри функции срез, по которому можно пройтись циклом. equivalent to []int. We can iterate over it with range, etc.
	MinMax := func(nums ...int) (int, int) {
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
	fmt.Println(MinMax(10, -5, 30, 0, 7))
	numberSlice := []int{100, -15, 20, 0}
	fmt.Println(MinMax(numberSlice...)) // Распаковка
}
