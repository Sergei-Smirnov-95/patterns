package main

import (
	"errors"
	"fmt"
)

type account struct{
	name string
}

func newAccount(accId string) *account{
	return &account{name: accId}
}

func (a *account) checkAccount(acc string) error{
	if a.name != acc {
		return errors.New("Wrong account")
	}
	fmt.Println("ok werification account")
	return nil
}