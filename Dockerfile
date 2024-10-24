FROM golang:latest as build

ENV CGP_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN useradd -u 10001 raven

ENV SRCDIR=/go/src

WORKDIR $SRCDIR

RUN apt-get update && apt-get install -y ca-certificates openssl tzdata
RUN update-ca-certificates

COPY . $SRCDIR

RUN make build

FROM busybox as package

WORKDIR /

COPY --from=build /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /go/src /

USER raven

EXPOSE 8080

CMD ["/main"]