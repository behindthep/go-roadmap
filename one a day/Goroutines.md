📦 Обработка ошибок в горутинах: безопасное управление ошибками через каналы

❓ Как обрабатывать ошибки в горутинах?

Обработка ошибок в горутинах требует передачи ошибок через каналы и их корректной обработки в основном потоке.

💡 Полная реализация:
func Worker(input <-chan int, result chan<- int, errChan chan<- error) {
    for value := range input {
        if value == 0 {
            errChan <- fmt.Errorf("ошибка: значение равно 0")
            continue
        }
        result <- 10 / value
    }
}

// Пример использования
func main() {
    input := make(chan int)
    result := make(chan int)
    errChan := make(chan error)

    go Worker(input, result, errChan)

    go func() {
        for i := -2; i <= 2; i++ {
            input <- i
        }
        close(input)
    }()

    for {
        select {
        case res := <-result:
            fmt.Println("Результат:", res)
        case err := <-errChan:
            fmt.Println("Ошибка:", err)
        case <-time.After(1 * time.Second):
            fmt.Println("Завершение работы")
            return
        }
    }
}

⚠️ Преимущества:

- Безопасная обработка ошибок.
- Отделение логики обработки ошибок от основной логики.
- Удобство использования каналов для передачи ошибок.

❗ Ограничения:

- Требует аккуратного управления каналами.
- Может быть избыточным для простых задач.
