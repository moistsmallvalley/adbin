.PHONY: run
run:
	go run . -dbname sample -password example

.PHONY: test
test:
	go test ./...

.PHONY: db/start
db/start:
	docker-compose up -d testdb

.PHONY: db/init
db/init:
	docker-compose down -v
	docker-compose up -d testdb

.PHONY: db/shell
db/shell:
	docker-compose exec testdb mysql -u root -pexample sample

.PHONY: lint
lint:
	docker run --rm -t \
		-v $$(pwd):/app \
		-v $$(pwd)/.cache/golangci-lint/v1.52.1:/root/.cache \
		-w /app \
		golangci/golangci-lint:v1.52.1 \
		golangci-lint run -v

.PHONY: linters
linters:
	docker run --rm -t \
		golangci/golangci-lint:v1.52.1 \
		golangci-lint help linters

