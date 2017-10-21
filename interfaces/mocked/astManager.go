package mocked

type AstManager struct{}

func (astM AstManager) GetImports(fileContent string) (importsPaths *[]string, err error) {
	return nil, nil
}
