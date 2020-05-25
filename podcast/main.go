package main

import (
	"fmt"
	"log"
	"path"

	"github.com/jwhitake/podcast/config"
	"github.com/jwhitake/podcast/files"
)

func main() {

	// test, err := files.Detect()
	// for _, usbList := range test {
	// 	fmt.Println(usbList)
	// }

	podPath := config.PodPath
	fmt.Println(podPath)

	dirList, err := files.GetSubDirectories(podPath)
	if err != nil {
		log.Fatal(err)
	}
	for index, list := range dirList {
		fmt.Println(list)
		dirName := dirList[index]
		absolutePath := path.Join(podPath, dirName)
		fileList, err := files.GetDirectoryFileNames(absolutePath)
		fmt.Println(absolutePath)
		if err != nil {
			log.Fatal(err)
		}
		for i := range fileList {
			fileName := fileList[i]
			fmt.Println(fileName)

			lines, err := files.ReadFile(path.Join(absolutePath, fileName))
			if err != nil {
				log.Fatal(err)
			}
			for _, line := range lines {
				fmt.Println(line)
			}
		}
	}
}
