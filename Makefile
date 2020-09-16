REGISTRY=docker.io
OWNER=goforbroke1006
SERVICE_NAME=fake-quotes-svc

VERSION ?= $(shell git describe --abbrev=0 --tags 2> /dev/null | git rev-parse --abbrev-ref HEAD 2> /dev/null | echo 'unknown')
HASH ?= $(shell git rev-parse --verify HEAD 2> /dev/null | echo 'unknown')
BUILD_DATE = $(shell date -u '+%Y-%m-%d %H:%M:%S %Z')

all: dep build lint test

dep:
	GOPROXY=direct go mod download

build:
	go build -ldflags "-s -w -X 'main.date=${BUILD_DATE}' -X main.version=${VERSION} -X main.commit=${HASH}" -o ${SERVICE_NAME} ./

lint:
	golangci-lint run

test:
	go test -race ./...

image:
	docker build -t ${REGISTRY}/${OWNER}/${SERVICE_NAME}:${VERSION}	./
	docker build -t ${REGISTRY}/${OWNER}/${SERVICE_NAME}:latest		./
	docker push ${REGISTRY}/${OWNER}/${SERVICE_NAME}:${VERSION}
	docker push ${REGISTRY}/${OWNER}/${SERVICE_NAME}:latest
