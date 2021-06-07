package pointers_and_errors

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) String() string {
	return fmt.Sprintf("wallet has %d coins", w.balance)
}
