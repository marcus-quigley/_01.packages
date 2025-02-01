package greeting_test

import (
	"bytes"
	"testing"

	greeting "github.com/mq/packages"
)

func TestDisplayName(t *testing.T) {

	want := "hello, marcus"
	buf := &bytes.Buffer{}
	g := greeting.New()
	g.Output = buf
	g.Name = "marcus"

	g.Greet()

	got := buf.String()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestReadName(t *testing.T) {
	want := "marcus"
	g := greeting.New()
	g.Input = bytes.NewBufferString("marcus\n")
	g.ReadName()

	got := g.Name
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

// func TestRun(t *testing.T) {
// 	g := greeting.Greetie{}
// 	g.Run()
// }
