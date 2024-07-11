package local

const dbFile = "etiketfs.db"

// LocalStorage ...
type LocalStorage struct {
	path string
}

// Load ...
func Load(path string) (*LocalStorage, error) {
	return &LocalStorage{path}, nil
}
