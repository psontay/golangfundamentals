package manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputPath  string
	OutputPath string
}

func NewFileManager(inputPath string, outputPath string) FileManager {
	return FileManager{
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
}

func (fileManager FileManager) ReadJSON(data interface{}) error {
	file, err := os.Open(fileManager.InputPath)
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(data)
	if err != nil {
		return fmt.Errorf("cannot decode json: %w", err)
	}
	return nil
}

func (fileManager FileManager) WriteJSON(data interface{}) error {
	file, err := os.Create(fileManager.OutputPath)
	if err != nil {
		return errors.New("Error creating file: " + fileManager.OutputPath)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Error encoding file: " + fileManager.OutputPath)
	}
	return nil
}
