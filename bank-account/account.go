package account

import (
	"sync"
)

// Account represents bank account
type Account struct {
	balance  int64
	isClosed bool
	mux      sync.Mutex
}

// Open account with initial deposit
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{
		balance: initialDeposit,
	}
}

// Close the account
func (a *Account) Close() (payout int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if a.isClosed || a.balance < 0 {
		return
	}

	payout, a.balance = a.balance, payout
	ok, a.isClosed = true, true
	return
}

// Balance gets balance of the account
func (a *Account) Balance() (balance int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if a.isClosed {
		return
	}

	return a.balance, true
}

// Deposit or withdraw amount
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if a.isClosed {
		return
	}

	if a.balance+amount < 0 {
		return
	}

	a.balance += amount
	return a.balance, true
}
