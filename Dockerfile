# Estágio 1: Compilação
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
# Compila o binário estático
RUN CGO_ENABLED=0 GOOS=linux go build -o rate-limiter ./cmd/main.go

# Estágio 2: Execução
FROM alpine:latest
WORKDIR /root/
# Copia apenas o binário do estágio anterior
COPY --from=builder /app/rate-limiter .
COPY --from=builder /app/.env . 

EXPOSE 8080
CMD ["./rate-limiter"]