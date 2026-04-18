package main

import (
	"fmt"
	"net/http"

	config "github.com/RodrigoMath/rateLimiter/config"
	infra "github.com/RodrigoMath/rateLimiter/internal/infra"
	rateLimiter "github.com/RodrigoMath/rateLimiter/internal/usecase"
)

func main() {
	mux := http.NewServeMux()

	cfg, err := config.LoadConfig()
	dbRepo := infra.DbFactory(cfg)

	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá! Você passou pelo Rate Limiter!"))
	})

	useCase := rateLimiter.NewRateLimiterUseCase(dbRepo, cfg.LimitIp, cfg.LimitToken, cfg.BlockTime)

	// Aplicando o middleware
	middleware := infra.NewRateLimiterMiddleware(useCase)

	handlerFinal := middleware(mux)

	fmt.Println("Servidor rodando na porta :8080...")
	http.ListenAndServe(":8080", handlerFinal)
}
