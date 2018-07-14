package main

// Bank represents a bank with accounts.
// Balances keyed by account number.
// Balances stored in minor units.
type Bank struct {
	accounts map[int64]int64
}

// NewBank returns a pointer to a new bank.
func NewBank() *Bank {
	return &Bank{
		accounts: make(map[int64]int64),
	}
}

// NewAccount creates a new bank account with the given account number and balance.
func (b *Bank) NewAccount(acct int64, balance int64) {
	b.accounts[acct] = balance
}

// Withdraw subtracts the given amount from the accounts balance. Returns true
// if the withdrawal was ok and false if there were insufficient funds.
func (b *Bank) Withdraw(acct int64, amount int64) bool {
	if b.accounts[acct] <= amount {
		return false
	}

	b.accounts[acct] -= amount

	return true
}

// GetBalance returns the current balance of the given account.
func (b *Bank) GetBalance(acct int64) int64 {
	return b.accounts[acct]
}
