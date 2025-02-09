package filecounter_test

import (
	"os"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/mq/packages/filecounter"
)

func TestFiles_CorrectlyListsFilesInTree(t *testing.T) {
	// fsys := fstest.MapFS{
	// 	"filecounter.go":      {},
	// 	"filecounter_test.go": {},
	// }
	fsys := os.DirFS("testdata/tree")
	want := []string{
		"another.go",
		"file.go",
		"subfolder/another.go",
	}
	got := filecounter.Files(fsys)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func BenchmarkFilesOnDisk(b *testing.B) {
	fsys := os.DirFS("testdata/tree")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = filecounter.Files(fsys)
	}
}

func BenchmarkFilesInMemory(b *testing.B) {
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = filecounter.Files(fsys)
	}
}
