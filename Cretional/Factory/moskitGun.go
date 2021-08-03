package main

type moskitGun struct {
	gun
	superPower int
}

func newMoskitGun() iGun{
	return &moskitGun{
		gun{"ak47",10,},
		0,
	}
}