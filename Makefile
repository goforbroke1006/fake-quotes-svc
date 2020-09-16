REGISTRY=docker.io
OWNER=goforbroke1006
SERVICE_NAME=fake-quotes-svc

VERSION ?= $(shell git describe --abbrev=0 --tags | git rev-parse --abbrev-ref HEAD | echo '0.1.0')
HASH ?= $(shell git rev-parse --verify HEAD | echo 'unknown')
BUILD_DATE = $(shell date -u '+%Y-%m-%d %H:%M:%S %Z')

all: build lint test

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
