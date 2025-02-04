package count_test

import (
	"bytes"
	"testing"

	"github.com/mq/packages/chap3/count"
)

func TestCountSuccess(t *testing.T) {
	input := bytes.NewBufferString("hello\nyou\ntoo\n")
	c, e := count.NewCounter(
		count.WithInput(input),
	)
	if e != nil {
		t.Fatalf("cant create coiuntie: %v", e.Error())
	}

	want := 3
	got := c.Lines()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
func TestWithInputFromArgs_SetsInputToGivenPath(t *testing.T) {
	args := []string{"testdata/three_lines.txt"}
	c, e := count.NewCounter(
		count.WithInputFromArgs(args),
	)
	if e != nil {
		t.Fatal(e)
	}
	want := 3
	got := c.Lines()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestWithInputFromArgs_IgnoresEmptyArgs(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := count.NewCounter(
		count.WithInput(inputBuf),
		count.WithInputFromArgs([]string{}),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
