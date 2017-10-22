package fileManager

import (
	"io/ioutil"
)

type Filer interface {
	GetPath() string
	GetContent() []byte
}

type FileManager struct{}

// GetFileContent : so it can be parsed
func (fM FileManager) GetFileContent(f Filer) (*string, error) {
	dat, err := ioutil.ReadFile(f.GetPath())
	if err != nil {
		return nil, err
	}

	str := string(dat)
	return &str, nil
}

// Write : should be used after the json formatter
func (fM FileManager) Write(f Filer) error {
	return ioutil.WriteFile(f.GetPath(), f.GetContent(), 0644)
}
