
FROM golang:1.21

# Установка рабочей директории внутри контейнера
WORKDIR /app

# Копирование файлов проекта внутрь контейнера
COPY . .

# Загрузка зависимостей (если у вас есть файл go.mod)
RUN go mod download

# запуск приложения
RUN go build -o main .

# Ваше приложение слушает порт 8080. Убедитесь, что этот порт доступен
EXPOSE 8000

# Команда для запуска приложения внутри контейнера
CMD ["./main"]



