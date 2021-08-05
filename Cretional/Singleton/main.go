package main

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}

type singleStruct struct{}

var single  *singleStruct

func getSingle() *singleStruct {
	if single == nil {
		mutex.Lock()
		defer  mutex.Unlock()
		if single == nil {
			single = &singleStruct{}
			fmt.Println("Now we create struct")
		}
		fmt.Println("struct already exist")
		return single
	}
	fmt.Println("struct already exist")
	return single
}

func main()  {
	for i:=0; i<10; i++ {
		go getSingle()
	}
	fmt.Scanln()
}