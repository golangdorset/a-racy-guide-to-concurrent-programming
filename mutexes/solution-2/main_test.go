package main

import "testing"

func BenchmarkBuyBeer(b *testing.B) {
	bank := NewBank()
	bank.NewAccount(1234, 50000)

	for i := 0; i < b.N; i++ {
		buyBeer(bank, 1234)
	}
}
