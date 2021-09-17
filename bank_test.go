package Golang_Structs_Interfaces

import "testing"

func TestWallet(t *testing.T){
	t.Run("Balance Check For Euro Account", func(t *testing.T) {
		//Create a wallet with 100 Eur in it and check the balance if its correctly updating the balance
		wallet := EuroWallet{ 100}
		got := PrintBalance(&wallet)
		want:= 100

		if got!=want{
			t.Errorf("we got: %d we wanted: %d",got,want)
		}
	})

	t.Run("Balance Check For Dollar Account", func(t *testing.T) {
		//Create a wallet with 100 Eur in it and check the balance if its correctly updating the balance
		wallet := DollarWallet{ 100}
		got := PrintBalance(&wallet)
		want:= 100

		if got!=want{
			t.Errorf("we got: %d we wanted: %d",got,want)
		}
	})

	t.Run("Deposit Check for Euro Account", func(t *testing.T) {
		//Create euro wallet with 0 Eur and then try to deposit 100 Eur to account and then get the balance amount to check if deposit was successful
		wallet := EuroWallet{ 0}
		wallet.DepositBalance(150)

		got:=PrintBalance(&wallet)

		want:= 150

		if got!=want{
			t.Errorf("we got: %d we wanted: %d",got,want)
		}
	})

	t.Run("Deposit Check for Dollar Account", func(t *testing.T) {
		//Create dollar wallet with 0 usd and then try to deposit 100 Eur to account and then get the balance amount to check if deposit was successful
		wallet := DollarWallet{ 0}
		wallet.DepositBalance(150)

		got:=PrintBalance(&wallet)

		want:= 150

		if got!=want{
			t.Errorf("we got: %d we wanted: %d",got,want)
		}
	})

}