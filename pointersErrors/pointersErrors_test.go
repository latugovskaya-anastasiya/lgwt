package pointersErrors_test

import (
	"testing"

	"github.com/latugovskaya-anastasiya/lgwt/pointersErrors"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := pointersErrors.Wallet{}
		wallet.Deposit(pointersErrors.Bitcoin(10))

		assertBalance(t, wallet, pointersErrors.Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := pointersErrors.Wallet{}
		wallet.SetBalance(pointersErrors.Bitcoin(20))
		err := wallet.Withdraw(pointersErrors.Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, pointersErrors.Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {

		wallet := pointersErrors.Wallet{}
		wallet.SetBalance(pointersErrors.Bitcoin(20))
		err := wallet.Withdraw(pointersErrors.Bitcoin(100))

		assertError(t, err, pointersErrors.ErrInsufficientFunds)
		assertBalance(t, wallet, pointersErrors.Bitcoin(20))

	})
}
func assertBalance(t testing.TB, wallet pointersErrors.Wallet, want pointersErrors.Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
