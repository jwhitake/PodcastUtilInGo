package main

import (
	"fmt"
	"path"
)

func main() {
	filePath := "C:/todel"

	absolutePath := path.Join(filePath, "test.txt")
	fmt.Println(absolutePath)

	testBase := path.Base(absolutePath)
	fmt.Println(testBase)

	testDir := path.Dir(absolutePath)
	fmt.Println(testDir)

	testExt := path.Ext(absolutePath)
	fmt.Println(testExt)

	dir, file := path.Split(absolutePath)
	fmt.Println(dir)
	fmt.Println(file)

	// path := "c:/todel/test.txt"
	// lines, err := files.ReadFile(path)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, line := range lines {
	// 	fmt.Println(line)
	// }

	// path = "C:/todel/test/"
	// dirList, err := files.GetDirectoryContents(path)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, list := range dirList {
	// 	fmt.Println(list)
	// }
}
