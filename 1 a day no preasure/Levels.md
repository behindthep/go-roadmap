🏗️ Три слоя бизнес-логики

❓ Что такое три слоя бизнес-логики?

Архитектурный подход, разделяющий приложение на три основных слоя для создания чистой, масштабируемой и поддерживаемой архитектуры.

💡 Характеристики слоёв:
- Слой представления (Presentation Layer):
  - Отвечает за взаимодействие с пользователем
  - Содержит контроллеры, обработчики HTTP-запросов
  - Минимальная бизнес-логика
  - Валидация входящих данных
  - Преобразование DTO

- Слой бизнес-логики (Business Logic Layer):
  - Реализация основных бизнес-процессов
  - Сервисы и UseCase
  - Сложная логика преобразования данных
  - Транзакционное управление
  - Бизнес-правила

- Слой данных (Data Access Layer):
  - Работа с базами данных
  - Репозитории
  - ORM-операции
  - Абстракция от конкретной базы данных

// Слой представления (Controller)
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    var dto UserDTO
    // Валидация входных данных
    if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    userService := NewUserService()
    user, err := userService.Create(dto)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user)
}

// Слой бизнес-логики (Service)
func (s *UserService) Create(dto UserDTO) (*User, error) {
    // Бизнес-правила
    if err := s.validator.Validate(dto); err != nil {
        return nil, err
    }

    // Сложная логика преобразования
    user := s.mapToUser(dto)

    // Транзакционность
    return s.repo.WithTransaction(func(tx *Transaction) (*User, error) {
        return tx.Save(user)
    })
}

// Слой данных (Repository)
func (r *UserRepository) Save(user *User) (*User, error) {
    return r.db.Create(user)
}


❗Почему это важно?
- Разделение ответственности: Каждый слой имеет чёткие границы
- Тестируемость: Легко изолировать и тестировать отдельные компоненты
- Масштабируемость: Возможность независимого развития слоёв
- Поддерживаемость: Clear separation of concerns

🎯 Три слоя — фундамент чистой архитектуры приложения
