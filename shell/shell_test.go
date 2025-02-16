package shell_test

import (
	"bytes"
	"io"
	"strings"
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

func TestRunProducesExpectedOutput(t *testing.T) {
	t.Parallel()
	in := strings.NewReader("echo hello\n\n")
	out := new(bytes.Buffer)
	session := shell.NewSession(in, out, io.Discard)
	session.DryRun = true
	session.Run()
	want := "> echo hello\n> > \nsee ya...\n"
	//want := "> hello\n> > \nsee ya...\n"
	got := out.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

// func TestNewSession_CreatesExpectedSession(t *testing.T) {
// 	t.Parallel()
// 	want := shell.Session{
// 		Stdin:  os.Stdin,
// 		Stdout: os.Stdout,
// 		Stderr: os.Stderr,
// 	}
// 	got := *shell.NewSession(os.Stdin, os.Stdout, os.Stderr)

// 	if want != got {
// 		t.Errorf("want %#v, got %#v", want, got)
// 	}
// }
