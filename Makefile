include .env
MOCKERY_PATH:=$(shell pwd)/bin/

run-api:
	go run ./cmd/api

build-api:
	@echo 'Building cmd/api...'
	go build -ldflags='-s' -o=./bin/api ./cmd/api

migrate-create:
	@echo 'Creating migration file for ${name}'
	./bin/tools/migrate create -seq -ext=.sql -dir=./migrations ${name}

migrate-up:
	@echo 'Up migrations'
	./bin/tools/migrate -path=./migrations -database=${DB_DSN} -verbose up

migrate-down:
	@echo 'Down migrations ${count}'
	./bin/tools/migrate -path=./migrations -database=${DB_DSN} -verbose down ${count}

migrate-fix:
	@echo 'Fix migrations of version ${version}'
	./bin/tools/migrate -path=./migrations -database=${DB_DSN} force ${version}

audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -count=1 -vet=off ./...

gen:
	go generate ./interal/...

gen-tools:
	go generate ./tools/tools.go

test:
	@echo '${MOCKERY_PATH}'