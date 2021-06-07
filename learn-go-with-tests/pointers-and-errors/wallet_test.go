package pointers_and_errors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		result := wallet.Balance()
		expected := 10

		fmt.Println(&wallet)

		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})

	// not done

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{10}

		wallet.Deposit(10)

		result := wallet.Balance()
		expected := 10

		fmt.Println(&wallet)

		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})
}
