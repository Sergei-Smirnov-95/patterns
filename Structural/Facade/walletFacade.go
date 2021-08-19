package main

import "fmt"

type walletFacade struct{
	account *account
	wallet *wallet
	notification *notification
}

func NewWalletFacade(accountId string, code int) *walletFacade {
	walletFacade := &walletFacade{
		account: newAccount(accountId),
		wallet: NewWallet(),
		notification: &notification{},
	}
	fmt.Println("account created")
	return walletFacade
}

func (w *walletFacade) addMoneyToWallet(accountID string, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	w.wallet.addBalance(amount)
	w.notification.SendWalletDebitNot()
	return nil
}
