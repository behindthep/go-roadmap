📦 Pipeline в Go: разбиение задачи на этапы с использованием горутин

❓ Что такое Pipeline?

Pipeline — это паттерн, который разбивает задачу на несколько этапов обработки данных. Каждый этап выполняется в отдельной горутине, что позволяет организовать плавный поток данных.

💡 Полная реализация:
func Stage(input <-chan int, process func(int) int) <-chan int {
    output := make(chan int)
    go func() {
        for value := range input {
            output <- process(value)
        }
        close(output)
    }()
    return output
}

// Пример использования
func main() {
    input := make(chan int)
    go func() {
        for i := 0; i < 5; i++ {
            input <- i
        }
        close(input)
    }()

    stage1 := Stage(input, func(x int) int { return x * 2 })
    stage2 := Stage(stage1, func(x int) int { return x + 1 })

    for result := range stage2 {
        fmt.Println(result)
    }
}

⚠️ Преимущества:

- Разделение задачи на этапы.
- Параллельная обработка данных.
- Гибкость в добавлении новых этапов.

❗ Ограничения:

- Требует аккуратного управления каналами.
- Может быть сложным для отладки.
