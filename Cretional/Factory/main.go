package main

import "fmt"

const (
	ak47 = iota
	moskit
)


func main(){
	ak47gun := getGunFactory(ak47)
	moskitgun := getGunFactory(moskit)
	fmt.Println(ak47gun.getName())
	fmt.Println(moskitgun.getPower())
}

func getGunFactory(gunType int)(iGun){
	switch gunType {
	case ak47:
		return newak47()
	case moskit:
		return newMoskitGun()
	default:
		return newak47()
	}
}