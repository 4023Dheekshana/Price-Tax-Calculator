package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Filemanager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm Filemanager) ReadFile() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("file not found")
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading the file failed.")
		fmt.Println(err)
		file.Close()
	}
	return lines, nil
}

func (fm Filemanager) WriteJson(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath) // to create new file os.create()
	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file) //creates new encoder to encode data in file is going to be encoder
	err = encoder.Encode(data)       //encoder helps to encode the datas in the file.
	if err != nil {
		file.Close()
		return errors.New("failed to convert data to json")
	}
	file.Close()
	return nil
}

func New(inputPath, outputPath string) Filemanager {
	return Filemanager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
