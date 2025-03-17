# README

## 📌 Команды

### 🛠 Работа с Protobuf
Для генерации кода из `.proto` файлов нужно перейти в папку `src`:
```sh
protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out internal/api/docs proto/main.proto
```
> 💡 При ошибке добавьте путь к `protoc` в `PATH`:
```sh
export PATH=$PATH:$(go env GOPATH)/bin
```

### 📦 Миграции
Создание новой миграции:
```sh
goose -dir db/migrations create init sql
```

Откат последней миграции:
```sh
goose -dir db/migrations down
```

Применение всех миграций:
```sh
goose -dir db/migrations up
```

### 🧪 Генерация моков для тестов
```sh
mockery --name=ProgressModelInterface --dir=src/internal/models --output=src/internal/models/mocks --case=underscore
```

### ✅ Запуск тестов
```sh
go test ./...
```

### 🔍 Линтер
```sh
golangci-lint run
```
