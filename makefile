DB_URL = mysql://root:nhn1832004@tcp(localhost:3306)/todo_app?charset=utf8mb4&parseTime=True

migrate-up:
	migrate -path ./internal/database/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path ./internal/database/migrations -database "$(DB_URL)" down

migrate-status:
	migrate -path ./internal/database/migrations -database "$(DB_URL)" version
