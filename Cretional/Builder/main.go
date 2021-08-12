package main

import "fmt"

func main() {

	nb := getBuilder("normal")
	nb.setWindowType()
	nb.setDoorType()
	nb.setNumFloor()
	normalHouse := nb.getHouse()

	ib := getBuilder("igloo")
	director := newDirector(ib)
	iglooHouse := director.buildHouse()

	fmt.Println(normalHouse)
	fmt.Println(iglooHouse)

}
