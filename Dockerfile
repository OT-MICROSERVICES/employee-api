FROM golang:1.20 as builder

WORKDIR /go/src/employee-api

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download
COPY client/ client/
COPY config/ config/
COPY middleware/ middleware/
COPY model/ model/
COPY routes/ routes/
COPY api/ api/
COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o employee-api main.go

FROM alpine:3.18.0
LABEL authors="Opstree Solution" \
      contact="opensource@opstree.com" \
      version="v0.1.0"

WORKDIR /
COPY employee-api .
USER 65532:65532

ENTRYPOINT ["/employee-api"]
