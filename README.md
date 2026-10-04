# Golang Roadmap

```
go mod init github.com
```

В Go принято разделять код на приватный (internal) и точку входа (cmd).

```
mkdir -p cmd/app internal/config internal/handler internal/service internal/repository
```

• cmd/app/ — только main.go. Точка входа в приложение.
• internal/ — код внутри защищен компилятором Go. Его невозможно случайно импортировать в чужие внешние проекты.
	• config/ — инициализация конфигов и файла .env.
	• handler/ — слой доставки (API, HTTP-контроллеры, обработчики запросов).
	• service/ — ядро проекта (бизнес-логика).
	• repository/ — слой работы с БД или внешними хранилищами.
