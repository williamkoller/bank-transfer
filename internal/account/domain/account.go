package domain

import (
	"errors"
	"time"

	"github.com/williamkoller/bank-transfer/pkg/uuid"
)

type AccountID = uuid.UUID

type Account struct {
	ID      AccountID
	Owner   string
	Balance float64
	Events  []DomainEvent
}

func NewAccount(id AccountID, owner string, initialBalance float64) (*Account, error) {
	if initialBalance < 0 {
		return nil, errors.New("initial balance cannot be negative")
	}
	return &Account{
		ID:      id,
		Owner:   owner,
		Balance: initialBalance,
	}, nil
}

func (a *Account) CanTransferTo(target *Account, amount float64) error {
	if a.ID == target.ID {
		return ErrCannotTransferToSelf
	}
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if a.Balance < amount {
		return ErrInsufficientFunds
	}
	return nil
}

func (a *Account) Debit(amount float64) {
	a.Balance -= amount
	a.Events = append(a.Events, NewAccountDebitedEvent(a.ID, amount, time.Now()))
}

func (a *Account) Credit(amount float64) {
	a.Balance += amount
	a.Events = append(a.Events, NewAccountCreditedEvent(a.ID, amount, time.Now()))
}
