FROM debian:stretch

COPY ./fake-quotes-svc /usr/local/bin/fake-quotes-svc
COPY ./config.yaml /app/config.yaml

WORKDIR /app/

ENTRYPOINT [ "fake-quotes-svc" ]

EXPOSE 8080
