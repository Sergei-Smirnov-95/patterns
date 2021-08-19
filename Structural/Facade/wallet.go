package main

type wallet struct {
	balance int
}

func NewWallet() *wallet {
	return &wallet{balance: 0}
}

func (w *wallet) addBalance(amount int){
	w.balance+=amount
}