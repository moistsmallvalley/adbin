.PHONY: lint

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

