package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
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

	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		// file.Close()
		return nil, errors.New("Error Reading the lines from file.")
	}

	// file.Close()
	return lines, nil
}

func (fm *FileManager) WriteResults(data any) error {

	time.Sleep(3 * time.Second)
	file, err := os.Create(fm.OutPutFile)
	if err != nil {
		return errors.New("Error Creating json file")
	}

	file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)

	if err != nil {
		// file.Close()
		return errors.New("Error Converting data into json.")
	}
	// file.Close()

	return nil

}
