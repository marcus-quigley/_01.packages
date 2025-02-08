package writer

import "os"

func WriteToFile(path string, data []byte) error {
	e := os.WriteFile(path, data, 0o600)
	if e != nil {
		return e
	}
	return os.Chmod(path, 0o600)
}
