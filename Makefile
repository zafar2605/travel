
migration-up:
	migrate -path migrations/ -database "postgresql://zafar:2605@localhost:5432/for_migration?sslmode=disable" -verbose up

migration-down:
	migrate -path migrations/ -database "postgresql://zafar:2605@localhost:5432/for_migration?sslmode=disable" -verbose down

gen-swag:
	swag init -g api/api.go -o api/docs

run:
	go run cmd/main.go

