BUILD_DATE=$(shell date +F%)
VERSION=v0.0.0
BIN=silverbullet
LD_FLAGS=-s -w

build:
	go build -ldflags "${LD_FLAGS} -X 'main.VERSION=${VERSION}' -X 'main.BUILD_DATE=${BUILD_DATE}'" \
		-buildmode=pie \
		-o ${BIN} cmd/${BIN}/*go
debug:
	go build -ldflags "-X 'main.VERSION=${VERSION}' -X 'main.BUILD_DATE=${BUILD_DATE}'" \
		-gcflags "-E" \
		-o ${BIN}-debug cmd/${BIN}/*go
check:
	go fmt cmd/${BIN}/*go
