package week2_

import "fmt"

// 2.1
type User struct {
	Name, Email string
	Age         int
	Address     // Встроенное поле (только тип структуры без имени поля)
}

func makeUser() {
	antony := User{Name: "Antony", Age: 24, Email: "antony@antoty.antony", Address: Address{"Havana", "First"}}
	fmt.Printf("%+v\n", antony) // {Name:Antony Email:antony@antoty.antony Age:24}
	fmt.Println(antony.IsAdult())

	fmt.Printf("Name: %s, City: %s, Street: %s.", antony.Name, antony.City, antony.Street)
}

// 2.2 приёмник
func (user User) IsAdult() bool {
	if user.Age >= 18 {
		return true
	}
	return false
}

// 2.3 Площадь прямоугольника. Структура Rectangle с Area(). Проверь результат через тест или печать.
type Rectangle struct {
	Width, Height float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

func makeRectangle() {
	basicRectangle := Rectangle{4.5, 3}
	fmt.Printf("%+v\n", basicRectangle)
	fmt.Println(basicRectangle.Area())
}

// 2.4 Композиция
type Address struct {
	City, Street string
}
