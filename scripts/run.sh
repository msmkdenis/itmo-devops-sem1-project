#!/bin/bash

echo "Обновление зависимостей"
go mod tidy

echo "Запуск приложения"
go run main.go &

# Ожидание доступности сервера
echo "Ожидаем, пока сервер будет готов..."
for i in {1..10}; do
  if curl -s http://localhost:8080 &> /dev/null; then
    echo "Сервер успешно запущен и готов к работе."
    exit 0
  fi
  echo "Попытка $i: сервер пока не готов..."
  sleep 2
done

echo "Ошибка: сервер не запустился."
exit 1