.PHONY: run test db-start db-init db-shell lint linters

run:
	go run . -dbname sample -password example

test:
	go test ./...

db-start:
	docker-compose up -d testdb

db-init:
	docker-compose down -v
	docker-compose up -d testdb

db-shell:
	docker-compose exec testdb mysql -u root -pexample sample

lint:
	docker run --rm -t \
		-v $$(pwd):/app \
		-v $$(pwd)/.cache/golangci-lint/v1.52.1:/root/.cache \
		-w /app \
		golangci/golangci-lint:v1.52.1 \
		golangci-lint run -v

linters:
	docker run --rm -t \
		golangci/golangci-lint:v1.52.1 \
		golangci-lint help linters

