package filecounter

import (
	"io/fs"
	"path/filepath"
)

func Files(filesys fs.FS) []string {
	//filesys := os.DirFS(path)
	var paths []string

	fs.WalkDir(filesys, ".",
		func(p string, d fs.DirEntry, e error) error {
			if filepath.Ext(p) == ".go" {
				paths = append(paths, p)
			}
			return nil
		})
	return paths
}
