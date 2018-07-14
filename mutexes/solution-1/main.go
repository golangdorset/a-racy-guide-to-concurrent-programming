package main

import "fmt"

var b *Bank

const (
	beerPrice = 450
	acctDan   = 1234
	acctBen   = 4321
)

func main() {
	// Create the bank.
	b = NewBank()

	// Set up dan's account.
	b.NewAccount(acctDan, 5000)

	// Set up bens's account.
	b.NewAccount(acctBen, 5000)

	// Dan buys 5 beers.
	for i := 0; i < 5; i++ {
		go buyBeer(b, acctDan)
	}

	// Ben is sensible and buys only 3 beers.
	for i := 0; i < 3; i++ {
		go buyBeer(b, acctBen)
	}
}

// buyBeer withdraws the cost of a beer from the specified bank account.
func buyBeer(b *Bank, acct int64) {
	if !b.Withdraw(acct, beerPrice) {
		fmt.Println("Broke!")
		return
	}

	fmt.Printf("Account %d has £%4.2f remaining\n", acct, float64(b.GetBalance(acct))/100)
}
