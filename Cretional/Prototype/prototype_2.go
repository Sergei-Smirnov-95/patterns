package main

import "fmt"

type folder struct {
	name string
	children []inode
}

func (f *folder) print(s string) {
	fmt.Println(s+f.name)
	for _,ch := range f.children {
		ch.print(s+"	")
	}
}

func (f *folder)Clone() inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var childrens []inode
	for _,ch := range f.children {
		copy := ch.Clone()
		childrens = append(childrens, copy)
	}
	cloneFolder.children = childrens
	return cloneFolder
}