package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertNoError := func(t testing.TB, err error) {
		t.Helper()

		if err != nil {
			t.Fatal("wanted no error but did get one")
		}
	}

	assertError := func(t testing.TB, err error, want error) {
		t.Helper()

		if err == nil {
			t.Fatal("wanted error but didn't get one")
		}

		if err != want {
			t.Errorf("wanted %s but got %s", want, err)
		}
	}

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", wallet.balance, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		// fmt.Printf("the address of balance in test is %p \n", &wallet.balance)

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {

		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		// fmt.Printf("the address of balance in test is %p \n", &wallet.balance)

		want := Bitcoin(10)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient fund", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, insufficientFundError)
		assertBalance(t, wallet, startingBalance)
	})
}
