package Golang_Structs_Interfaces

import "errors"

type EuroWallet struct {
	euroBalance int
}

type DollarWallet struct {
	dollarBalance int
}

type Currency interface{
	CheckBalance() int
	DepositBalance(amount int)
	Withdraw(amount int) error
}

func (euroWallet *EuroWallet) CheckBalance() int{
	return euroWallet.euroBalance
}

func(dollarWallet *DollarWallet) CheckBalance() int{
	return dollarWallet.dollarBalance
}

func (euroWallet *EuroWallet) DepositBalance(amount int){
	euroWallet.euroBalance+=amount
}

func(dollarWallet *DollarWallet) DepositBalance(amount int) {
	dollarWallet.dollarBalance+=amount
}

func (euroWallet *EuroWallet) Withdraw(amount int) error{
	if amount > euroWallet.CheckBalance(){
		return errors.New("can not withdraw, insufficient funds")
	}
	euroWallet.euroBalance-=amount
	return nil
}

func(dollarWallet *DollarWallet) Withdraw(amount int) error {
	if amount > dollarWallet.CheckBalance(){
		return errors.New("can not withdraw, insufficient funds")
	}
	dollarWallet.dollarBalance-=amount
	return nil

}

func PrintBalance(c Currency) int {
	return c.CheckBalance()
}

func DepositBalance(c Currency,amount int){
	c.DepositBalance(amount)
}
