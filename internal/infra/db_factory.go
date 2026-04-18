package infra

import (
	"github.com/RodrigoMath/rateLimiter/config"
	"github.com/RodrigoMath/rateLimiter/internal/db"
	useCase "github.com/RodrigoMath/rateLimiter/internal/usecase"
	"github.com/redis/go-redis/v9"
)

func DbFactory(cfg *config.Config) useCase.RateLimiterRepository {
	switch cfg.Strategy {
	case "redis":
		// Criamos o cliente real aqui!
		client := redis.NewClient(&redis.Options{
			Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
			Password: cfg.RedisPassword,
			DB:       0,
		})

		return &db.RedisRepository{
			Client: client, // Agora não é mais nil
		}
	default:
		return nil
	}

}
