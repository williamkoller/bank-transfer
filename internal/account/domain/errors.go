package domain

import "errors"

var (
	ErrCannotTransferToSelf = errors.New("cannot transfer to same account")
	ErrInsufficientFunds    = errors.New("insufficient funds")
	ErrInvalidAmount        = errors.New("amount must be greater than zero")
)
