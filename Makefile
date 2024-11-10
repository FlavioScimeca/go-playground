.PHONY: dev build clean docker-up docker-down migrate docker-build

dev:
	go run cmd/main.go

build:
	go build -o bin/server cmd/main.go

clean:
	rm -rf bin/

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down -v

docker-build:
	docker-compose build

migrate:
	go run cmd/main.go migrate