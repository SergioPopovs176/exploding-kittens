APP_NAME = expkit

.PHONY: build run stop logs restart docker-build docker-up docker-down

# Docker: сборка образа
docker-build:
	docker compose build

# Docker: поднять сервисы в фоне
up:
	sudo docker compose up -d

# Docker: остановить сервисы и удалить контейнеры
down:
	sudo docker compose down

# Docker: перезапуск
restart: docker-down docker-up

# Логи
logs:
	docker compose logs -f
