package domain

import "time"

type DomainEvent interface {
	EventName() string
}

type AccountDebited struct {
	AccountID AccountID
	Amount    float64
	At        time.Time
}

func NewAccountDebitedEvent(id AccountID, amount float64, at time.Time) *AccountDebited {
	return &AccountDebited{AccountID: id, Amount: amount, At: at}
}

func (e *AccountDebited) EventName() string {
	return "AccountDebited"
}

type AccountCredited struct {
	AccountID AccountID
	Amount    float64
	At        time.Time
}

func NewAccountCreditedEvent(id AccountID, amount float64, at time.Time) *AccountCredited {
	return &AccountCredited{AccountID: id, Amount: amount, At: at}
}

func (e *AccountCredited) EventName() string {
	return "AccountCredited"
}
