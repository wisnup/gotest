package pointers

import (
	"errors"
	"fmt"
)

var insufficientFundError = errors.New("insufficient fund")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {

	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
	// fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return insufficientFundError
	}
	w.balance -= amount
	// fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	return nil
}

func (w *Wallet) Balance() Bitcoin {

	return w.balance
}
