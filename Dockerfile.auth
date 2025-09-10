FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o auth ./cmd/auth 

FROM scratch

WORKDIR /app

COPY --from=builder /app/auth .

EXPOSE 8081

CMD ["./auth"]