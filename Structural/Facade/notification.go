package main

import "fmt"

type notification struct {
}

func (n *notification) SendWalletDebitNot(){
	fmt.Println("sending wallet debit  notification")
}

func (n *notification) SendWalletCreditNot(){
	fmt.Println("sending wallet credit notification")
}