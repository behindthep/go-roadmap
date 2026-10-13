package composition

import "fmt"

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

func Composition() {

	fmt.Println("Композиция (встраивание структур)-------------------------------------------------")
	/*
		В Go нет наследования. «Композиция лучше наследования». Не строим иерархии, а собираем сложные объекты из простых, LEGO.
	*/

	company := Company{"Jarvis", "Moscow", Address{"Saint Petersburg", "Green 2"}}
	fmt.Println(company.City)
	fmt.Println(company.Address.City)

	// Сквозной доступ к полям и методам (Promoted fields)

	myCompany, _ := NewCompany("Stark", "Boston", Address{"York", "Sleepy 11"})
	fmt.Println(myCompany.FullAddress()) // метод «поднялся» наверх
}
