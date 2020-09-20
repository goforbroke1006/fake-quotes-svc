FROM golang:1.14 as build-env

WORKDIR /code/

COPY go.mod .
COPY go.sum .

#RUN GOPROXY=direct go mod download
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o fake-quotes-svc ./

FROM debian:stretch

COPY --from=build-env /code/fake-quotes-svc /usr/local/bin/fake-quotes-svc
COPY ./config.yaml.dist /app/config.yaml

WORKDIR /app/

ENTRYPOINT [ "fake-quotes-svc" ]

EXPOSE 8080
