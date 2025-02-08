package count_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/mq/packages/chap3/count"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"count": count.Main,
	}))
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func Test_LinesSuccess(t *testing.T) {
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

func Test_WordsSuccess(t *testing.T) {
	//input := bytes.NewBufferString("hello you\nits a nice day\nchamp\n")
	input := bytes.NewBufferString("1\n2 words\n3 this time")
	c, e := count.NewCounter(
		count.WithInput(input),
	)
	if e != nil {
		t.Fatalf("cant create coiuntie: %v", e.Error())
	}
	want := 6
	got := c.Words()
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
