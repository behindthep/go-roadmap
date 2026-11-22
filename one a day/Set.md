📦 Set в Golang: эффективная работа с уникальными элементами

❓ Что такое Set?

Set — структура данных для хранения уникальных элементов. В Go реализуется через map с пустой структурой в качестве значения для оптимизации памяти.

💡 Полная реализация:
type Set struct {
    items map[interface{}]struct{}
}

func NewSet() *Set {
    return &Set{
        items: make(map[interface{}]struct{}),
    }
}

// Основные операции
func (s *Set) Add(item interface{}) {
    s.items[item] = struct{}{}
}

func (s *Set) Remove(item interface{}) {
    delete(s.items, item)
}

func (s *Set) Contains(item interface{}) bool {
    _, exists := s.items[item]
    return exists
}

// Операции над множествами
func (s *Set) Union(other *Set) *Set {
    result := NewSet()
    for item := range s.items {
        result.Add(item)
    }
    for item := range other.items {
        result.Add(item)
    }
    return result
}

func (s *Set) Intersection(other *Set) *Set {
    result := NewSet()
    for item := range s.items {
        if other.Contains(item) {
            result.Add(item)
        }
    }
    return result
}

⚠️ Преимущества:

- O(1) для основных операций
- Эффективное использование памяти
- Удобные операции над множествами
- Отсутствие дубликатов

❗Ограничения:

- Неупорядоченная структура
- Ключи должны быть сравнимыми
- Нет встроенной реализации
- Потребность в type assertions при использовании interface{}
