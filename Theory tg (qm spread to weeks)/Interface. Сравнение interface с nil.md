При сравнении интерфейса с nil возвращается true, только если и тип, и значение интерфейса = nil.

Интерфейс в Go — структура, состоит из двух компонентов:
1. Указатель на type: Содержит информацию о типе данных, который реализует интерфейс: *MyStruct.
2. Указатель на value: Содержит конкретное значение этого типа: указатель на структуру или nil.

```go
package main

import "fmt"

type MyInterface interface {
  Method()
}

type MyStruct struct{}

func (m *MyStruct) Method() {
  fmt.Println("Hi, Mark!")
}

func main() {
  var i MyInterface     // интерфейс не инициализирован, тип и значение = nil
  fmt.Println(i == nil) // true

  var s *MyStruct = nil // указатель на структуру = nil
  i = s                 // интерфейс содержит тип *MyStruct, значение nil
  fmt.Println(i == nil) // false
}
```

Сравниваете интерфейс с nil, проверяет оба компонента и, если хотя бы один из них != nil, интерфейс не считается nil.
