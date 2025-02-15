package shell_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mq/packages/shell"
)

func TestCmdFromString_CreatesExpectedCmd(t *testing.T) {
	input := "/bin/ls -l main.go"
	want := []string{"/bin/ls", "-l", "main.go"}

	cmd, e := shell.CreateCommand(input)
	if e != nil {
		t.Error(e)
	}
	got := cmd.Args
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCmdFromString_ErrorsOnEmptyInput(t *testing.T) {
	_, e := shell.CreateCommand("")
	if e == nil {
		t.Error("got nil want error on empty input")
	}
}
