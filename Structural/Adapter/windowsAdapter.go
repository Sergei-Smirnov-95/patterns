package main

import "fmt"

type windowsAdapter struct{
	windowMachine *windows
}

func (w *windowsAdapter) insertIntoLightingPort()  {
	fmt.Println("Adapter converts Lighting signal to USB")
	w.windowMachine.insertIntoUSBPort()
}
