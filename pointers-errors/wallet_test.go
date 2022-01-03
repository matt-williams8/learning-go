package pointerserrors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("can accept deposits", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertWalletBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		t.Run("withdraws when there is available balance", func(t *testing.T) {
			wallet := Wallet{balance: 100}

			err := wallet.Withdraw(Bitcoin(20))

			assertNoError(t, err)
			assertWalletBalance(t, wallet, Bitcoin(80))
		})

		t.Run("errors when there is insufficent balance", func(t *testing.T) {
			wallet := Wallet{balance: 20}
			err := wallet.Withdraw(Bitcoin(50))

			assertError(t, err, ErrInsufficientBalance)
			assertWalletBalance(t, wallet, Bitcoin(20))

		})
	})
}

func assertWalletBalance(tb testing.TB, wallet Wallet, expectedBalance Bitcoin) {
	tb.Helper()
	actualBalance := wallet.Balance()
	if expectedBalance != actualBalance {
		tb.Errorf("expected %v, got %v", expectedBalance, actualBalance)
	}
}

func assertNoError(tb testing.TB, err error) {
	if err != nil {
		tb.Errorf("expected no error, got %v", err)
	}
}

func assertError(tb testing.TB, err, expectedError error) {
	tb.Helper()
	if err == nil {
		tb.Fatal("expected an error but did not get one")
	}

	if expectedError != err {
		tb.Errorf("expected error %v, got %v", expectedError, err)
	}
}
