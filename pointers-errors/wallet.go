package pointerserrors

import (
	"errors"
	"fmt"
)

var ErrInsufficientBalance = errors.New("insufficient balance")

type Bitcoin int

func (bitcoin Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", bitcoin)
}

type Wallet struct {
	balance Bitcoin
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return ErrInsufficientBalance
	}
	wallet.balance -= amount
	return nil
}

func (wallet *Wallet) Balance() (balance Bitcoin) {
	return wallet.balance
}
