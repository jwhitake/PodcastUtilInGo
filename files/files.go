package files

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
		dir.Mode()
		if dir.IsDir() {
			dirList = append(dirList, dir.Name())
		}
	}
	return dirList, err
}

//DeleteFilesByExtension is for
func DeleteFilesByExtension(path, ext string) {
	listing, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range listing {
		if strings.HasSuffix(f.Name(), ext) {
			os.Remove(path + f.Name())
		}
	}
}

//Copy is for 
func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
