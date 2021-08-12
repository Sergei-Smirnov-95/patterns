package main

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) setWindowType() {
	b.windowType = "snow window"
}

func (b *iglooBuilder) setDoorType() {
	b.doorType = "ice door"
}

func (b *iglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *iglooBuilder) getHouse() house {
	return house{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor}
}
