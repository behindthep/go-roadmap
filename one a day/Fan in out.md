📦 Fan-in и Fan-out в Go: распараллеливание задач и сбор результатов

❓ Что такое Fan-in и Fan-out?

Fan-out позволяет распараллелить выполнение задачи на несколько горутин, а Fan-in собирает результаты этих горутин в один поток данных.

💡 Полная реализация:
// Fan-out
func FanOut(input <-chan int, workers int) []<-chan int {
    outputs := make([]<-chan int, workers)
    for i := 0; i < workers; i++ {
        output := make(chan int)
        go func() {
            for value := range input {
                output <- value * 2
            }
            close(output)
        }()
        outputs[i] = output
    }
    return outputs
}

// Fan-in
func FanIn(inputs []<-chan int) <-chan int {
    output := make(chan int)
    var wg sync.WaitGroup
    for _, input := range inputs {
        wg.Add(1)
        go func(ch <-chan int) {
            for value := range ch {
                output <- value
            }
            wg.Done()
        }(input)
    }
    go func() {
        wg.Wait()
        close(output)
    }()
    return output
}

// Пример использования
func main() {
    input := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            input <- i
        }
        close(input)
    }()

    outputs := FanOut(input, 3)
    result := FanIn(outputs)

    for value := range result {
        fmt.Println(value)
    }
}

⚠️ Преимущества:

- Распараллеливание задач.
- Эффективное использование ресурсов.
- Удобство сбора результатов.

❗ Ограничения:

- Требует синхронизации (например, `sync.WaitGroup`).
- Может быть сложным для управления при большом количестве горутин.
