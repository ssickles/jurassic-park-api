# syntax = docker/dockerfile:1

FROM golang:1.21-alpine AS builder
WORKDIR /src
COPY go.* ./
RUN go mod download
COPY . ./
RUN --mount=type=cache,target=/root/.cache/go-build go build -o jurassic-park-api .

FROM alpine:latest

# Configure entrypoint
WORKDIR /
COPY docker-entrypoint.sh .
RUN chmod +x docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]

# Set up goose and migration sql files
ADD "https://github.com/pressly/goose/releases/download/v3.6.1/goose_linux_x86_64" /goose
RUN chmod +x /goose
WORKDIR /app/migrations/
COPY ./migrations/ .

# Add jurassic-park-api binary
WORKDIR /app/
COPY --from=builder /src/jurassic-park-api ./

CMD ["/app/jurassic-park-api"]
