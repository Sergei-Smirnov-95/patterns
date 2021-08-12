package main

import "fmt"

func main() {
	file1 := &file{"file_1"}
	file2 := &file{"file_2"}
	folder1 := &folder{name: "folder_1",
		children: []inode{file1},
	}
	folder2 := &folder{name: "folder_2",
		children: []inode{folder1, file2},
	}

	fmt.Println("print object:")
	folder2.print("")
	clone := folder2.Clone()
	fmt.Println("print clone:")
	clone.print("")
}
