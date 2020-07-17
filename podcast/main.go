package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"

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


	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)


	currentDir := files.GetCurrentDirectory()
	fmt.Println(currentDir + "/menu.txt")
	menu, err := files.ReadFile(currentDir + "/menu.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Point 1")
	for _, lineItem := range menu {
		fmt.Println(lineItem)
	}
	// reader := bufio.NewReader(os.Stdin)
	// text, _ := reader.ReadString('\n')
	// text = strings.Replace(text, "\n", "", -1)
	// fmt.Println(text)
	// test, err := files.Detect()
	// for _, usbList := range test {
	// 	fmt.Println(usbList)
	// }
	fmt.Println("Point 2")
	configuration := config.GetConfig()
	fmt.Println("Point 3")
	fmt.Println(configuration.ArchiveFiles)
	fmt.Println(configuration.ArchivePath)
	fmt.Println(configuration.SourcePath)
	// dirList, err := files.GetSubDirectories(configuration.SourcePath)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for index, list := range dirList {
	// 	fmt.Println(list)
	// 	dirName := dirList[index]
	// 	absolutePath := path.Join(configuration.SourcePath, dirName)
	// 	fileList, err := files.GetDirectoryFileNames(absolutePath)
	// 	fmt.Println(absolutePath)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	for i := range fileList {
	// 		fileName := fileList[i]
	// 		fmt.Println(fileName)

	// 		lines, err := files.ReadFile(path.Join(absolutePath, fileName))
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		for _, line := range lines {
	// 			fmt.Println(line)
	// 		}
	// 	}
	// }
}
