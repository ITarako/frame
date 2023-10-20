include .env

run-api:
	go run ./cmd/api

build-api:
	@echo 'Building cmd/api...'
	go build -ldflags='-s' -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux_amd64/api ./cmd/api

migration-create:
	@echo 'Creating migration file for ${name}'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

migration-up:
	@echo 'Up migrations'
	migrate -path=./migrations -database=${PROJECT_DB_DSN} -verbose up

migration-down:
	@echo 'Down migrations ${count}'
	migrate -path=./migrations -database=${PROJECT_DB_DSN} -verbose down ${count}

migration-fix:
	@echo 'Fix migrations of version ${version}'
	migrate -path=./migrations -database=${PROJECT_DB_DSN} force ${version}

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