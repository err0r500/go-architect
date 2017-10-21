package mocked

type FileManager struct{}

func (fM FileManager) GetFileContent(path string) (*string, error) {
	return nil, nil
}

func (fM FileManager) WriteToFile() error {
	return nil
}
