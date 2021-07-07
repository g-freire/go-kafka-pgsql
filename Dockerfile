#FROM golang:1.14.2-alpine as builder
#RUN apk add alpine-sdk
#WORKDIR ./go-api/
#COPY . .
#RUN go mod download
#RUN GOOS=linux GOARCH=amd64 go build -o rest-api -tags musl
#
#FROM alpine:latest as runner
#WORKDIR /root/
#COPY --from=builder /go/app/rest-api .
#ENTRYPOINT /root/rest-api


# first stage
FROM golang:1.14.2-stretch as build_base
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download

# second stage
FROM build_base AS builder
WORKDIR ./
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags netgo -o /bin/app -ldflags "-w -s -X" ./cmd/**/main.go

# final stage
FROM debian:10.3-slim
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
COPY --from=builder /bin/app /bin/app
EXPOSE 80
ENTRYPOINT ["/bin/app"]