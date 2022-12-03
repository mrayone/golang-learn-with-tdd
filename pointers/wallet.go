package pointers

import (
	"errors"
	"fmt"
)

var ErrorInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	//(*w).balance way two
	return w.balance //automatically dereference that pointer as (*w).balance
}

func (w *Wallet) Withdraw(amount Bitcoin) (Bitcoin, error) {
	if amount > w.balance {
		return Bitcoin(0), ErrorInsufficientFunds
	}

	w.balance -= amount
	return amount, nil
}
