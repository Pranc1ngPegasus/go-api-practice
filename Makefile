.PHONY: wiregen go-build build up down start restart stop logs migrate-up migrate-down

BASE_DOCKER_COMPOSE = ./docker/docker-compose.yml
COMPOSE_OPTS = -f "$(BASE_DOCKER_COMPOSE)" -p go-api-practice

wiregen:
	third_party/bin/wire gen ./...

go-build: wiregen
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./.artifacts/server-linux-amd64 ./cmd/server

build: wiregen
	docker-compose $(COMPOSE_OPTS) up -d --build

setup: build

up:
	docker-compose $(COMPOSE_OPTS) up -d

down:
	docker-compose $(COMPOSE_OPTS) down

start:
	docker-compose $(COMPOSE_OPTS) start

restart:
	docker-compose $(COMPOSE_OPTS) restart

stop:
	docker-compose $(COMPOSE_OPTS) stop

logs:
	docker-compose $(COMPOSE_OPTS) logs -f

sqlboiler:
	docker-compose $(COMPOSE_OPTS) exec -T app sh -c "./third_party/bin/sqlboiler --wipe mysql"

migrate-up:
	docker-compose $(COMPOSE_OPTS) exec -T app sh -c "./third_party/bin/sql-migrate up -config=dbconfig.yml"

migrate-down:
	docker-compose $(COMPOSE_OPTS) exec -T app sh -c "./third_party/bin/sql-migrate down -config=dbconfig.yml"
