package kafka

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func ReadJsonFiles() [][]byte {
	var byteCollection [][]byte
	var fileAbsolutePrefix string
	var filePath string

	// WIN
	//absPath, _ := filepath.Abs("/go-api")
	//fileAbsolutePrefix = absPath

	// MAC
	//absPath, _ := filepath.Abs("")
	//fileAbsolutePrefix = absPath

	// ALPINE
	//absPath, _ := filepath.Abs("./producer/test_files/")
	//fileAbsolutePrefix = absPath

	files, err := ioutil.ReadDir(fileAbsolutePrefix)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	totalAmountFiles := len(files)

	fmt.Printf("\n START PARSING %d TEST FILES ", totalAmountFiles)
	for i := 0; i < totalAmountFiles; i++ {
		filePath = fileAbsolutePrefix + "/" + strconv.Itoa(i) + ".json"

		fmt.Printf("\n PARSING FILE %d / %d @ %s:  ", i, totalAmountFiles, filePath)
		jsonFile, err := os.Open(filePath)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		byteCollection = append(byteCollection, byteValue)
	}
	return byteCollection
}
