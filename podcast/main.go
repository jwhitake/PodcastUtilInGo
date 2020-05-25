package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/jwhitake/podcast/config"
	"github.com/jwhitake/podcast/files"
)

//Configuration holds json config
type Configuration struct {
	SourcePath   string `json:"SourcePath"`
	ArchivePath  string `json:"ArchivePath"`
	ArchiveFiles string `json:"ArchiveFiles"`
}

func main() {
	// test, err := files.Detect()
	// for _, usbList := range test {
	// 	fmt.Println(usbList)
	// }

    configuration := config.GetConfig()
	dirList, err := files.GetSubDirectories(configuration.SourcePath)
	if err != nil {
		log.Fatal(err)
	}
	for index, list := range dirList {
		fmt.Println(list)
		dirName := dirList[index]
		absolutePath := path.Join(configuration.SourcePath, dirName)
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
