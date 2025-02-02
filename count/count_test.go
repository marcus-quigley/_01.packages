package count_test

import (
	"bytes"
	"testing"

	"github.com/mq/packages/count"
)

func TestCountSuccess(t *testing.T) {
	input := bytes.NewBufferString("hello\nyou\ntoo\n")
	c, e := count.NewCountie(
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
