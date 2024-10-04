#!/bin/bash

# Применим миграции
# ВАЖНО! Для локального запуска вызвать из корня проекта, предварительно подгрузив переменные окружения в терминал командой  `set -a; source <путь к .env файлу>; set +a`
# Установка goose: go install github.com/pressly/goose/v3/cmd/goose@latest
goose -dir ./migrations postgres "user=$NOTIFICATION_SERVICE_POSTGRES_USER password=$NOTIFICATION_SERVICE_POSTGRES_PASSWORD dbname=$NOTIFICATION_SERVICE_POSTGRES_DB host=$NOTIFICATION_SERVICE_POSTGRES_HOST port=$NOTIFICATION_SERVICE_POSTGRES_PORT sslmode=disable" down