# Stage 1: build
FROM golang:1.25-alpine AS builder

WORKDIR /build

# Install git (needed for go modules)
RUN apk add --no-cache git

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd


# Stage 2: minimal runtime container
FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/server .

EXPOSE 8000

CMD ["./server"]
