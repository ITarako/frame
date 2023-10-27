include .env

ifdef ${.CURDIR}
	CURRENT_DIR = ${.CURDIR}
else
	CURRENT_DIR = ${CURDIR}
endif
export TOOLS_DIR=${CURRENT_DIR}/bin/tools

gen:
	go generate ${CURRENT_DIR}/internal/...

run-api:
	go run ${CURRENT_DIR}/cmd/api

build-api:
	@echo 'Building cmd/api...'
	go build -ldflags='-s' -o=${CURRENT_DIR}/bin/api ${CURRENT_DIR}/cmd/api

migrate-create:
	@echo 'Creating migration file for ${name}'
	${TOOLS_DIR}/migrate create -seq -ext=.sql -dir=${CURRENT_DIR}/migrations ${name}

migrate-up:
	@echo 'Up migrations'
	${TOOLS_DIR}/migrate -path=${CURRENT_DIR}/migrations -database=${DB_DSN} -verbose up

migrate-down:
	@echo 'Down migrations ${count}'
	${TOOLS_DIR}/migrate -path=${CURRENT_DIR}/migrations -database=${DB_DSN} -verbose down ${count}

migrate-fix:
	@echo 'Fix migrations of version ${version}'
	${TOOLS_DIR}/migrate -path=${CURRENT_DIR}/migrations -database=${DB_DSN} force ${version}

audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ${CURRENT_DIR}/...
	@echo 'Vetting code...'
	go vet ${CURRENT_DIR}/...
	staticcheck ${CURRENT_DIR}/...
	@echo 'Running tests...'
	go test -race -count=1 -vet=off ${CURRENT_DIR}/...