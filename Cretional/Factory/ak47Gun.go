package main

type ak47Gun struct {
	gun
	familyName string
}

func (a *ak47Gun)getfamilyName() string{
	return a.familyName
}

func newak47() iGun{
	return &ak47Gun{
		gun{"ak47",10,},
		"Kalashnikov",
	}
}