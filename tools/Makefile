ifdef ${.CURDIR}
	CURRENT_DIR = ${.CURDIR}
else
	CURRENT_DIR = ${CURDIR}
endif

gen:
	go generate ${CURRENT_DIR}/tools.go

audit:
	go mod tidy
	go mod verify