FROM golang:1.23.4-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o grpc-server ./cmd/main.go


FROM alpine:3.21.3

WORKDIR /app

COPY --from=builder /build/grpc-server .
RUN chmod +x ./grpc-server

COPY migrations ./migrations
COPY proto ./proto

EXPOSE 50051

CMD ["./grpc-server"]
