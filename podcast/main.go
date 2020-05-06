package main

import (
	"github.com/jwhitake/podcast/config"
	"github.com/jwhitake/podcast/files"
	"fmt"
	"log"
	"path"
)

func main() {
	// filePath := "C:/todel"

	// absolutePath := path.Join(filePath, "test.txt")
	// fmt.Println(absolutePath)

	// testBase := path.Base(absolutePath)
	// fmt.Println(testBase)

	// testDir := path.Dir(absolutePath)
	// fmt.Println(testDir)

	// testExt := path.Ext(absolutePath)
	// fmt.Println(testExt)

	// dir, file := path.Split(absolutePath)
	// fmt.Println(dir)
	// fmt.Println(file)

	podPath := config.PodPath
	fmt.Println(podPath)

	dirList, err := files.GetDirectoryContents(podPath)
	if err != nil {
		log.Fatal(err)
	}
	for index, list := range dirList {
		fmt.Println(list)
		dirName := dirList[index]
		absolutePath := path.Join(podPath, dirName)
		fileList, err := files.GetDirectoryContents(absolutePath)
		if err != nil{
			log.Fatal(err)
		}
		for i := range fileList{
			fileName := fileList[i]
			fmt.Println(fileName)
		}
		
	}

	// lines, err := files.ReadFile(path)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, line := range lines {
	// 	fmt.Println(line)
	// }
}
