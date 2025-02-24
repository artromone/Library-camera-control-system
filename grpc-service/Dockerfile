FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o grpc-server ./cmd/main.go

FROM alpine:latest 

WORKDIR /app

COPY --from=builder /app/grpc-server .

COPY migrations ./migrations

COPY proto ./proto

EXPOSE 50051

CMD ["./grpc-server"]
