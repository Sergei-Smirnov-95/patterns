package main

import "fmt"

type client struct{
	name string
}

func (c *client) insertLightingConnectorIntoComputer(com computer){
	fmt.Println("client inserts Lighting connector into computer")
	com.insertIntoLightingPort()
}