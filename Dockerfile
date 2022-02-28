# syntax=docker/dockerfile:1

FROM golang AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /istio-unittest-time-server

##
## Deploy
##
FROM alpine

WORKDIR /

COPY --from=build /istio-unittest-time-server /istio-unittest-time-server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/istio-unittest-time-server"]
