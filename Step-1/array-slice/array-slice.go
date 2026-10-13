package arrayslice

import (
	"fmt"
	"slices"
)

func ArraySlice() {

	fmt.Println("1.4.1: Массивы и Срезы (Slices) — Устройство, Капасити и Специфика append-------------------------------------------------")
	/*
		Массивы в Go используются редко. У них фиксированная длина, является частью их типа. [3]int и [5]int — два разных типа, их нельзя сравнивать или присваивать друг другу. Массивы передаются в функции по значению (копируются целиком), что неэффективно для больших объемов.

		Срез — динамическая «обертка» над массивом. typed only by the elements they contain (not the number of elements). Не хранит данные. Внутри структуры среза (Header) три элемента:
		1. Указатель (Pointer) на первый элемент базового массива.
		2. Длина (Length / len)
		3. Емкость (Capacity / cap) — сколько элементов вместит базовый массив, с точки старта среза.
	*/

	var mySlice []int // len=0, cap=0, указатель=nil.
	fmt.Println("nil-slice:", mySlice, "len =", len(mySlice), "cap =", cap(mySlice))
	for i := 1; i <= 10; i++ {
		mySlice = append(mySlice, i) // Аллокатор (подсистема управления оперативной памятью, отвечает за выделение и освобождение памяти для переменных и объектов во время работы программы). Go проверяет: хватает ли емкости (cap)?
		// Если хватает (len < cap), Go пишет элемент в базовый массив и увеличивает len на 1.
		// Если НЕ хватает (len == cap), Go совершает «тяжелую» операцию: выделяет в памяти новый базовый массив (x2 cap от прежнего), копирует данные, добавляет новый элемент, меняет указатель среза на новый массив и пересчитывает len и cap. Старый срез об этом не знает.
		fmt.Println(mySlice, "len =", len(mySlice), "cap =", cap(mySlice))
	}

	arr := [5]int{10, 20, 4: 50} // If you specify the index with :, the elements in between will be zeroed.
	s1 := arr[1:4]               // [:4] [2:]
	s1[0] = 999                  // указывает на срез s1, а не на базовый массив
	fmt.Println(s1)              // [999 0 0]
	fmt.Println(arr)             // [10 999 0 0 50] срез s1 — это лишь окно в массив arr. Изменение s1[0] затерло arr[1]
	cpy := make([]int, len(s1))
	copy(cpy, s1)
	fmt.Println("cpy:", cpy)
	if slices.Equal(s1, cpy) { // same for maps.Equal
		fmt.Println("s1 == cpy")
	}

	// Срез передается в функцию по значению, но копируется только заголовок (указатель, len, cap). Указатель копируется - все еще смотрит на тот же базовый массив.
	ModifySlice := func(s []int) {
		s[0] = 777         // изменение элемента внутри функции, изменит и снаружи. базовая память общая
		s = append(s, 888) // Если внутри функции сделать append, который превысит текущий cap, срез внутри функции переключится на новый массив, а старый срез снаружи функции останется смотреть на старый массив и ничего не узнает о новых элементах.
		// внутри функции произошел "escape" (выделение нового массива из-за нехватки cap), поэтому снаружи 888 мы не увидели.
		fmt.Println("Slice inside func:", s)
	}
	mainSlice := make([]int, 2) // len=2 (By default an array is zero-valued, which for ints means 0s), cap=2
	fmt.Println(mainSlice, len(mainSlice), cap(mainSlice))
	// append, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value
	mainSlice = append(mainSlice, 1, 2)
	fmt.Println(mainSlice)
	ModifySlice(mainSlice)
	fmt.Println(mainSlice) // 777 применилось, а 888 пропало

	optimizedSlice := make([]int, 0, 1000000) // срез с предварительно выделенной памятью
	fmt.Println("optimizedSlice:", len(optimizedSlice), cap(optimizedSlice))
	var notOptimizedSlice []int
	fmt.Println("notOptimizedSlice:", len(notOptimizedSlice), cap(notOptimizedSlice))
	for i := 1; i <= 1000000; i++ {
		optimizedSlice = append(optimizedSlice, i)
		notOptimizedSlice = append(notOptimizedSlice, i) // растёт скачками
		if i == 600 {
			// Оптимизация через make, чтобы компилятор Go не тратил ресурсы на постоянное выделение новой памяти и копирование элементов при append.
			fmt.Println("optimizedSlice:", len(optimizedSlice), cap(optimizedSlice))          // 600 1000000
			fmt.Println("notOptimizedSlice:", len(notOptimizedSlice), cap(notOptimizedSlice)) // 600 848 сколько лишних перевыделений памяти (848 -> 1000+ и т.д.) сделал бы Go для второго среза.
			break
		}
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // 2d:  [[0] [1 2] [2 3 4]]
}
