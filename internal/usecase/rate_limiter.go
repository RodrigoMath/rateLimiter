package useCase

import (
	"context"
)

// func (u *RateLimiterUseCase) chooseRepositoryStrategy(strategy string) RateLimiterRepository {

// 	case strategy == "redis":
// 		return

// 	return u.Repository
// }

type RateLimiterRepository interface {
	Increment(ctx context.Context, key string, expiration int) (int, error)
	IsBlocked(ctx context.Context, key string) (bool, error)
	Block(ctx context.Context, key string, duration int) error
}

type RateLimiterUseCase struct {
	Repository RateLimiterRepository
	IPLimit    int
	TokenLimit int
	BlockTime  int
}

func NewRateLimiterUseCase(repo RateLimiterRepository, ipLimit, tokenLimit, blockTime int) *RateLimiterUseCase {
	return &RateLimiterUseCase{
		Repository: repo,
		IPLimit:    ipLimit,
		TokenLimit: tokenLimit,
		BlockTime:  blockTime,
	}
}

func (u *RateLimiterUseCase) Execute(ctx context.Context, id string, isToken bool) (bool, error) {
	// 1. Definir o limite baseado no tipo (Precedência)
	limit := u.IPLimit
	keyPrefix := "ip:"
	if isToken {
		limit = u.TokenLimit
		keyPrefix = "token:"
	}

	// 2. Chaves para o Redis
	countKey := keyPrefix + id
	blockKey := "block:" + keyPrefix + id

	// 3. Verificar se já está bloqueado (IsBlocked)
	blocked, err := u.Repository.IsBlocked(ctx, blockKey)
	if err != nil {
		return false, err
	}
	if blocked {
		return false, nil
	}

	// 4. Incrementar o contador (Increment)
	currentCount, err := u.Repository.Increment(ctx, countKey, u.BlockTime)
	if err != nil {
		return false, err
	}

	// 5. Se estourou o limite, aplicar o bloqueio (Block)
	if currentCount > limit {
		u.Repository.Block(ctx, blockKey, int(u.BlockTime))
		return false, nil
	}

	return true, nil
}
