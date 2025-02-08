package writer_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/mq/packages/writer"
)

func TestWriteToFile_NoDataDoesntReturnError(t *testing.T) {
	e := writer.WriteToFile("data", []byte{1, 2, 3})
	if e != nil {
		t.Error("Should not have got error", e)
	}
}

func TestWriteToFile_WritesGivenDataToFile(t *testing.T) {
	want := []byte{1, 2, 3}
	filepath := t.TempDir() + "/write_test.txt"

	e := writer.WriteToFile(filepath, want)
	if e != nil {
		t.Fatal("Should not have got error", e)
	}
	got, e := os.ReadFile(filepath)
	if e != nil {
		t.Fatal("cant read data", e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestWriteToFile_ReturnsErrorForUnwritableFile(t *testing.T) {
	path := "bogusdir/write_test.txt"
	e := writer.WriteToFile(path, []byte{})
	if e == nil {
		t.Fatal("want error when file not writable")
	}
}

func TestWriteToFile_ClobbersExistingFile(t *testing.T) {
	path := t.TempDir() + "/clobber_test.txt"
	e := os.WriteFile(path, []byte{3, 4, 5}, 0o600)
	if e != nil {
		t.Fatal(e)
	}
	want := []byte{1, 2, 3}
	e = writer.WriteToFile(path, want)
	if e != nil {
		t.Fatal(e)
	}
	got, e := os.ReadFile(path)
	if e != nil {
		t.Fatal(e)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestWriteToFile_ChangesPermsOnExistingFile(t *testing.T) {
	path := t.TempDir() + "/perms_test.txt"
	// Pre-create empty file with open perms
	err := os.WriteFile(path, []byte{}, 0o644)
	if err != nil {
		t.Fatal(err)
	}
	err = writer.WriteToFile(path, []byte{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}
	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	perm := stat.Mode().Perm()
	if perm != 0o600 {
		t.Errorf("want file mode 0o600, got 0o%o", perm)
	}
}
