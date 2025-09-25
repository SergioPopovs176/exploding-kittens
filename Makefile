APP_NAME = exploding-kittens

.PHONY: logs restart build up down db-shell

# Docker: сборка образа
build:
	docker compose build

# Docker: поднять сервисы в фоне
up:
	@echo "Starting services (Go + Postgres)..."
	sudo docker compose up -d

# Docker: остановить сервисы и удалить контейнеры
down:
	@echo "Stopping services..."
	sudo docker compose down

# Docker: перезапуск
restart: docker-down docker-up

# Логи
logs:
	sudo docker compose logs -f

# Подключиться в psql внутри контейнера базы
db-shell:
	@echo "🐘 Connecting to Postgres..."
	sudo docker compose -p $(APP_NAME) exec db psql -U postgres -d expkit
