package domain

// File, very simple representation of a file
// implementing the Filer interface
type File struct {
	Path    string
	Content []byte
}

func (f File) GetPath() string {
	return f.Path
}

func (f File) GetContent() []byte {
	return f.Content
}
