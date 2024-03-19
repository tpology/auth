FROM golang:1.22 AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod go install -v ./cmd/auth

FROM debian:bookworm-slim

COPY --from=builder /go/bin/auth /usr/local/bin/auth

ENTRYPOINT ["/usr/local/bin/auth"]
