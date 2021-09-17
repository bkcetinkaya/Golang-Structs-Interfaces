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
		DepositBalance(&wallet,100)

		got:=PrintBalance(&wallet)
		want:= 150

		if got!=want{
			t.Errorf("we got: %d we wanted: %d",got,want)
		}
	})

	t.Run("Deposit Check for Dollar Account", func(t *testing.T) {
		//Create dollar wallet with 0 usd and then try to deposit 100 Eur to account and then get the balance amount to check if deposit was successful
		wallet := DollarWallet{ 0}
		DepositBalance(&wallet,100)

		got:=PrintBalance(&wallet)
		want:= 150

		if got!=want{
			t.Errorf("we got: %d we wanted: %d",got,want)
		}
	})

	t.Run("Withdraw Check for Eur Account", func(t *testing.T) {
		// Create an eur wallet with 100 in it, and then try to withdraw 50 eur
		wallet := EuroWallet{ 100}
		//Since our withdrawal method returns an error we can handle the error if it occurs
		err := wallet.Withdraw(50)
		//check if withdraw worked
		got:=PrintBalance(&wallet)
		want:= 50

		if err!=nil{
			t.Error(err)
			t.Errorf("we got: %d we wanted: %d",got,want)
		}
	})
	t.Run("Withdraw Check for usd Account", func(t *testing.T) {
		// Create an usd wallet with 100 in it, and then try to withdraw 50 usd
		wallet := DollarWallet{ 100}
		//Since our withdrawal method returns an error we can handle the error if it occurs
		err := wallet.Withdraw(50)
		//check if withdraw worked
		got:=PrintBalance(&wallet)
		want:= 50

		if err!=nil{
			t.Error(err)
		}

		if got!=want{
			t.Errorf("we got: %d we wanted: %d",got,want)
			t.Errorf("we got: %d we wanted: %d",got,want)
		}
	})




}