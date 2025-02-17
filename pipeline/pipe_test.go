package pipeline_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mq/packages/pipeline"
)

func TestStdoutPrintsMessageToOutput(t *testing.T) {
	want := "hello world\n"
	p := pipeline.FromString(want)

	buf := new(bytes.Buffer)
	p.Output = buf
	p.Stdout()
	if p.Error != nil {
		t.Fatal(p.Error)
	}

	got := buf.String()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestStdoutPrintsNothingOnError(t *testing.T) {
	p := pipeline.FromString("hello world\n")
	p.Error = fmt.Errorf("an error")

	buf := new(bytes.Buffer)
	p.Output = buf
	p.Stdout()
	got := buf.String()
	if got != "" {
		t.Errorf("want no output from Stdout after error, got %q", got)
	}
}

func TestFromFile_ReadsAllDataFromFile(t *testing.T) {

	want := []byte("Hello world\n")
	p := pipeline.FromFile("testdata/hello.txt")
	if p.Error != nil {
		t.Fatal(p.Error)
	}
	got, e := io.ReadAll(p.Input)
	if e != nil {
		t.Fatal(e)
	}
	if !cmp.Equal(got, want) {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestFromFile_SetsErrorGivenNonexistentFile(t *testing.T) {
	t.Parallel()
	p := pipeline.FromFile("doesnt-exist.txt")
	if p.Error == nil {
		t.Fatal("want error opening non-existent file, got nil")
	}
}
