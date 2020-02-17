package files

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

//ReadFile is for reading in and returning a file from the passed
//in path variable
func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

//GetDirectoryContents is for
func GetDirectoryContents(path string) ([]string, error) {
	listing, err := ioutil.ReadDir(path)
	var dirList []string
	for _, dir := range listing {
		if dir.IsDir() {
			dirList = append(dirList, dir.Name())
		}
	}
	return dirList, err
}

func DelteFilesByExtension(path string, ext string)(bool, error){
	listing, err := ioutil.ReadDir(path)
	for _, f := range listing{
		if f.Name := strings.HasSuffix(ext){
			os.Remove(path + f.Name)
		}
	}
}