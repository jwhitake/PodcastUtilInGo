package files

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"bytes"
	"log"
	"os"
	"os/exec"
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

//GetSubDirectories is for
func GetSubDirectories(path string) ([]string, error) {
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

//GetDirectoryFileNames is for
func GetDirectoryFileNames(path string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
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

// Detect returns a list of file paths pointing to the root folder of
// USB storage devices connected to the system.
// Source: https://github.com/deepakjois/gousbdrivedetector
func Detect() ([]string, error) {
	var drives []string
	driveMap := make(map[string]bool)

	cmd := "wmic"
	args := []string{"logicaldisk", "where", "drivetype=2", "get", "deviceid"}
	out, err := exec.Command(cmd, args...).Output()

	if err != nil {
		return drives, err
	}

	s := bufio.NewScanner(bytes.NewReader(out))
	for s.Scan() {
		line := s.Text()
		if strings.Contains(line, ":") {
			rootPath := strings.TrimSpace(line) + string(os.PathSeparator)
			driveMap[rootPath] = true
		}
	}

	for k := range driveMap {
		_, err := os.Open(k)
		if err == nil {
			drives = append(drives, k)
		}
	}

	return drives, nil
}
