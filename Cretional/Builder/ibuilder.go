package main

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(buildertype string) iBuilder {
	switch buildertype {
	case "normal":
		return &normalBuilder{}
	case "igloo":
		return &iglooBuilder{}
	default:
		return &normalBuilder{}
	}
}
