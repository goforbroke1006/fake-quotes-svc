FROM debian:stretch

COPY ./fake-quotes-svc /usr/local/bin/fake-quotes-svc

ENTRYPOINT fake-quotes-svc
