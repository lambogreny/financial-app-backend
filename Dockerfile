FROM golang:1.13.4 AS builder
ENV DATA_DIRECTORY /go/src/9phum.com/financial-app-backend
WORKDIR $DATA_DIRECTORY
ARG APP_VERSION
ARG CGO_ENABLED=0
COPY . .
RUN go build -ldflags="-X 9phum.com/financial-app-backend/internal/config.Version=$APP_VERSION" 9phum.com/financial-app-backend/cmd/server

FROM alpine:3.10
ENV DATA_DIRECTORY /go/src/9phum.com/financial-app-backend/
RUN apk add --update --no-cache \
    ca-certificates
COPY internal/database/migrations ${DATA_DIRECTORY}internal/database/migrations   
COPY --from=builder ${DATA_DIRECTORY}server /financial-app-backend

ENTRYPOINT ["/financial-app-backend"]
