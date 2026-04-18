package db

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	Client *redis.Client
}

func (r *RedisRepository) Increment(ctx context.Context, key string, expiration int) (int, error) {
	pipe := r.Client.Pipeline()
	incr := pipe.Incr(ctx, key)
	pipe.ExpireNX(ctx, key, time.Duration(expiration)*time.Second) // Expire em 60 segundos

	_, err := pipe.Exec(ctx)
	if err != nil {
		return 0, err
	}

	return int(incr.Val()), nil
}

func (r *RedisRepository) IsBlocked(ctx context.Context, key string) (bool, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil { // Chave não existe = Não está bloqueado
		return false, nil
	}
	if err != nil {
		return false, err // Erro real de conexão ou algo assim
	}
	return val == "blocked", nil
}

func (r *RedisRepository) Block(ctx context.Context, key string, duration int) error {
	return r.Client.Set(ctx, key, "blocked", time.Duration(duration)*time.Second).Err()
}
