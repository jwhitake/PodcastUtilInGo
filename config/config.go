package config

import(
	"fmt"
	"encoding/json"
	"os"
	"io/ioutil"

	"github.com/jwhitake/podcast/files"
)

//Configuration holds json config
type Configuration struct {
	SourcePath   string `json:"SourcePath"`
	ArchivePath  string `json:"ArchivePath"`
	ArchiveFiles string `json:"ArchiveFiles"`
}


//GetConfig loads and returns a Configuration struct
func GetConfig() Configuration{
	currentPath := files.GetCurrentDirectory()
	jsonFile, err := os.Open(currentPath + "/config.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(currentPath)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var configuration Configuration
	json.Unmarshal(byteValue, &configuration)
	return configuration
}
