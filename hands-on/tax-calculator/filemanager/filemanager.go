package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFile  string
	OutPutFile string
}

func New(in, out string) *FileManager {
	return &FileManager{
		InputFile:  in,
		OutPutFile: out,
	}
}

func (fm *FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFile)

	if err != nil {
		return nil, errors.New("Error Opening the file.")
	}

	lines := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Error Reading the lines from file.")
	}
	return lines, nil
}

func (fm *FileManager) WriteResults(data any) error {
	file, err := os.Create(fm.OutPutFile)
	if err != nil {
		return errors.New("Error Creating json file")
	}

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("Error Converting data into json.")
	}
	file.Close()

	return nil

}
