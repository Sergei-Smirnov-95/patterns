package main

type inode interface {
	print(string)
	Clone() inode
}
