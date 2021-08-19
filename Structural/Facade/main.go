package main

func main(){
	wallet := NewWalletFacade("pupkin",999)
	err := wallet.addMoneyToWallet("pupkin",100)
	if err != nil{
		panic(err.Error())
	}
}
