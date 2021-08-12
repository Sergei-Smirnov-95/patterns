package main

func main(){
	clientpc := &client{}
	macpc := &mac{}
	clientpc.insertLightingConnectorIntoComputer(macpc)

	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}

	clientpc.insertLightingConnectorIntoComputer(windowsMachineAdapter)
}
