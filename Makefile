create-migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name

migrate-up:
	migrate -database postgres://postgres:password@localhost:5432/dealls?sslmode=disable -path ./migrations up

migrate-down:
	migrate -database postgres://postgres:password@localhost:5432/dealls?sslmode=disable -path ./migrations down

up-db:
	docker run --name dealls-db \
		-e POSTGRES_DB=dealls \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-p 5432:5432 \
		-d postgres

gen-doc:
	swag init --parseInternal

test:
	k6 run e2e/script.js

run:
	go run main.go