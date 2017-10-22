package mocked

import "github.com/err0r500/go-architect/domain"

type FileManager struct{}

func (fM FileManager) GetFileContent(domain.File) (*string, error) {
	return nil, nil
}

func (fM FileManager) Write(domain.File) error {
	return nil
}
