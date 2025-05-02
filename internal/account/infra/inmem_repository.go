package infra

import (
	"context"
	"errors"
	"sync"

	"github.com/williamkoller/bank-transfer/internal/account/domain"
)

type InMemoryAccountRepo struct {
	data map[domain.AccountID]*domain.Account
	mu   sync.RWMutex
}

func NewInMemoryAccountRepo() *InMemoryAccountRepo {
	return &InMemoryAccountRepo{
		data: make(map[domain.AccountID]*domain.Account),
	}
}

func (r *InMemoryAccountRepo) FindByID(ctx context.Context, id domain.AccountID) (*domain.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	acc, ok := r.data[id]
	if !ok {
		return nil, errors.New("account not found")
	}
	return acc, nil
}

func (r *InMemoryAccountRepo) Save(ctx context.Context, account *domain.Account) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[account.ID] = account
	return nil
}
