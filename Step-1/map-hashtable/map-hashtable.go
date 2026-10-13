package maphashtable

import "fmt"

func MapHashtable() {

	fmt.Println("1.4.2: Хеш-таблицы (Map) — Устройство, инициализация и Runtime Panic-------------------------------------------------")
	/*
		map — встроенный тип ассоциативного массива (ключ-значение).
	*/

	// В Go разница между объявлением map и её инициализацией:
	var m map[string]int // nil-map. len = 0.
	// Разделение на var m ... и затем m = make(...) используют редко, когда объявление переменной и её заполнение разнесены по разным блокам кода (например, в условиях).
	fmt.Println(m["key"]) // Из nil-map можно читать ключи. Ключа нет - Go вернет дефолтное значение (zero value) для типа.
	// m["key"] = 42  panic: assignment to entry in nil map. в nil-map запрещено записывать данные. Runtime Panic (программа упадет).

	// Перед записью в карту, её инициализировать:
	// Если размер неизвестен и она пустая: m := make(map[string]int) или m := map[string]int{} — идентичны.
	// makeMap := make(map[string]int, 100) Если карта создается пустой, но известна её будущая емкость (туда запишут 100 элементов): make с указанием размера. Это выделит память под бакеты заранее и сбережет ресурсы рантайма на переаллокацию.
	// literalMap := map[string]int{"Alex": 5} Если карта заполняется сразу: используй короткую форму и литерал
	m = make(map[string]int)
	// Под капотом map в Go — заголовок (структура hmap из пакета runtime), который содержит метаданные, и набор бакетов (buckets / ведер).
	// Добавляешь элемент, Go берет хеш от ключа, определяет нужный бакет и складывает пару туда.
	m["key"] = 42
	fmt.Println(m["key"])

	words := []string{"go", "java", "go", "python", "go", "java"}
	wordMap := map[string]int{}
	for i := 0; i < len(words); i++ {
		wordMap[words[i]]++
	}
	fmt.Println(wordMap)

	// Если ключа в карте нет, Go возвращает zero value. В карте лежит реальный 0 или ключа просто не существует?
	thereIsStudent := func(name string) {
		grades := map[string]int{"Alex": 5, "Jane": 0}
		val, ok := grades[name] // Проверка существования ключа.
		// значение можно опустить blank identifier _, ok
		if ok {
			fmt.Println("У", name, "оценка", val)
		} else {
			fmt.Println("Студента", name, "нет в базе")
		}
	}
	thereIsStudent("Jane")
	thereIsStudent("Tom")

	DeleteKey := func(m map[string]int, key string) {
		delete(m, key)
	}
	mapInFunc := map[string]int{"Alex": 5, "Jane": 0, "Angel": 12}
	fmt.Println(mapInFunc)
	DeleteKey(mapInFunc, "Jane")
	// В отличие от срезов, map — под капотом является указателем (pointer) на hmap. Передаешь map в функцию, копия заголовка не создается. Функция получает тот же указатель. Изменения карты внутри функции видны на карте снаружи, без небоходимости возврата значения через аналоги append.
	fmt.Println(mapInFunc)

	clear(mapInFunc)
	fmt.Println("cleared map:", mapInFunc) // remove all pairs map[]
}
