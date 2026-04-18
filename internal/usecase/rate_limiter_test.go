package useCase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

// Implementando os métodos exigidos pela interface RateLimiterRepository
func (m *MockRepository) Increment(ctx context.Context, key string, expiration int) (int, error) {
	args := m.Called(ctx, key, expiration)
	return args.Int(0), args.Error(1)
}

func (m *MockRepository) IsBlocked(ctx context.Context, key string) (bool, error) {
	args := m.Called(ctx, key)
	return args.Bool(0), args.Error(1)
}

func (m *MockRepository) Block(ctx context.Context, key string, duration int) error {
	args := m.Called(ctx, key, duration)
	return args.Error(0)
}

func TestRateLimiterUseCase_shouldPass_Execute(t *testing.T) {

	mockRepo := new(MockRepository)

	useCase := NewRateLimiterUseCase(mockRepo, 10, 100, 60)

	ctx := context.Background()

	mockRepo.On("IsBlocked", ctx, mock.Anything).Return(false, nil)

	mockRepo.On("Increment", ctx, mock.Anything, mock.Anything).Return(9, nil)

	isBlocked, err := useCase.Execute(ctx, "127.0.0.1", false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	assert.True(t, isBlocked, "Expected to be allowed, but was blocked")

}

func TestRateLimiterUseCase_shouldFail_Execute(t *testing.T) {

	mockRepo := new(MockRepository)

	useCase := NewRateLimiterUseCase(mockRepo, 10, 100, 60)

	ctx := context.Background()

	mockRepo.On("IsBlocked", ctx, mock.Anything).Return(false, nil)

	mockRepo.On("Increment", ctx, mock.Anything, mock.Anything).Return(11, nil)

	mockRepo.On("Block", ctx, mock.Anything, mock.Anything).Return(nil)

	isBlocked, err := useCase.Execute(ctx, "127.0.0.1", false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	assert.False(t, isBlocked, "Expected to be not allowed")

}
