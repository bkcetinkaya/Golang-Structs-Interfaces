package Golang_Structs_Interfaces

type EuroWallet struct {
	euroBalance int
}

type DollarWallet struct {
	dollarBalance int
}

type Currency interface{
	CheckBalance() int
}

func (euroWallet *EuroWallet) CheckBalance() int{
	return euroWallet.euroBalance
}

func(dollarWallet *DollarWallet) CheckBalance() int{
	return dollarWallet.dollarBalance
}

func PrintBalance(c Currency) int {
	return c.CheckBalance()
}

