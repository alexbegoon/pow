ARG GO_VERSION=1.18
FROM golang:${GO_VERSION} as base

ENV APPPATH="/application"
ENV GOPATH="/usr/go"
ENV GO111MODULE="on"

COPY . ${APPPATH}
WORKDIR ${APPPATH}

FROM base as dev

RUN curl -sSfl https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin/

FROM dev as dev-server

CMD ${GOPATH}/bin/air -c "${APPPATH}/configs/server.air.toml"

FROM dev as dev-client

CMD ${GOPATH}/bin/air -c "${APPPATH}/configs/client.air.toml"

FROM base as server

RUN go build -mod readonly -a -o server cmd/server/main.go

FROM base as client

RUN go build -mod readonly -a -o client cmd/client/main.go