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
}