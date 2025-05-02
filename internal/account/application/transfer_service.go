package application

import (
	"context"

	"github.com/williamkoller/bank-transfer/internal/account/domain"
)

type AccountRepository interface {
	FindByID(ctx context.Context, id domain.AccountID) (*domain.Account, error)
	Save(ctx context.Context, account *domain.Account) error
}

type TransferService struct {
	repo AccountRepository
}

func NewTransferService(repo AccountRepository) *TransferService {
	return &TransferService{repo: repo}
}

func (s *TransferService) Transfer(ctx context.Context, fromID, toID domain.AccountID, amount float64) error {
	from, err := s.repo.FindByID(ctx, fromID)
	if err != nil {
		return err
	}
	to, err := s.repo.FindByID(ctx, toID)
	if err != nil {
		return err
	}

	if err := from.CanTransferTo(to, amount); err != nil {
		return err
	}

	from.Debit(amount)
	to.Credit(amount)

	if err := s.repo.Save(ctx, from); err != nil {
		return err
	}
	return s.repo.Save(ctx, to)
}
