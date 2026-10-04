# ==============================================================================
# Переменные (Настройки проекта)
# ==============================================================================
MODULE_NAME := $(shell head -n 1 go.mod | cut -d' ' -f2)
BINARY_NAME := my-go-app
MAIN_PATH   := ./main.go
DOCKER_TAG  := latest

# Цвета для вывода в консоль
IMAGE_NAME  := $(BINARY_NAME)
GREEN       := $(shell printf "\033[32m")
RESET       := $(shell printf "\033[0m")

# ==============================================================================
# Команды (Phony Targets)
# ==============================================================================
.PHONY: all help tidy fmt lint build run test test-cover clean docker-build

all: help

help:
	@echo "Доступные команды:"
	@echo "  make tidy         - Скачать и очистить зависимости Go"
	@echo "  make fmt          - Автоматическое форматирование кода (go fmt)"
	@echo "  make lint         - Проверка кода линтером (golangci-lint)"
	@echo "  make build        - Скомпилировать бинарный файл"
	@echo "  make run          - Собрать и запустить приложение локально"
	@echo "  make test         - Запустить модульные тесты"
	@echo "  make test-cover   - Запустить тесты и открыть отчет о покрытии"
	@echo "  make clean        - Удалить скомпилированные файлы"
	@echo "  make docker-build - Собрать локальный Docker-образ"

tidy:
	@echo "$(GREEN)⇒ Downloading and cleaning dependencies...$(RESET)"
	go mod tidy

fmt:
	@echo "$(GREEN)⇒ Formatting code...$(RESET)"
	go fmt ./...

lint:
	@echo "$(GREEN)⇒ Running linter...$(RESET)"
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "⚠️ golangci-lint не установлен. Пропустите или установите через: go install ://github.com"; \
	fi

build: tidy fmt
	@echo "$(GREEN)⇒ Building binary [$(BINARY_NAME)]...$(RESET)"
	CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/$(BINARY_NAME) $(MAIN_PATH)

# автоматически скачает зависимости, отформатирует код, скомпилирует его в папку ./bin/ и сразу запустит ваше приложение.
run: build
	@echo "$(GREEN)⇒ Running application...$(RESET)"
	./bin/$(BINARY_NAME)

# запустит все тесты с включенным детектором состояния гонки (-race).
test:
	@echo "$(GREEN)⇒ Running unit tests...$(RESET)"
	go test -v -race -cover ./...

# запустит тесты, посчитает процент покрытия кодом и автоматически откроет красивую интерактивную страницу в вашем браузере, где покажет, какие строки кода покрыты тестами, а какие нет.
test-cover:
	@echo "$(GREEN)⇒ Running tests with coverage report...$(RESET)"
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# очистит проект от временных файлов сборки и логов тестов.
clean:
	@echo "$(GREEN)⇒ Cleaning build artifacts...$(RESET)"
	rm -rf ./bin
	rm -f coverage.out

docker-build:
	@echo "$(GREEN)⇒ Building Docker image [$(IMAGE_NAME):$(DOCKER_TAG)]...$(RESET)"
	docker build -t $(IMAGE_NAME):$(DOCKER_TAG) .
