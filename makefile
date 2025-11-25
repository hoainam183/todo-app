include .env
export
# Tạo DB_URL từ các biến
DB_URL = mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?charset=utf8mb4&parseTime=True

migrate-up:
	@migrate -path ./internal/database/migrations -database "$(DB_URL)" up
migrate-down:
	migrate -path ./internal/database/migrations -database "$(DB_URL)" down

migrate-status:
	migrate -path ./internal/database/migrations -database "$(DB_URL)" version
